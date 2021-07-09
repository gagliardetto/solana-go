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

	bin "github.com/dfuse-io/binary"
)

// GetSlot returns the current slot the node is processing.
func (cl *Client) GetSlot(ctx context.Context, commitment CommitmentType) (out bin.Uint64, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}

	err = cl.rpcClient.CallFor(&out, "getSlot", params)
	return
}
