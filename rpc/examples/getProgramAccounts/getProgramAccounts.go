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

	out, err := client.GetProgramAccounts(
		context.TODO(),
		solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(len(out))
	spew.Dump(out) // NOTE: this can generate a lot of output
}
