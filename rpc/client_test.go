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
	"fmt"
	"testing"

	"go.uber.org/zap"

	"github.com/dfuse-io/solana-go"
	"github.com/stretchr/testify/require"
)

func TestClient_GetAccountInfo(t *testing.T) {
	c := NewClient("http://api.mainnet-beta.solana.com:80/rpc")
	pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
	accInfo, err := c.GetAccountInfo(context.Background(), pubKey)
	require.NoError(t, err)
	d, err := json.MarshalIndent(accInfo, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))
	pubKey = solana.MustPublicKeyFromBase58("EXnGBBSamqzd3uxEdRLUiYzjJkTwQyorAaFXdfteuGXe")
	accInfo, err = c.GetAccountInfo(context.Background(), pubKey)
	require.NoError(t, err)
	d, err = json.MarshalIndent(accInfo, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))
}

func TestClient_GetConfirmedSignaturesForAddress2(t *testing.T) {
	c := NewClient("http://api.mainnet-beta.solana.com:80/rpc")
	account := solana.MustPublicKeyFromBase58("H7ATJQGhwG8Uf8sUntUognFpsKixPy2buFnXkvyNbGUb")
	accInfo, err := c.GetConfirmedSignaturesForAddress2(context.Background(), account, &GetConfirmedSignaturesForAddress2Opts{Limit: 1})
	require.NoError(t, err)

	d, err := json.MarshalIndent(accInfo, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))

}

func TestClient_GetConfirmedTransaction(t *testing.T) {
	zlog, _ = zap.NewDevelopment()
	c := NewClient("http://api.mainnet-beta.solana.com:80/rpc")
	c.Debug = true
	signature := "53hoZ98EsCMA6L63GWM65M3Bd3WqA4LxD8bcJkbKoKWhbJFqX9M1WZ4fSjt8bYyZn21NwNnV2A25zirBni9Qk6LR"
	trx, err := c.GetConfirmedTransaction(context.Background(), signature)
	require.NoError(t, err)

	d, err := json.MarshalIndent(trx, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))

	signature = "4ZK6ofUodMP8NrB8RGkKFpXWVKMk5eqjkBTbq7DKiDu34gbdrpgctJHp3cU79ZGEBgTaohbjy56KJwhraVmgYq9i"
	trx, err = c.GetConfirmedTransaction(context.Background(), signature)
	require.NoError(t, err)

	d, err = json.MarshalIndent(trx, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))

}

func TestClient_getMinimumBalanceForRentExemption(t *testing.T) {
	zlog, _ = zap.NewDevelopment()
	c := NewClient("http://api.mainnet-beta.solana.com:80/rpc")
	c.Debug = true
	lamport, err := c.GetMinimumBalanceForRentExemption(context.Background(), 100)
	require.NoError(t, err)
	require.Equal(t, 1586880, lamport)
}
