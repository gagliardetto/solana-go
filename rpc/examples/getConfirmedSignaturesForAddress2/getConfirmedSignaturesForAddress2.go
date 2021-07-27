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
	{
		// deprecated and is going to be removed in solana-core v1.8
		out, err := client.GetConfirmedSignaturesForAddress2(
			context.TODO(),
			pubKey,
			// TODO:
			nil,
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}
