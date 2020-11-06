package rpc

import (
	"context"
	"net/http"

	//"github.com/dfuse-io/solana-go"
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

func (c *Client) GetBalance(ctx context.Context, publicKey string, commitment CommitmentType) (out *GetBalanceResult, err error) {
	params := []interface{}{publicKey}
	if commitment != "" {
		params = append(params, string(commitment))
	}

	err = c.rpcClient.CallFor(&out, "getBalance", params...)
	return
}

func (c *Client) GetRecentBlockhash(ctx context.Context, commitment CommitmentType) (out *GetRecentBlockhashResult, err error) {
	var params []interface{}
	if commitment != "" {
		params = append(params, string(commitment))
	}

	err = c.rpcClient.CallFor(&out, "getRecentBlockhash", params...)
	return
}

func (c *Client) GetSlot(ctx context.Context, commitment CommitmentType) (out GetSlotResult, err error) {
	var params []interface{}
	if commitment != "" {
		params = append(params, string(commitment))
	}

	err = c.rpcClient.CallFor(&out, "getSlot", params...)
	return
}

func (c *Client) GetConfirmedBlock(ctx context.Context, slot uint64, encoding string) (out *GetConfirmedBlockResult, err error) {
	if encoding == "" {
		encoding = "json"
	}
	params := []interface{}{slot, encoding}

	err = c.rpcClient.CallFor(&out, "getConfirmedBlock", params...)
	return
}

func (c *Client) GetAccountInfo(ctx context.Context, publicKey string) (out *GetAccountInfoResult, err error) {
	obj := map[string]interface{}{
		"encoding": "base64",
	}
	params := []interface{}{publicKey, obj}

	err = c.rpcClient.CallFor(&out, "getAccountInfo", params...)
	return
}

func (c *Client) GetProgramAccounts(ctx context.Context, publicKey string, opts *GetProgramAccountsOpts) (out *GetProgramAccountsResult, err error) {
	params := []interface{}{publicKey}
	if opts != nil {
		params = append(params, opts)
	}

	err = c.rpcClient.CallFor(&out, "getProgramAccounts", params...)
	return
}
