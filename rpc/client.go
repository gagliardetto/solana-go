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
	return NewClientWithOpts(rpcURL, nil)
}

func NewClientWithOpts(rpcURL string, opts *jsonrpc.RPCClientOpts) *Client {
	rpcClient := jsonrpc.NewClientWithOpts(rpcURL, opts)
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

// GetBalance returns the balance of the account of provided publicKey.
func (cl *Client) GetBalance(
	ctx context.Context,
	publicKey string, // Pubkey of account to query, as base-58 encoded string
	commitment CommitmentType,
) (out *GetBalanceResult, err error) {
	params := []interface{}{publicKey}
	if commitment != "" {
		params = append(params, string(commitment))
	}

	err = cl.rpcClient.CallFor(&out, "getBalance", params)

	return
}

// GetRecentBlockhash returns a recent block hash from the ledger,
// and a fee schedule that can be used to compute the cost of submitting a transaction using it.
func (c *Client) GetRecentBlockhash(ctx context.Context, commitment CommitmentType) (out *GetRecentBlockhashResult, err error) {
	var params []interface{}
	if commitment != "" {
		commit := map[string]string{
			"commitment": string(commitment),
		}
		params = append(params, commit)
	}

	err = c.rpcClient.CallFor(&out, "getRecentBlockhash", params)
	return
}

// GetSlot returns the current slot the node is processing.
func (c *Client) GetSlot(ctx context.Context, commitment CommitmentType) (out bin.Uint64, err error) {
	var params []interface{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}

	err = c.rpcClient.CallFor(&out, "getSlot", params)
	return
}

// GetAccountInfo returns all information associated with the account of provided publicKey.
func (cl *Client) GetAccountInfo(ctx context.Context, account solana.PublicKey) (out *GetAccountInfoResult, err error) {
	return cl.GetAccountInfoWithOpts(
		ctx,
		account,
		EncodingBase64,
		"",
		nil,
		nil,
	)
}

type M map[string]interface{}

type EncodingType string

const (
	EncodingBase58     EncodingType = "base58"      // limited to Account data of less than 129 bytes
	EncodingBase64     EncodingType = "base64"      // will return base64 encoded data for Account data of any size
	EncodingBase64Zstd EncodingType = "base64+zstd" // compresses the Account data using Zstandard and base64-encodes the result

	// attempts to use program-specific state parsers to
	// return more human-readable and explicit account state data.
	// If "jsonParsed" is requested but a parser cannot be found,
	// the field falls back to "base64" encoding, detectable when the data field is type <string>.
	// Cannot be used if specifying dataSlice parameters (offset, length).
	EncodingJSONParsed EncodingType = "jsonParsed"

	EncodingJSON EncodingType = "json"
)

// GetAccountInfoWithOpts returns all information associated with the account of provided publicKey.
// You can limit the returned account data with the offset and length parameters.
// You can specify the encoding of the returned data with the encoding parameter.
func (cl *Client) GetAccountInfoWithOpts(
	ctx context.Context,
	account solana.PublicKey,
	encoding EncodingType,
	commitment CommitmentType,
	offset *uint,
	length *uint,
) (out *GetAccountInfoResult, err error) {

	obj := M{}

	if encoding != "" {
		obj["encoding"] = encoding
	}
	if commitment != "" {
		obj["commitment"] = commitment
	}
	if offset != nil && length != nil {
		obj["dataSlice"] = M{
			"offset": offset,
			"length": length,
		}
		if encoding == EncodingJSONParsed {
			return nil, errors.New("cannot use dataSlice with EncodingJSONParsed")
		}
	}

	params := []interface{}{account}
	if len(obj) > 0 {
		params = append(params, obj)
	}

	err = cl.rpcClient.CallFor(&out, "getAccountInfo", params)
	if err != nil {
		return nil, err
	}

	if out.Value == nil {
		return nil, ErrNotFound
	}

	return out, nil
}

// GetAccountInfo populates the provided `inVar` parameter with all information associated with the account of provided publicKey.
func (c *Client) GetAccountDataIn(ctx context.Context, account solana.PublicKey, inVar interface{}) (err error) {
	resp, err := c.GetAccountInfo(ctx, account)
	if err != nil {
		return err
	}

	return bin.NewDecoder(resp.Value.Data).Decode(inVar)
}

// GetProgramAccounts returns all accounts owned by the provided program publicKey.
func (c *Client) GetProgramAccounts(
	ctx context.Context,
	publicKey solana.PublicKey,
	opts *GetProgramAccountsOpts,
) (out GetProgramAccountsResult, err error) {
	obj := M{
		"encoding": "base64",
	}
	if opts != nil {
		if opts.Commitment != "" {
			obj["commitment"] = string(opts.Commitment)
		}
		if len(opts.Filters) != 0 {
			obj["filters"] = opts.Filters
		}
		if opts.Encoding != "" {
			// TODO: remove option?
			obj["encoding"] = opts.Encoding
		}
		// if opts.WithContext != nil {
		// 	obj["withContext"] = opts.WithContext
		// }
		if opts.DataSlice != nil {
			obj["dataSlice"] = M{
				"offset": opts.DataSlice.Offset,
				"length": opts.DataSlice.Length,
			}
		}
	}

	params := []interface{}{publicKey, obj}

	err = c.rpcClient.CallFor(&out, "getProgramAccounts", params)
	return
}

// GetMinimumBalanceForRentExemption returns minimum balance required to make account rent exempt.
func (c *Client) GetMinimumBalanceForRentExemption(
	ctx context.Context,
	dataSize int,
	commitment CommitmentType,
) (lamport int, err error) {
	params := []interface{}{dataSize}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = c.rpcClient.CallFor(&lamport, "getMinimumBalanceForRentExemption", params)
	return
}

type SimulateTransactionResponse struct {
	Err  interface{}
	Logs []string
}

// SimulateTransaction simulates sending a transaction.
func (c *Client) SimulateTransaction(ctx context.Context, transaction *solana.Transaction) (*SimulateTransactionResponse, error) {
	buf := new(bytes.Buffer)
	if err := bin.NewEncoder(buf).Encode(transaction); err != nil {
		return nil, fmt.Errorf("send transaction: encode transaction: %w", err)
	}
	trxData := buf.Bytes()

	obj := M{
		"encoding": "base64",
	}

	b64Data := base64.StdEncoding.EncodeToString(trxData)
	params := []interface{}{
		b64Data,
		obj,
	}

	var out *SimulateTransactionResponse
	if err := c.rpcClient.CallFor(&out, "simulateTransaction", params); err != nil {
		return nil, fmt.Errorf("send transaction: rpc send: %w", err)
	}

	return out, nil

}

// SendTransaction submits a signed transaction to the cluster for processing.
func (c *Client) SendTransaction(ctx context.Context, transaction *solana.Transaction) (signature string, err error) {

	buf := new(bytes.Buffer)

	if err := bin.NewEncoder(buf).Encode(transaction); err != nil {
		return "", fmt.Errorf("send transaction: encode transaction: %w", err)
	}

	trxData := buf.Bytes()

	obj := M{
		"encoding": "base64",
	}

	params := []interface{}{
		base64.StdEncoding.EncodeToString(trxData),
		obj,
	}

	if err := c.rpcClient.CallFor(&signature, "sendTransaction", params); err != nil {
		return "", fmt.Errorf("send transaction: rpc send: %w", err)
	}
	return
}

// RequestAirdrop requests an airdrop of lamports to a publicKey.
func (cl *Client) RequestAirdrop(
	ctx context.Context,
	account solana.PublicKey,
	lamport uint64,
	commitment CommitmentType,
) (signature solana.Signature, err error) {
	params := []interface{}{
		account,
		lamport,
	}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&signature, "requestAirdrop", params)
	return
}
