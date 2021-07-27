package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
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

	limit := uint64(4)
	out, err := client.GetBlocksWithLimit(
		context.TODO(),
		uint64(example.Context.Slot-10),
		limit,
		rpc.CommitmentType("finalized"),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
