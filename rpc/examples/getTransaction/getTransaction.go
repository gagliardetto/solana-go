// Copyright 2021 github.com/gagliardetto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
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

		decodedTx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(out.Transaction.GetBinary()))
		if err != nil {
			panic(err)
		}
		spew.Dump(decodedTx)
	}
}
