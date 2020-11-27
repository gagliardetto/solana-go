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
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/ybbus/jsonrpc"
)

var ErrNotFound = errors.New("not found")

type Client struct {
	rpcURL    string
	rpcClient jsonrpc.RPCClient
	headers   http.Header
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
	commit := map[string]string{
		"commitment": string(commitment),
	}
	var params []interface{}
	if commitment != "" {
		params = append(params, commit)
	}

	err = c.rpcClient.CallFor(&out, "getRecentBlockhash", params)
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

func (c *Client) GetAccountInfo(ctx context.Context, account solana.PublicKey) (out *GetAccountInfoResult, err error) {
	obj := map[string]interface{}{
		"encoding": "base64",
	}
	params := []interface{}{account, obj}

	err = c.rpcClient.CallFor(&out, "getAccountInfo", params...)
	if err != nil {
		return nil, err
	}

	if out.Value == nil {
		return nil, ErrNotFound
	}

	return out, nil
}

func (c *Client) GetAccountDataIn(ctx context.Context, account solana.PublicKey, inVar interface{}) (err error) {
	resp, err := c.GetAccountInfo(ctx, account)
	if err != nil {
		return err
	}

	return bin.NewDecoder(resp.Value.Data).Decode(inVar)
}

func (c *Client) GetConfirmedTransaction(ctx context.Context, signature string) (out TransactionWithMeta, err error) {
	params := []interface{}{signature, "json"}

	err = c.rpcClient.CallFor(&out, "getConfirmedTransaction", params...)
	return
}

func (c *Client) GetConfirmedSignaturesForAddress2(ctx context.Context, address solana.PublicKey, opts *GetConfirmedSignaturesForAddress2Opts) (out GetConfirmedSignaturesForAddress2Result, err error) {

	params := []interface{}{address.String(), opts}

	err = c.rpcClient.CallFor(&out, "getConfirmedSignaturesForAddress2", params...)
	return
}

func (c *Client) GetProgramAccounts(ctx context.Context, publicKey solana.PublicKey, opts *GetProgramAccountsOpts) (out GetProgramAccountsResult, err error) {
	obj := map[string]interface{}{
		"encoding": "base64",
	}
	if opts != nil {
		if opts.Commitment != "" {
			obj["commitment"] = string(opts.Commitment)
		}
		if len(opts.Filters) != 0 {
			obj["filters"] = opts.Filters
		}
	}

	params := []interface{}{publicKey, obj}

	err = c.rpcClient.CallFor(&out, "getProgramAccounts", params...)
	return
}

func (c *Client) GetMinimumBalanceForRentExemption(ctx context.Context, dataSize int) (lamport int, err error) {
	params := []interface{}{dataSize}
	err = c.rpcClient.CallFor(&lamport, "getMinimumBalanceForRentExemption", params...)
	return
}

type SimulateTransactionResponse struct {
	Err  interface{}
	Logs []string
}

func (c *Client) SimulateTransaction(ctx context.Context, transaction *solana.Transaction) (*SimulateTransactionResponse, error) {
	buf := new(bytes.Buffer)
	if err := bin.NewEncoder(buf).Encode(transaction); err != nil {
		return nil, fmt.Errorf("send transaction: encode transaction: %w", err)
	}
	trxData := buf.Bytes()

	obj := map[string]interface{}{
		"encoding": "base64",
	}

	b64Data := base64.StdEncoding.EncodeToString(trxData)
	params := []interface{}{
		b64Data,
		obj,
	}

	var out *SimulateTransactionResponse
	if err := c.rpcClient.CallFor(&out, "simulateTransaction", params...); err != nil {
		return nil, fmt.Errorf("send transaction: rpc send: %w", err)
	}

	return out, nil

}

func (c *Client) SendTransaction(ctx context.Context, transaction *solana.Transaction) (signature string, err error) {

	buf := new(bytes.Buffer)

	if err := bin.NewEncoder(buf).Encode(transaction); err != nil {
		return "", fmt.Errorf("send transaction: encode transaction: %w", err)
	}

	trxData := buf.Bytes()

	obj := map[string]interface{}{
		"encoding": "base64",
	}

	params := []interface{}{
		base64.StdEncoding.EncodeToString(trxData),
		obj,
	}

	if err := c.rpcClient.CallFor(&signature, "sendTransaction", params...); err != nil {
		return "", fmt.Errorf("send transaction: rpc send: %w", err)
	}
	return
}

func (c *Client) RequestAirdrop(ctx context.Context, account *solana.PublicKey, lamport uint64, commitment CommitmentType) (signature string, err error) {

	obj := map[string]interface{}{
		"commitment": commitment,
	}

	params := []interface{}{
		account.String(),
		lamport,
		obj,
	}

	if err := c.rpcClient.CallFor(&signature, "requestAirdrop", params...); err != nil {
		return "", fmt.Errorf("send transaction: rpc send: %w", err)
	}
	return
}
