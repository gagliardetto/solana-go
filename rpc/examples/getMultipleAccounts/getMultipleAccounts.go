package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_MainNetBeta
	client := rpc.New(endpoint)

	{
		out, err := client.GetMultipleAccounts(
			context.TODO(),
			solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"),  // serum token
			solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R"), // raydium token
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
	{
		out, err := client.GetMultipleAccountsWithOpts(
			context.TODO(),
			[]solana.PublicKey{solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"), // serum token
				solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R"), // raydium token
			},
			&rpc.GetMultipleAccountsOpts{
				Encoding:   solana.EncodingBase64Zstd,
				Commitment: rpc.CommitmentType("finalized"),
				// You can get just a part of the account data by specify a DataSlice:
				// DataSlice: &rpc.DataSlice{
				// 	Offset: pointer.ToUint64(0),
				// 	Length: pointer.ToUint64(1024),
				// },
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
	}
}
