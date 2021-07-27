package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_TestNet
	client := rpc.New(endpoint)

	example, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentType("finalized"))
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
				Commitment: rpc.CommitmentType("finalized"),
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
