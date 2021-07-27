package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_MainNetBeta
	client := rpc.New(endpoint)

	{
		pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
		// basic usage
		resp, err := client.GetAccountInfo(
			context.TODO(),
			pubKey,
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(resp)

		var mint token.Mint
		err = bin.NewDecoder(resp.Value.Data.GetBinary()).Decode(&mint)
		if err != nil {
			panic(err)
		}
		spew.Dump(mint)
		// NOTE: The supply is mint.Supply, with the mint.Decimals:
		// mint.Supply = 9998022451607088
		// mint.Decimals = 6
		// ... which means that the supply is 9998022451.607088
	}
	{
		// Or you can use `GetAccountDataIn` which does all of the above in one call:
		pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
		var mint token.Mint
		// Get the account, and decode its data into the provided mint object:
		err := client.GetAccountDataIn(
			context.TODO(),
			pubKey,
			&mint,
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(mint)
	}
	{
		pubKey := solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R") // raydium token
		// advanced usage
		resp, err := client.GetAccountInfoWithOpts(
			context.TODO(),
			pubKey,
			// You can specify more options here:
			&rpc.GetAccountInfoOpts{
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
		spew.Dump(resp)

		var mint token.Mint
		err = bin.NewDecoder(resp.Value.Data.GetBinary()).Decode(&mint)
		if err != nil {
			panic(err)
		}
		spew.Dump(mint)
	}
}
