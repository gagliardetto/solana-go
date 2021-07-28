package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	recent, err := client.GetRecentBlockhash(
		context.TODO(),
		rpc.CommitmentFinalized,
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
