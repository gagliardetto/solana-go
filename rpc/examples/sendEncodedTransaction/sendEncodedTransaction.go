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
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)
	base64Tx := "AepZ357SXFu9tOBeX8v8aqkPNyGq/RUN4EvDzWAT/Fsg1htQu0Zp6F6OxO645I5z98VI4XacPKJp8rtHOE7Q9Q0BAAED8gJOoYf0IvSirfhOq6ZFf3R3ekB5vFWlaGTEq3irvwFakK+tD9il+9jzzs+gU1wzZxkmXZqyeBhbaXogNlk1GQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAou6RU2rKBCg2RzcJqZwEr4y2VBsWbkYLYMKFcOz64j4BAgIAAQwCAAAAAOH1BQAAAAA="

	sig, err := client.SendEncodedTransaction(context.TODO(), base64Tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Submitted tx signature: ", sig.String())
}
