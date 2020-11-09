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
	"fmt"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/require"
)

func TestWSClient_ProgramSubscribe(t *testing.T) {
	zlog, _ = zap.NewDevelopment()

	c, err := Dial(context.Background(), "ws://api.mainnet-beta.solana.com:80/rpc")
	defer c.Close()
	require.NoError(t, err)

	sub, err := c.ProgramSubscribe("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o", "")
	require.NoError(t, err)

	data, err := sub.Recv()
	if err != nil {
		fmt.Println("receive an error: ", err)
		return
	}
	fmt.Println("data received: ", data.(*ProgramWSResult).Value.Account.Owner)
	return

}
