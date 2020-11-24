package ws

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/dfuse-io/solana-go"

	"github.com/dfuse-io/solana-go/rpc"
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

type ProgramResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		PubKey  solana.PublicKey `json:"pub_key"`
		Account rpc.Account      `json:"account"`
	} `json:"value"`
}

type AccountResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		Account rpc.Account `json:"account"`
	} `json:"value"`
}

type SlotResult struct {
	Parent uint64
	Root   uint64
	Slot   uint64
}
