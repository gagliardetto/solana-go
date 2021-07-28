package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetRecentBlockhash(
		context.TODO(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}

	{
		endSlot := uint64(example.Context.Slot)
		// deprecated and is going to be removed in solana-core v1.8
		out, err := client.GetConfirmedBlocks(
			context.TODO(),
			uint64(example.Context.Slot-3),
			&endSlot,
			rpc.CommitmentFinalized,
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}
