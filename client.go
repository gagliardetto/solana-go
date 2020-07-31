package solana

import (
	"context"
	"net/http"

	"github.com/ybbus/jsonrpc"
)

type Client struct {
	rpcURL    string
	rpcClient jsonrpc.RPCClient
	headers   http.Header

	Debug bool
}

func NewClient(rpcURL string) *Client {
	rpcClient := jsonrpc.NewClient(rpcURL)
	return &Client{
		rpcURL:    rpcURL,
		rpcClient: rpcClient,
	}
}

func (c *Client) SetHeader(k, v string) {
	if c.headers == nil {
		c.headers = http.Header{}
	}
	c.headers.Set(k, v)
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

func (c *Client) GetProgramAccounts(ctx context.Context, publicKey string, opts *GetProgramAccountsOpts) (out *GetProgramAccountsRPCResult, err error) {
	params := []interface{}{publicKey}
	if opts != nil {
		params = append(params, opts)
	}

	err = c.rpcClient.CallFor(&out, "getProgramAccounts", params...)
	if err != nil {
		return nil, err
	}

	return
}
