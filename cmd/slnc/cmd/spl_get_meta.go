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
	"context"
	"os"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/programs/tokenregistry"
	_ "github.com/dfuse-io/solana-go/programs/tokenregistry"
	"github.com/dfuse-io/solana-go/text"
	"github.com/spf13/cobra"
)

var getTokenMetaCmd = &cobra.Command{
	Use:   "meta {account}",
	Short: "Retrieve token meta for a specific account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		address := args[0]
		pubKey, err := solana.PublicKeyFromBase58(address)
		errorCheck("public key", err)

		accountInfo, err := client.GetAccountInfo(ctx, pubKey)

		var tm *tokenregistry.TokenMeta
		err = bin.NewDecoder(accountInfo.Value.Data).Decode(&tm)
		errorCheck("decode", err)

		err = text.NewEncoder(os.Stdout).Encode(tm, nil)
		errorCheck("textEncoding", err)

		//fmt.Println("raw data", hex.EncodeToString(accountInfo.Value.Data))
		return nil
	},
}

func init() {
	splGetCmd.AddCommand(getTokenMetaCmd)
}
