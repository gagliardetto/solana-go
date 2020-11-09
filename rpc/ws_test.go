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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWSClient_ProgramSubscribe(t *testing.T) {

	c, err := NewWSClient("ws://api.mainnet-beta.solana.com:80/rpc")
	require.NoError(t, err)

	stream, _, err := c.ProgramSubscribe("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o", "")
	require.NoError(t, err)

	select {
	case <-stream:
	case <-time.After(2000 * time.Millisecond):
		t.Error("failed to run the giving time")
	}
}
