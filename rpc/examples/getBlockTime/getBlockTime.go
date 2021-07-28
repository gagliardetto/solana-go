package main

import (
	"context"
	"time"

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

	out, err := client.GetBlockTime(
		context.TODO(),
		uint64(example.Context.Slot),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	spew.Dump(out.Time().Format(time.RFC1123))
}
