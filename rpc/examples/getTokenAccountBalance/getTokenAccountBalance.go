package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_TestNet
	client := rpc.New(endpoint)

	pubKey := solana.MustPublicKeyFromBase58("EzK5qLWhftu8Z2znVa5fozVtobbjhd8Gdu9hQHpC8bec")
	out, err := client.GetTokenAccountBalance(
		context.TODO(),
		pubKey,
		rpc.CommitmentType("finalized"),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
