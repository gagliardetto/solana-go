package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_TestNet
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
				Commitment: rpc.CommitmentType("finalized"),
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
