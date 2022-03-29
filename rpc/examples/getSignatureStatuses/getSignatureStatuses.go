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
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
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
