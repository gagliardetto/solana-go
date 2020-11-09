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
