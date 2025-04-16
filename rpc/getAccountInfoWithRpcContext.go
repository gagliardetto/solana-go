// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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

// GetAccountInfoWithRpcContext is similar to GetAccountInfoWithOpts but will return rpcContext and nil account if account is not found
func (cl *Client) GetAccountInfoWithRpcContext(
	ctx context.Context,
	account solana.PublicKey,
	opts *GetAccountInfoOpts,
) (*Account, *RPCContext, error) {
	out, err := cl.getAccountInfoWithOpts(ctx, account, opts)
	if err != nil {
		return nil, nil, err
	}
	if out == nil {
		return nil, nil, nil
	} else {
		return out.Value, &out.RPCContext, nil
	}
}
