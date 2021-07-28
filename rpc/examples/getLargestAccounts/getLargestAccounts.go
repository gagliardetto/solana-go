package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	out, err := client.GetLargestAccounts(
		context.TODO(),
		rpc.CommitmentFinalized,
		rpc.LargestAccountsFilterCirculating,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
