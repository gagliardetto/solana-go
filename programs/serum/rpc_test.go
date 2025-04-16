// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
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

package serum

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gagliardetto/solana-go/rpc/ws"

	"github.com/gagliardetto/solana-go/rpc"

	"github.com/stretchr/testify/require"

	"github.com/gagliardetto/solana-go"
)

func TestFetchMarket(t *testing.T) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("Setup 'RPC_URL' to run test i.e. 'https://api.mainnet-beta.solana.com'")
		return
	}

	//

	client := rpc.New(rpcURL)
	ctx := context.Background()

	openOrderAdd, err := solana.PublicKeyFromBase58("jFoHUkNDC767PyK11cZM4zyNcpjLqFnSjaqEYp5GVBr")
	require.NoError(t, err)

	openOrders, err := FetchOpenOrders(ctx, client, openOrderAdd)
	require.NoError(t, err)

	cnt, err := json.MarshalIndent(openOrders.OpenOrders, "", " ")

	require.NoError(t, err)

	fmt.Println(string(cnt))
}

func TestStreamOpenOrders(t *testing.T) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("Setup 'RPC_URL' to run test i.e. 'wss://api.mainnet-beta.solana.com'")
		return
	}
	client, err := ws.Connect(context.Background(), rpcURL)
	require.NoError(t, err)

	err = StreamOpenOrders(context.Background(), client)
	require.NoError(t, err)
}
