// Copyright 2021 github.com/gagliardetto
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

	"github.com/gagliardetto/solana-go"
)

// GetSupply returns information about the current supply.
func (cl *Client) GetSupply(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out *GetSupplyResult, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getSupply", params)
	return
}

type GetSupplyResult struct {
	RPCContext
	Value *SupplyResult `json:"value"`
}

type SupplyResult struct {
	// Total supply in lamports
	Total uint64 `json:"total"`

	// Circulating supply in lamports.
	Circulating uint64 `json:"circulating"`

	// Non-circulating supply in lamports.
	NonCirculating uint64 `json:"nonCirculating"`

	// An array of account addresses of non-circulating accounts.
	NonCirculatingAccounts []solana.PublicKey `json:"nonCirculatingAccounts"`
}
