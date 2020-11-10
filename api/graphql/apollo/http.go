package apollo

import (
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const protocolGraphQLWS = "graphql-ws"

var upgrader = websocket.Upgrader{
	CheckOrigin:  func(r *http.Request) bool { return true },
	Subprotocols: []string{protocolGraphQLWS},
}

type Middleware struct {
	service GraphQLService
}

func NewMiddleware(service GraphQLService) *Middleware {
	return &Middleware{service: service}
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		connectionProtocols := websocket.Subprotocols(r)
		for _, subprotocol := range connectionProtocols {
			if subprotocol == protocolGraphQLWS {
				ws, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					zlog.Debug("unable to upgrade HTTP connection", zap.Error(err))
					return
				}

				if ws.Subprotocol() != protocolGraphQLWS {
					zlog.Debug("created websocket connection is not using right subprotocol",
						zap.String("expected_protocol", protocolGraphQLWS),
						zap.String("actual_protocol", ws.Subprotocol()),
					)

					ws.Close()
					return
				}

				zlog.Debug("websocket connection initialized correctly, continuing connection process")
				go Connect(ws, m.service)
				return
			}
		}

		zlog.Debug("this connection didn't had expected protocol, assuming it's a normal HTTP connection",
			zap.String("expected_protocol", protocolGraphQLWS),
			zap.Strings("received_protocols", connectionProtocols),
		)

		next.ServeHTTP(w, r)
	})
}
