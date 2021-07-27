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

	pubKey := solana.MustPublicKeyFromBase58("EW2p7QCJNHMVj5nQCcW7Q2BDETtNBXn68FyucU4RCjvb")
	out, err := client.GetStakeActivation(
		context.TODO(),
		pubKey,
		rpc.CommitmentType("finalized"),
		nil,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
