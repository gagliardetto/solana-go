package main

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_TestNet
	client := rpc.New(endpoint)

	example, err := client.GetRecentBlockhash(
		context.TODO(),
		rpc.CommitmentType("finalized"),
	)
	if err != nil {
		panic(err)
	}

	{ // deprecated and is going to be removed in solana-core v1.8
		out, err := client.GetConfirmedBlock(
			context.TODO(),
			uint64(example.Context.Slot),
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
	{
		slot := uint64(example.Context.Slot)
		out, err := client.GetConfirmedBlockWithOpts(
			context.TODO(),
			slot,
			// You can specify more options here:
			&rpc.GetConfirmedBlockOpts{
				Encoding:   solana.EncodingBase64,
				Commitment: rpc.CommitmentType("finalized"),
				// Get only signatures:
				TransactionDetails: rpc.TransactionDetailsSignatures,
				// Exclude rewards:
				Rewards: pointer.ToBool(false),
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}
