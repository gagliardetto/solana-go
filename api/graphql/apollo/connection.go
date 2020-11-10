package apollo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/graph-gophers/graphql-go"
	gqerrors "github.com/graph-gophers/graphql-go/errors"
	"go.uber.org/zap"
)

type operationMessageType string

// https://github.com/apollographql/subscriptions-transport-ws/blob/a56491c6feacd96cab47b7a3df8c2cb1b6a96e36/src/message-types.ts
const (
	typeComplete            operationMessageType = "complete"
	typeConnectionAck       operationMessageType = "connection_ack"
	typeConnectionError     operationMessageType = "connection_error"
	typeConnectionInit      operationMessageType = "connection_init"
	typeConnectionKeepAlive operationMessageType = "ka"
	typeConnectionTerminate operationMessageType = "connection_terminate"
	typeData                operationMessageType = "data"
	typeError               operationMessageType = "error"
	typeStart               operationMessageType = "start"
	typeStop                operationMessageType = "stop"

	// We support "pong" as keep alive message to support `@dfuse/client` JavaScript client sending them as the keep alive message
	typePong operationMessageType = "pong"
)

type wsConnection interface {
	Close() error
	ReadJSON(v interface{}) error
	SetReadLimit(limit int64)
	SetWriteDeadline(t time.Time) error
	WriteMessage(messageType int, data []byte) error
	WriteJSON(v interface{}) error
}

type sendFunc func(id string, omType operationMessageType, payload json.RawMessage)

type operationMessage struct {
	ID      string               `json:"id,omitempty"`
	Payload json.RawMessage      `json:"payload,omitempty"`
	Type    operationMessageType `json:"type"`
}

type startMessagePayload struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

// GraphQLService interface
type GraphQLService interface {
	Subscribe(ctx context.Context, document string, operationName string, variableValues map[string]interface{}) (payloads <-chan interface{}, err error)
}

type connection struct {
	cancel       func()
	service      GraphQLService
	writeTimeout time.Duration
	ws           wsConnection
	request      *http.Request
}

// ReadLimit limits the maximum size of incoming messages
func ReadLimit(limit int64) func(conn *connection) {
	return func(conn *connection) {
		conn.ws.SetReadLimit(limit)
	}
}

// WriteTimeout sets a timeout for outgoing messages
func WriteTimeout(d time.Duration) func(conn *connection) {
	return func(conn *connection) {
		conn.writeTimeout = d
	}
}

// Connect implements the apollographql subscriptions-transport-ws protocol@v0.9.4
// https://github.com/apollographql/subscriptions-transport-ws/blob/v0.9.4/PROTOCOL.md
func Connect(ws wsConnection, service GraphQLService, options ...func(conn *connection)) {
	conn := &connection{
		service: service,
		ws:      ws,
	}

	defaultOpts := []func(conn *connection){
		ReadLimit(4096),
		WriteTimeout(10 * time.Second),
	}

	for _, opt := range append(defaultOpts, options...) {
		opt(conn)
	}

	ctx, cancel := context.WithCancel(context.Background())
	conn.cancel = cancel

	// This is a blocking call and share the connection lifecycle, so will end only when connection closes
	zlog.Debug("starting read loop blocking call")
	conn.readLoop(ctx, conn.writeLoop(ctx))
}

func (conn *connection) writeLoop(ctx context.Context) sendFunc {
	stop := make(chan struct{})
	out := make(chan *operationMessage)

	send := func(id string, omType operationMessageType, payload json.RawMessage) {
		select {
		case <-stop:
			return
		case out <- &operationMessage{ID: id, Type: omType, Payload: payload}:
		}
	}

	go func() {
		defer close(stop)
		defer conn.close()

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-out:
				select {
				case <-ctx.Done():
					return
				default:
				}

				zlog.Debug("setting connection timeout value", zap.Duration("write_timeout", conn.writeTimeout))
				if err := conn.ws.SetWriteDeadline(time.Now().Add(conn.writeTimeout)); err != nil {
					return
				}

				if err := conn.ws.WriteJSON(msg); err != nil {
					return
				}
			}
		}
	}()

	return send
}

func (conn *connection) close() {
	zlog.Info("closing websocket connection")
	conn.cancel()
	conn.ws.Close()
}

func (conn *connection) readLoop(ctx context.Context, send sendFunc) {
	defer conn.close()

	opDone := map[string]func(){}
	for {
		var msg operationMessage
		err := conn.ws.ReadJSON(&msg)
		if err != nil {
			zlog.Debug("got an error while trying to read message from websocket")
			return
		}

		switch msg.Type {
		case typeConnectionInit:
			var initMsg map[string]interface{}
			if err := json.Unmarshal(msg.Payload, &initMsg); err != nil {
				ep := errPayload(fmt.Errorf("invalid payload for type: %s", msg.Type))
				send("", typeConnectionError, ep)
				continue
			}

			send("", typeConnectionAck, nil)

		case typeStart:
			if msg.ID == "" {
				ep := errPayload(errors.New("missing mongoID for start operation"))
				send("", typeConnectionError, ep)
				continue
			}

			var osp startMessagePayload
			if err := json.Unmarshal(msg.Payload, &osp); err != nil {
				ep := errPayload(fmt.Errorf("invalid payload for type: %s", msg.Type))
				send(msg.ID, typeConnectionError, ep)
				continue
			}

			zlog.Debug("starting stream due to start message received from client", zap.String("id", msg.ID))

			opCtx, cancel := context.WithCancel(ctx)

			c, err := conn.service.Subscribe(opCtx, osp.Query, osp.OperationName, osp.Variables)
			if err != nil {
				cancel()
				send(msg.ID, typeError, errPayload(err))
				continue
			}

			opDone[msg.ID] = cancel

			go func() {
				defer cancel()
				for {
					select {
					case <-opCtx.Done():
						return

					case payload, more := <-c:
						if !more {
							zlog.Info("notifying completion of stream due to no more data", zap.String("id", msg.ID))
							send(msg.ID, typeComplete, nil)
							return
						}

						if resp, ok := payload.(*graphql.Response); ok {
							// We assume there will be a single error, we should handle the multi-error variant one day...
							terminalErr := getTerminalQueryError(resp.Errors)
							if terminalErr != nil {
								var err = resp.Errors[0]
								zlog.Info("notifying completion of stream due to service error", zap.String("id", msg.ID), zap.Error(err))
								send(msg.ID, typeError, errPayload(err))
								return
							}
						}

						jsonPayload, err := json.Marshal(payload)
						if err != nil {
							send(msg.ID, typeError, errPayload(err))
							return
						}

						send(msg.ID, typeData, jsonPayload)
					}
				}
			}()

		case typePong:
			// We simply discard pong messages as they act as keep alive messages

		case typeStop:
			zlog.Debug("stopping stream due to stop message received from client", zap.String("id", msg.ID))
			onDone, ok := opDone[msg.ID]
			if ok {
				delete(opDone, msg.ID)
				onDone()
			}

		case typeConnectionTerminate:
			zlog.Info("terminating client connection")
			err := conn.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				zlog.Warn("unable to send close message to client, discarding")
				return
			}

			return

		default:
			ep := errPayload(fmt.Errorf("unknown operation message of type: %s", msg.Type))
			send(msg.ID, typeError, ep)
		}
	}
}

func errPayload(err error) json.RawMessage {
	b, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})

	return b
}

func getTerminalQueryError(responseErrors []*gqerrors.QueryError) error {
	if len(responseErrors) <= 0 {
		return nil
	}

	for _, err := range responseErrors {
		if err.Extensions == nil {
			continue
		}

		value, exist := err.Extensions["terminal"]
		if !exist {
			continue
		}

		if v, ok := value.(bool); ok && v {
			return err
		}
	}

	return nil
}
