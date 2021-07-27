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

	txSig := solana.MustSignatureFromBase58("4bjVLV1g9SAfv7BSAdNnuSPRbSscADHFe4HegL6YVcuEBMY83edLEvtfjE4jfr6rwdLwKBQbaFiGgoLGtVicDzHq")
	{
		out, err := client.GetTransaction(
			context.TODO(),
			txSig,
			nil,
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
		spew.Dump(out.Transaction.GetParsedTransaction())
	}
	{
		out, err := client.GetTransaction(
			context.TODO(),
			txSig,
			&rpc.GetTransactionOpts{
				Encoding: solana.EncodingJSON,
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
		spew.Dump(out.Transaction.GetParsedTransaction())
	}
	{
		out, err := client.GetTransaction(
			context.TODO(),
			txSig,
			&rpc.GetTransactionOpts{
				Encoding: solana.EncodingBase58,
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
		spew.Dump(out.Transaction.GetBinary())
	}
	{
		out, err := client.GetTransaction(
			context.TODO(),
			txSig,
			&rpc.GetTransactionOpts{
				Encoding: solana.EncodingBase64,
			},
		)
		if err != nil {
			panic(err)
		}
		spew.Dump(out)
		spew.Dump(out.Transaction.GetBinary())
	}
}
