package ws

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type request struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      uint64      `json:"id"`
}

func newRequest(params []interface{}, method string, configuration map[string]interface{}) *request {
	if params != nil && configuration != nil {
		params = append(params, configuration)
	}
	return &request{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      uint64(rand.Int63()),
	}
}

func (c *request) encode() ([]byte, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("encode request: json marshall: %w", err)
	}
	return data, nil
}

type response struct {
	Version string           `json:"jsonrpc"`
	Params  *params          `json:"params"`
	Error   *json.RawMessage `json:"error"`
}

type params struct {
	Result       *json.RawMessage `json:"result"`
	Subscription int              `json:"subscription"`
}
