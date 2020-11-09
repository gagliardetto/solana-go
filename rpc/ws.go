package rpc

import (
	"fmt"
	"log"

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
		"encoding": "base64",
	}
	if commitment != "" {
		conf["commitment"] = string(commitment)
	}

	params = append(params, conf)
	data, err := json2.EncodeClientRequest("programSubscribe", params)
	if err != nil {
		return fmt.Errorf("program subscribe: encode request: %w", err)
	}

	conn, err := w.getConnection()
	if err != nil {
		return fmt.Errorf("program subscribe: get connection: %w", err)
	}

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return fmt.Errorf("program subscribe: write message: %w", err)
	}

	return nil
}
