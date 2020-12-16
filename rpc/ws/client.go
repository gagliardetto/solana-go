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

package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sync"
	"time"

	"github.com/dfuse-io/solana-go/rpc"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type result interface{}

type Client struct {
	rpcURL                  string
	conn                    *websocket.Conn
	lock                    sync.RWMutex
	subscriptionByRequestID map[uint64]*Subscription
	subscriptionByWSSubID   map[uint64]*Subscription
	reconnectOnErr          bool
}

func Dial(ctx context.Context, rpcURL string) (c *Client, err error) {
	c = &Client{
		rpcURL:                  rpcURL,
		subscriptionByRequestID: map[uint64]*Subscription{},
		subscriptionByWSSubID:   map[uint64]*Subscription{},
	}

	c.conn, _, err = websocket.DefaultDialer.DialContext(ctx, rpcURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new ws client: dial: %w", err)
	}

	go func() {
		for {
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
			time.Sleep(20 * time.Second)
		}
	}()
	go c.receiveMessages()
	return c, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) receiveMessages() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.closeAllSubscription(err)
			return
		}
		c.handleMessage(message)
	}
}

func (c *Client) handleMessage(message []byte) {
	// when receiving message with id. the result will be a subscription number.
	// that number will be associated to all future message destine to this request
	if gjson.GetBytes(message, "id").Exists() {
		requestID := uint64(gjson.GetBytes(message, "id").Int())
		subID := uint64(gjson.GetBytes(message, "result").Int())
		c.handleNewSubscriptionMessage(requestID, subID)
		return
	}

	c.handleSubscriptionMessage(uint64(gjson.GetBytes(message, "params.subscription").Int()), message)

}

func (c *Client) handleNewSubscriptionMessage(requestID, subID uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if traceEnabled {
		zlog.Debug("received new subscription message",
			zap.Uint64("message_id", requestID),
			zap.Uint64("subscription_id", subID),
		)
	}

	callBack, found := c.subscriptionByRequestID[requestID]
	if !found {
		zlog.Error("cannot find websocket message handler for a new stream.... this should not happen",
			zap.Uint64("request_id", requestID),
			zap.Uint64("subscription_id", subID),
		)
		return
	}
	callBack.subID = subID
	c.subscriptionByWSSubID[subID] = callBack

	zlog.Debug("registered ws subscription",
		zap.Uint64("subscription_id", subID),
		zap.Uint64("request_id", requestID),
		zap.Int("subscription_count", len(c.subscriptionByWSSubID)),
	)
	return
}

func (c *Client) handleSubscriptionMessage(subID uint64, message []byte) {
	if traceEnabled {
		zlog.Debug("received subscription message",
			zap.Uint64("subscription_id", subID),
		)
	}

	c.lock.RLock()
	sub, found := c.subscriptionByWSSubID[subID]
	c.lock.RUnlock()
	if !found {
		zlog.Warn("unable to find subscription for ws message", zap.Uint64("subscription_id", subID))
		return
	}

	//getting and instantiate the return type for the call back.
	resultType := reflect.New(sub.reflectType)
	result := resultType.Interface()
	err := decodeResponse(bytes.NewReader(message), &result)
	if err != nil {
		fmt.Println("*****************************")
		c.closeSubscription(sub.req.ID, fmt.Errorf("unable to decode client response: %w", err))
		return
	}

	// this cannot be blocking or else
	// we  will no read any other message
	if len(sub.stream) >= cap(sub.stream) {
		zlog.Warn("closing ws client subscription... not consuming fast en ought",
			zap.Uint64("request_id", sub.req.ID),
		)
		c.closeSubscription(sub.req.ID, fmt.Errorf("reached channel max capacity %d", len(sub.stream)))
		return
	}

	sub.stream <- result
	return
}

func (c *Client) closeAllSubscription(err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, sub := range c.subscriptionByRequestID {
		sub.err <- err
	}

	c.subscriptionByRequestID = map[uint64]*Subscription{}
	c.subscriptionByWSSubID = map[uint64]*Subscription{}
}

func (c *Client) closeSubscription(reqID uint64, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	sub, found := c.subscriptionByRequestID[reqID]
	if !found {
		return
	}

	sub.err <- err

	err = c.unsubscribe(sub.subID, sub.unsubscribeMethod)
	if err != nil {
		zlog.Warn("unable to send rpc unsubscribe call",
			zap.Error(err),
		)
	}

	delete(c.subscriptionByRequestID, sub.req.ID)
	delete(c.subscriptionByWSSubID, sub.subID)
}

func (c *Client) unsubscribe(subID uint64, method string) error {
	req := newRequest([]interface{}{subID}, method, map[string]interface{}{})
	data, err := req.encode()
	if err != nil {
		return fmt.Errorf("unable to encode unsubscription message for subID %d and method %s", subID, method)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return fmt.Errorf("unable to send unsubscription message for subID %d and method %s", subID, method)
	}
	return nil
}

func (c *Client) subscribe(params []interface{}, conf map[string]interface{}, subscriptionMethod, unsubscribeMethod string, commitment rpc.CommitmentType, resultType interface{}) (*Subscription, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if commitment != "" {
		conf["commitment"] = string(commitment)
	}

	req := newRequest(params, subscriptionMethod, conf)
	data, err := req.encode()
	if err != nil {
		return nil, fmt.Errorf("subscribe: unable to encode subsciption request: %w", err)
	}

	sub := newSubscription(req, reflect.TypeOf(resultType), func(err error) {
		c.closeSubscription(req.ID, err)
	}, unsubscribeMethod)

	c.subscriptionByRequestID[req.ID] = sub
	zlog.Info("added new subscription to websocket client", zap.Int("count", len(c.subscriptionByRequestID)))

	zlog.Debug("writing data to conn", zap.String("data", string(data)))
	err = c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, fmt.Errorf("unable to write request: %w", err)
	}

	return sub, nil
}

func decodeResponse(r io.Reader, reply interface{}) (err error) {
	var c *response
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
