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

	pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
	// Let's get a valid transaction to use in the example:
	example, err := client.GetConfirmedSignaturesForAddress2(
		context.TODO(),
		pubKey,
		nil,
	)
	if err != nil {
		panic(err)
	}

	out, err := client.GetConfirmedTransaction(
		context.TODO(),
		example[0].Signature,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
