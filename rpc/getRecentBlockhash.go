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
)

// GetRecentBlockhash returns a recent block hash from the ledger,
// and a fee schedule that can be used to compute the cost of submitting a transaction using it.
func (cl *Client) GetRecentBlockhash(ctx context.Context, commitment CommitmentType) (out *GetRecentBlockhashResult, err error) {
	var params []interface{}
	if commitment != "" {
		commit := map[string]string{
			"commitment": string(commitment),
		}
		params = append(params, commit)
	}

	err = cl.rpcClient.CallFor(&out, "getRecentBlockhash", params)
	return
}
