package solana

import (
	"context"

	"github.com/ybbus/jsonrpc"
)

type Client struct {
	rpcURL    string
	rpcClient jsonrpc.RPCClient
}

func NewClient(rpcURL string) *Client {
	rpcClient := jsonrpc.NewClient(rpcURL)
	return &Client{
		rpcURL:    rpcURL,
		rpcClient: rpcClient,
	}
}

func (c *Client) GetBalance(ctx context.Context, publicKey string, commitment CommitmentType) (out *GetBalanceRPCResult, err error) {
	params := []interface{}{publicKey}
	if commitment != "" {
		params = append(params, string(commitment))
	}

	err = c.rpcClient.CallFor(&out, "getBalance", params...)
	if err != nil {
		return nil, err
	}

	return
}

func (c *Client) GetAccountInfo(ctx context.Context, publicKey string, commitment CommitmentType) (out *GetAccountInfoRPCResult, err error) {
	params := []interface{}{publicKey}
	if commitment != "" {
		params = append(params, string(commitment))
	}

	err = c.rpcClient.CallFor(&out, "getAccountInfo", params...)
	if err != nil {
		return nil, err
	}

	return
}
