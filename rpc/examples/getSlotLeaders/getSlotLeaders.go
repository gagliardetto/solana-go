package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_TestNet
	client := rpc.New(endpoint)

	recent, err := client.GetRecentBlockhash(
		context.TODO(),
		rpc.CommitmentType("finalized"),
	)
	if err != nil {
		panic(err)
	}

	out, err := client.GetSlotLeaders(
		context.TODO(),
		uint64(recent.Context.Slot),
		10,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
