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

	pubKey := solana.MustPublicKeyFromBase58("6dmNQ5jwLeLk5REvio1JcMshcbvkYMwy26sJ8pbkvStu")

	out, err := client.GetInflationReward(
		context.TODO(),
		[]solana.PublicKey{
			pubKey,
		},
		&rpc.GetInflationRewardOpts{
			Commitment: rpc.CommitmentType("finalized"),
		},
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
