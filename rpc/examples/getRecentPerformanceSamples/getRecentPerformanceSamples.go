package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_TestNet
	client := rpc.New(endpoint)

	limit := uint(3)
	out, err := client.GetRecentPerformanceSamples(
		context.TODO(),
		&limit,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
