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

package ws

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_AccountSubscribe(t *testing.T) {
	t.Skip("Never ending test, revisit me to not depend on actual network calls, or hide between env flag")

	zlog, _ = zap.NewDevelopment()

	c, err := Dial(context.Background(), "ws://api.mainnet-beta.solana.com:80")
	defer c.Close()
	require.NoError(t, err)

	accountID := solana.MustPublicKeyFromBase58("SqJP6vrvMad5XBQK5PCFEZjeuQSFi959sdpqtSNvnsX")
	sub, err := c.AccountSubscribe(accountID, "")
	require.NoError(t, err)

	data, err := sub.Recv()
	if err != nil {
		fmt.Println("receive an error: ", err)
		return
	}
	text.NewEncoder(os.Stdout).Encode(data, nil)
	fmt.Println("OpenOrders: ", data.Value.Account.Owner)
	fmt.Println("data: ", data.Value.Account.Data)
	return
}

func Test_ProgramSubscribe(t *testing.T) {
	t.Skip("Never ending test, revisit me to not depend on actual network calls, or hide between env flag")

	zlog, _ = zap.NewDevelopment()

	fmt.Println("Dialing")
	c, err := Dial(context.Background(), "wss://solana-api.projectserum.com")
	fmt.Println("Hello?")
	defer c.Close()
	require.NoError(t, err)

	programID := solana.MustPublicKeyFromBase58("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o")
	sub, err := c.ProgramSubscribe(programID, "")
	require.NoError(t, err)

	for {
		data, err := sub.Recv()
		if err != nil {
			fmt.Println("receive an error: ", err)
			return
		}
		fmt.Println("data received: ", data.Value.Pubkey)
	}

}
func Test_SlotSubscribe(t *testing.T) {
	t.Skip("Never ending test, revisit me to not depend on actual network calls, or hide between env flag")

	zlog, _ = zap.NewDevelopment()

	c, err := Dial(context.Background(), "ws://api.mainnet-beta.solana.com:80")
	defer c.Close()
	require.NoError(t, err)

	sub, err := c.SlotSubscribe()
	require.NoError(t, err)

	data, err := sub.Recv()
	if err != nil {
		fmt.Println("receive an error: ", err)
		return
	}
	fmt.Println("data received: ", data.Parent)
	return
}
