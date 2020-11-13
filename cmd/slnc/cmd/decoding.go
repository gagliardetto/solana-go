// Copyright 2020 dfuse Platform Inc.
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

package cmd

import (
	"fmt"
	"log"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/token"
)

func decode(owner solana.PublicKey, data []byte) (interface{}, error) {
	bdx, _ := solana.PublicKeyFromBase58("BdxDnkFufu8tjAE5gdPkWdjGfQ3Lz2v6ozfiBDMKxDFW")
	tkn, _ := solana.PublicKeyFromBase58("TokenSVp5gheXUvJ6jGWGeCsgPKgnE3YgdGKRVCMY9o")
	switch owner {
	case bdx, tkn:
		return decodeAsToken(data)
	}
	return nil, nil
}

func decodeAsToken(data []byte) (out interface{}, err error) {

	switch len(data) {
	case 120:
		var tokenAcct token.Account

		if err := bin.NewDecoder(data).Decode(&tokenAcct); err != nil {
			return nil, fmt.Errorf("failed unpacking: %w", err)
		}

		return tokenAcct, nil
	case 40:
		var mint token.Mint
		if err := bin.NewDecoder(data).Decode(&mint); err != nil {
			log.Fatalln("failed unpack", err)
		}

		return mint, nil

		// cnt, _ := json.MarshalIndent(mint, "", "  ")
		// fmt.Println(string(cnt))
	case 7777:
		// decode the Multisig struct
	}
	return
}
