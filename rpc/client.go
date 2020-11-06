package rpc

import (
	"fmt"

	"github.com/dfuse-io/solana-go"
	"github.com/ybbus/jsonrpc"
)

type Client struct {
	client             jsonrpc.RPCClient
	solanaEndpointAddr string
}

func NewClient(endpointAddr string) *Client {
	return &Client{
		solanaEndpointAddr: endpointAddr,
		client:             jsonrpc.NewClient(endpointAddr),
	}
}

func (c *Client) GetAccountInfo(publicKey solana.PublicKey) (*solana.GetAccountInfoResult, error) {
	cfg := map[string]interface{}{
		"encoding": "base64",
	}

	var reply solana.GetAccountInfoResult
	err := c.client.CallFor(&reply, "getAccountInfo", publicKey.String(), cfg)

	if err != nil {
		return nil, fmt.Errorf("getAccountInfo: %w", err)
	}
	return &reply, nil
}

func (c *Client) GetConfirmedBlock(slot uint64) error {
	panic("need to full model here ...")
	var reply map[string]interface{}
	err := c.client.CallFor(&reply, "getConfirmedBlock", slot, "json")

	if err != nil {
		return fmt.Errorf("calling getConfirmedBlock: %w", err)
	}

	fmt.Println("reply:", reply)
	return nil
}
