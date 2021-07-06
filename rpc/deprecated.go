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
	"context"

	"github.com/gagliardetto/solana-go"
)

// DEPRECATED: GetConfirmedBlock returns identity and transaction information about a confirmed block in the ledger.
// DEPRECATED: Please use getBlock instead This method is expected to be removed in solana-core v1.8
func (cl *Client) GetConfirmedBlock(
	ctx context.Context,
	slot uint64,
	encoding string,
) (out *GetConfirmedBlockResult, err error) {
	if encoding == "" {
		encoding = "json"
	}
	params := []interface{}{slot, encoding}

	err = cl.rpcClient.CallFor(&out, "getConfirmedBlock", params)
	return
}

// DEPRECATED: GetConfirmedTransaction returns transaction details for a confirmed transaction.
func (cl *Client) GetConfirmedTransaction(
	ctx context.Context,
	signature string,
) (out TransactionWithMeta, err error) {
	params := []interface{}{signature, "json"}

	err = cl.rpcClient.CallFor(&out, "getConfirmedTransaction", params)
	return
}

// DEPRECATED: GetConfirmedSignaturesForAddress2 returns confirmed signatures for transactions involving an
// address backwards in time from the provided signature or most recent confirmed block.
func (cl *Client) GetConfirmedSignaturesForAddress2(ctx context.Context, address solana.PublicKey, opts *GetConfirmedSignaturesForAddress2Opts) (out GetConfirmedSignaturesForAddress2Result, err error) {

	params := []interface{}{address.String(), opts}

	err = cl.rpcClient.CallFor(&out, "getConfirmedSignaturesForAddress2", params)
	return
}
