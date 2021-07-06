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
	"encoding/json"
	"testing"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetAccountInfo(t *testing.T) {
	server, closer := mockJSONRPC(t, json.RawMessage(`{"jsonrpc":"2.0","result":{"context":{"slot":1},"value":{"data":["dGVzdA==","base64"]}},"id":0}`))
	defer closer()
	client := NewClient(server.URL)

	pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
	out, err := client.GetAccountInfo(context.Background(), pubKey)
	require.NoError(t, err)

	assert.Equal(t, map[string]interface{}{"id": float64(0), "jsonrpc": "2.0", "method": "getAccountInfo", "params": []interface{}{
		"7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932",
		map[string]interface{}{"encoding": "base64"},
	}}, server.RequestBody(t))

	assert.Equal(t, &GetAccountInfoResult{RPCContext: RPCContext{Context{Slot: 1}}, Value: &Account{Data: []byte{0x74, 0x65, 0x73, 0x74}}}, out)
}

func TestClient_GetConfirmedSignaturesForAddress2(t *testing.T) {
	server, closer := mockJSONRPC(t, json.RawMessage(`{"jsonrpc":"2.0","result":[{"err":null,"memo":null,"signature":"mgw5vw4tnbou1wVStKckVcVncbpRwfZPcMNbVBoigbSPXBMa3857CNzhwoCkRzM5K7nG32wcbpVJDHttQeBRaHB","slot":1}],"id":0}`))
	defer closer()
	client := NewClient(server.URL)

	account := solana.MustPublicKeyFromBase58("H7ATJQGhwG8Uf8sUntUognFpsKixPy2buFnXkvyNbGUb")
	out, err := client.GetConfirmedSignaturesForAddress2(context.Background(), account, &GetConfirmedSignaturesForAddress2Opts{Limit: 1})
	require.NoError(t, err)

	assert.Equal(t, map[string]interface{}{"id": float64(0), "jsonrpc": "2.0", "method": "getConfirmedSignaturesForAddress2", "params": []interface{}{
		"H7ATJQGhwG8Uf8sUntUognFpsKixPy2buFnXkvyNbGUb",
		map[string]interface{}{"limit": float64(1)},
	}}, server.RequestBody(t))

	expected := []*TransactionSignature{
		{Slot: 1, Signature: "mgw5vw4tnbou1wVStKckVcVncbpRwfZPcMNbVBoigbSPXBMa3857CNzhwoCkRzM5K7nG32wcbpVJDHttQeBRaHB"},
	}

	assert.Equal(t, GetConfirmedSignaturesForAddress2Result(expected), out)
}

func TestClient_GetConfirmedTransaction(t *testing.T) {
	server, closer := mockJSONRPC(t, json.RawMessage(`{"jsonrpc":"2.0","result":{"meta":{"err":null,"fee":5000,"innerInstructions":[],"logMessages":[],"postBalances":[],"preBalances":[],"status":{"Ok":null}},"slot":48291656,"transaction":{"message":{"accountKeys":["GKu2xfGZopa8C9K11wduQWgP4W4H7EEcaNdsUb7mxhyr"],"header":{"numReadonlySignedAccounts":0,"numReadonlyUnsignedAccounts":3,"numRequiredSignatures":1},"instructions":[{"accounts":[1,2,3,0],"data":"3yZe7d","programIdIndex":4}],"recentBlockhash":"uoEAQCWCKjV9ecsBvngctJ7upNBZX7hpN4SfdR6TaUz"},"signatures":["53hoZ98EsCMA6L63GWM65M3Bd3WqA4LxD8bcJkbKoKWhbJFqX9M1WZ4fSjt8bYyZn21NwNnV2A25zirBni9Qk6LR"]}},"id":0}`))
	defer closer()
	client := NewClient(server.URL)

	out, err := client.GetConfirmedTransaction(context.Background(), "53hoZ98EsCMA6L63GWM65M3Bd3WqA4LxD8bcJkbKoKWhbJFqX9M1WZ4fSjt8bYyZn21NwNnV2A25zirBni9Qk6LR")
	require.NoError(t, err)

	assert.Equal(t, map[string]interface{}{"id": float64(0), "jsonrpc": "2.0", "method": "getConfirmedTransaction", "params": []interface{}{
		"53hoZ98EsCMA6L63GWM65M3Bd3WqA4LxD8bcJkbKoKWhbJFqX9M1WZ4fSjt8bYyZn21NwNnV2A25zirBni9Qk6LR",
		"json",
	}}, server.RequestBody(t))

	signature, err := solana.SignatureFromBase58("53hoZ98EsCMA6L63GWM65M3Bd3WqA4LxD8bcJkbKoKWhbJFqX9M1WZ4fSjt8bYyZn21NwNnV2A25zirBni9Qk6LR")
	require.NoError(t, err)

	assert.Equal(t, TransactionWithMeta{
		Transaction: &solana.Transaction{
			Message: solana.Message{
				Header:          solana.MessageHeader{NumRequiredSignatures: 1, NumReadonlySignedAccounts: 0, NumReadonlyUnsignedAccounts: 3},
				RecentBlockhash: solana.MustPublicKeyFromBase58("uoEAQCWCKjV9ecsBvngctJ7upNBZX7hpN4SfdR6TaUz"),
				AccountKeys:     []solana.PublicKey{solana.MustPublicKeyFromBase58("GKu2xfGZopa8C9K11wduQWgP4W4H7EEcaNdsUb7mxhyr")},
				Instructions: []solana.CompiledInstruction{
					{Accounts: []uint8{1, 2, 3, 0}, Data: solana.Base58([]byte{0x74, 0x65, 0x73, 0x74}), ProgramIDIndex: 4},
				},
			},
			Signatures: []solana.Signature{signature},
		},
		Meta: &TransactionMeta{Fee: 5000, PreBalances: []bin.Uint64{}, PostBalances: []bin.Uint64{}},
	}, out)
}

func TestClient_getMinimumBalanceForRentExemption(t *testing.T) {
	server, closer := mockJSONRPC(t, json.RawMessage(`{"jsonrpc":"2.0","result":1586880,"id":0}`))
	defer closer()
	client := NewClient(server.URL)

	out, err := client.GetMinimumBalanceForRentExemption(context.Background(), 100)
	require.NoError(t, err)

	assert.Equal(t, map[string]interface{}{"id": float64(0), "jsonrpc": "2.0", "method": "getMinimumBalanceForRentExemption", "params": []interface{}{
		float64(100),
	}}, server.RequestBody(t))

	assert.Equal(t, int(1586880), out)

}
