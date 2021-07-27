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

	out, err := client.GetSignatureStatuses(
		context.TODO(),
		true,
		// All the transactions you want the get the status for:
		solana.MustSignatureFromBase58("2CwH8SqVZWFa1EvsH7vJXGFors1NdCuWJ7Z85F8YqjCLQ2RuSHQyeGKkfo1Tj9HitSTeLoMWnxpjxF2WsCH8nGWh"),
		solana.MustSignatureFromBase58("5YJHZPeHZuZjhunBc1CCB1NDRNf2tTJNpdb3azGsR7PfyEncCDhr95wG8EWrvjNXBc4wCKixkheSbCxoC2NCG3X7"),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
