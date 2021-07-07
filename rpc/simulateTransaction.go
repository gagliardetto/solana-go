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
	"fmt"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type SimulateTransactionResponse struct {
	Err  interface{} `json:"err,omitempty"`
	Logs []string    `json:"logs,omitempty"`
}

// SimulateTransaction simulates sending a transaction.
func (cl *Client) SimulateTransaction(
	ctx context.Context,
	transaction *solana.Transaction,
) (out *SimulateTransactionResponse, err error) {
	return cl.SimulateTransactionWithOpts(
		ctx,
		transaction,
		false,
		"",
		false,
	)
}

// SimulateTransaction simulates sending a transaction.
func (cl *Client) SimulateTransactionWithOpts(
	ctx context.Context,
	transaction *solana.Transaction,
	sigVerify bool, // if true the transaction signatures will be verified (default: false, conflicts with replaceRecentBlockhash)
	commitment CommitmentType, // Commitment level to simulate the transaction at (default: "finalized").
	replaceRecentBlockhash bool, // if true the transaction recent blockhash will be replaced with the most recent blockhash. (default: false, conflicts with sigVerify)
) (out *SimulateTransactionResponse, err error) {
	buf := new(bytes.Buffer)
	if err := bin.NewEncoder(buf).Encode(transaction); err != nil {
		return nil, fmt.Errorf("send transaction: encode transaction: %w", err)
	}
	trxData := buf.Bytes()

	obj := M{
		"encoding": "base64",
	}
	if sigVerify {
		obj["sigVerify"] = sigVerify
	}
	if commitment != "" {
		obj["commitment"] = commitment
	}
	if replaceRecentBlockhash {
		obj["replaceRecentBlockhash"] = replaceRecentBlockhash
	}

	b64Data := base64.StdEncoding.EncodeToString(trxData)
	params := []interface{}{
		b64Data,
		obj,
	}

	err = cl.rpcClient.CallFor(&out, "simulateTransaction", params)
	return
}
