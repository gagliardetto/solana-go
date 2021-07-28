package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	{
		out, err := client.GetBlockProduction(context.TODO())
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
	{
		out, err := client.GetBlockProductionWithOpts(
			context.TODO(),
			&rpc.GetBlockProductionOpts{
				Commitment: rpc.CommitmentFinalized,
				// Range: &rpc.SlotRangeRequest{
				// 	FirstSlot: XXXXXX,
				// 	Identity:  solana.MustPublicKeyFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				// },
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}
