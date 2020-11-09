// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"sync"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type result interface{}
type callBackInfo struct {
	requestID    uint64
	subscription uint64
	stream       chan result
	reflectType  reflect.Type
}

type WSClient struct {
	currentID               uint64
	rpcURL                  string
	conn                    *websocket.Conn
	lock                    sync.Mutex
	callbacksByRequestID    map[uint64]*callBackInfo
	callbacksBySubscription map[uint64]*callBackInfo
}

func NewWSClient(rpcURL string) (*WSClient, error) {
	c := &WSClient{
		currentID:               0,
		rpcURL:                  rpcURL,
		callbacksByRequestID:    map[uint64]*callBackInfo{},
		callbacksBySubscription: map[uint64]*callBackInfo{},
	}

	conn, _, err := websocket.DefaultDialer.Dial(rpcURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new ws client: dial: %w", err)
	}
	c.conn = conn

	c.receiveMessages()

	return c, nil
}

func (c *WSClient) receiveMessages() {
	zlog.Info("ready to receive message")
	go func() {
		k := 0
		for {
			k++
			_, message, err := c.conn.ReadMessage()
			zlog.Info("")
			if err != nil {
				zlog.Error("message reception", zap.Error(err))
				continue
			}

			// when receiving message with id. the result will be a subscription number.
			// that number will be associated to all future message destine to this request
			if gjson.GetBytes(message, "id").Exists() {
				zlog.Info("received subscription for id")
				id := uint64(gjson.GetBytes(message, "id").Int())
				sub := uint64(gjson.GetBytes(message, "result").Int())

				//moving pending callback info to the actual callbacksBySubscription list
				c.lock.Lock()
				callBack := c.callbacksByRequestID[id]
				callBack.subscription = sub
				c.callbacksBySubscription[sub] = callBack
				c.lock.Unlock()

				zlog.Info("move sub from pending to callback", zap.Uint64("id", uint64(id)), zap.Uint64("subscription", uint64(sub)))
				continue
			}

			//getting the callback
			sub := uint64(gjson.GetBytes(message, "params.subscription").Int())
			callBack := c.callbacksBySubscription[sub]

			//getting and instantiate the return type for the call back.
			resultType := reflect.New(callBack.reflectType)
			result := resultType.Interface()

			err = decodeClientResponse(bytes.NewReader(message), &result)
			if err != nil {
				zlog.Error("failed to decode result", zap.Uint64("subscription", uint64(sub)), zap.Error(err))
			}
			callBack.stream <- result
		}
	}()
}

type clientRequest struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint64      `json:"id"`
}

func encodeClientRequest(method string, args interface{}) ([]byte, uint64, error) {
	c := &clientRequest{
		Version: "2.0",
		Method:  method,
		Params:  args,
		Id:      uint64(rand.Int63()),
	}
	data, err := json.Marshal(c)
	if err != nil {
		return nil, 0, fmt.Errorf("encode request: json marshall: %w", err)
	}
	return data, c.Id, nil
}

type wsClientResponse struct {
	Version string                  `json:"jsonrpc"`
	Params  *wsClientResponseParams `json:"params"`
	Error   *json.RawMessage        `json:"error"`
}

type wsClientResponseParams struct {
	Result       *json.RawMessage `json:"result"`
	Subscription int              `json:"subscription"`
}

func decodeClientResponse(r io.Reader, reply interface{}) (err error) {
	var c *wsClientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return err
	}

	if c.Error != nil {
		jsonErr := &json2.Error{}
		if err := json.Unmarshal(*c.Error, jsonErr); err != nil {
			return &json2.Error{
				Code:    json2.E_SERVER,
				Message: string(*c.Error),
			}
		}
		return jsonErr
	}

	if c.Params == nil {
		return json2.ErrNullResult
	}

	return json.Unmarshal(*c.Params.Result, &reply)
}

type ProgramWSResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		Account Account `json:"account"`
	} `json:"value"`
}

func (c *WSClient) ProgramSubscribe(programID string, commitment CommitmentType) (stream chan result, id uint64, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	stream = make(chan result, 200)

	params := []interface{}{programID}
	conf := map[string]interface{}{
		"encoding": "jsonParsed",
	}
	if commitment != "" {
		conf["commitment"] = string(commitment)
	}

	params = append(params, conf)
	data, id, err := encodeClientRequest("programSubscribe", params)
	if err != nil {
		return nil, 0, fmt.Errorf("program subscribe: encode request: %c", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, 0, fmt.Errorf("program subscribe: write message: %c", err)
	}

	c.callbacksByRequestID[id] = &callBackInfo{
		requestID:   id,
		stream:      stream,
		reflectType: reflect.TypeOf(ProgramWSResult{}),
	}

	return stream, id, nil
}

func (c *WSClient) ProgramUnsubscribe(reqID int) {

}
