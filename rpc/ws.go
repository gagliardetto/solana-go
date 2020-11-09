package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/gorilla/websocket"
)

type WSClient struct {
	rpcURL string
	conn   *websocket.Conn
}

func NewWSClient(rpcURL string) *WSClient {
	return &WSClient{
		rpcURL: rpcURL,
	}
}

func (w *WSClient) getConnection() (*websocket.Conn, error) {
	if w.conn != nil {
		return w.conn, nil
	}
	c, _, err := websocket.DefaultDialer.Dial(w.rpcURL, nil)
	if err != nil {
		return nil, fmt.Errorf("ws dial: %w", err)
	}
	w.conn = c
	return c, nil
}

func (w *WSClient) ProgramSubscribe(programID string, commitment CommitmentType) error {
	params := []interface{}{programID}
	conf := map[string]interface{}{
		"encoding": "jsonParsed",
	}
	if commitment != "" {
		conf["commitment"] = string(commitment)
	}

	params = append(params, conf)
	data, err := encodeClientRequest("programSubscribe", params)
	if err != nil {
		return fmt.Errorf("program subscribe: encode request: %w", err)
	}

	conn, err := w.getConnection()
	if err != nil {
		return fmt.Errorf("program subscribe: get connection: %w", err)
	}

	go func() {
		k := 0
		for {
			k++
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			o := map[string]interface{}{}
			subscription, err := decodeClientResponse(bytes.NewReader(message), &o)
			if err != nil {
				panic(err)
			}
			log.Printf("recv: %d: %s", subscription, o)
		}
	}()

	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return fmt.Errorf("program subscribe: write message: %w", err)
	}

	return nil
}

type clientRequest struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params interface{} `json:"params"`

	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id uint64 `json:"id"`
}

// EncodeClientRequest encodes parameters for a JSON-RPC client request.
func encodeClientRequest(method string, args interface{}) ([]byte, error) {
	c := &clientRequest{
		Version: "2.0",
		Method:  method,
		Params:  args,
		Id:      uint64(rand.Int63()),
	}
	return json.Marshal(c)
}

type wsClientResponse struct {
	Version string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Result  int                     `json:"result"` //Ça c'est pas cool
	Params  *wsClientResponseParams `json:"params"`
	Error   *json.RawMessage        `json:"error"`
}

type wsClientResponseParams struct {
	Result       *json.RawMessage `json:"result"`
	Subscription int              `json:"subscription"`
}

func decodeClientResponse(r io.Reader, reply interface{}) (subscriptionID int, err error) {
	var c *wsClientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return -1, err
	}

	if c.Error != nil {
		jsonErr := &json2.Error{}
		if err := json.Unmarshal(*c.Error, jsonErr); err != nil {
			return -1, &json2.Error{
				Code:    json2.E_SERVER,
				Message: string(*c.Error),
			}
		}
		return -1, jsonErr
	}

	if c.Method == "" {
		return c.Result, nil
	}

	if c.Params == nil {
		return -1, json2.ErrNullResult
	}

	return c.Params.Subscription, json.Unmarshal(*c.Params.Result, reply)
}
