package ws

import (
	"encoding/json"
	"fmt"
	"math/rand"

	bin "github.com/dfuse-io/binary"
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

type AccountResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		rpc.Account
	} `json:"value"`
}

type LogResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		Signature solana.Signature `json:"signature"`
		Err       interface{}      `json:"err"`
		Logs      []string         `json:"logs"`
	} `json:"value"`
}

type ProgramResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		PubKey  solana.PublicKey `json:"pubkey"`
		Account rpc.Account      `json:"account"`
	} `json:"value"`
}

type SignatureResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		Err interface{} `json:"err"`
	} `json:"value"`
}

type SlotResult struct {
	Parent uint64 `json:"parent"`
	Root   uint64 `json:"root"`
	Slot   uint64 `json:"slot"`
}

type RootResult bin.Uint64

type VoteResult struct {
	Hash  solana.Hash  `json:"hash"`
	Slots []bin.Uint64 `json:"slots"`
	// TODO:
	// Timestamp interface{} `json:"timestamp"`
}
