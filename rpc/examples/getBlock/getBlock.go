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

package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	{
		out, err := client.GetBlock(context.TODO(), uint64(example.Context.Slot))
		if err != nil {
			panic(err)
		}
		// spew.Dump(out) // NOTE: This generates a lot of output.
		spew.Dump(len(out.Transactions))
	}

	{
		includeRewards := false
		out, err := client.GetBlockWithOpts(
			context.TODO(),
			uint64(example.Context.Slot),
			// You can specify more options here:
			&rpc.GetBlockOpts{
				Encoding:   solana.EncodingBase64,
				Commitment: rpc.CommitmentFinalized,
				// Get only signatures:
				TransactionDetails: rpc.TransactionDetailsSignatures,
				// Exclude rewards:
				Rewards: &includeRewards,
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}
