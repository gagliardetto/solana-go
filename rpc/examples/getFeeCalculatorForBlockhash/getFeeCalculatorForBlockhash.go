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

	out, err := client.GetFeeCalculatorForBlockhash(
		context.TODO(),
		example.Value.Blockhash,
		rpc.CommitmentType("finalized"),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
