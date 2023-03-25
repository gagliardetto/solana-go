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

// TODO: add comments
func (cl *Client) GetRecentPrioritizationFees(
	ctx context.Context,
	accounts solana.PublicKeySlice, // optional
) (out []PriorizationFeeResult, err error) {
	params := []interface{}{accounts}
	err = cl.rpcClient.CallForInto(ctx, &out, "getRecentPrioritizationFees", params)

	return
}

type PriorizationFeeResult struct {
	Slot              uint64 `json:"slot"`
	PrioritizationFee uint64 `json:"prioritizationFee"`
}
