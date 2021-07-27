package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	out, err := client.GetHealth(
		context.TODO(),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	spew.Dump(out == rpc.HealthOk)
}
