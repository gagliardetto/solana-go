package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	amount := solana.LAMPORTS_PER_SOL // 1 sol
	pubKey := solana.MustPublicKeyFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	out, err := client.RequestAirdrop(
		context.TODO(),
		pubKey,
		amount,
		"",
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
