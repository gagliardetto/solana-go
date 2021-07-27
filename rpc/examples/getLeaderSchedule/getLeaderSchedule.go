package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_TestNet
	client := rpc.New(endpoint)

	out, err := client.GetLeaderSchedule(
		context.TODO(),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out) // NOTE: this creates a lot of output
}
