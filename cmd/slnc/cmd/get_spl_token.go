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
	"os"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/programs/token"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/dfuse-io/solana-go/text"
	"github.com/spf13/cobra"
)

var getSPLTokenCmd = &cobra.Command{
	Use:   "spl-token",
	Short: "Retrieve and decide spl token",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := cmd.Context()

		resp, err := client.GetProgramAccounts(
			ctx,
			solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
			&rpc.GetProgramAccountsOpts{
				Filters: []rpc.RPCFilter{
					{
						DataSize: 82,
					},
				},
			},
		)
		if err != nil {
			return err
		}

		if resp == nil {
			return fmt.Errorf("program account not found")
		}

		for _, keyedAcct := range resp {
			acct := keyedAcct.Account
			//fmt.Println("Data len:", len(acct.Data), keyedAcct.Pubkey)
			var mint *token.Mint
			if err := bin.NewDecoder(acct.Data).Decode(&mint); err != nil {
				log.Fatalln("failed unpack", err)
			}

			text.EncoderColorCyan.Print("Address: ")
			fmt.Println(keyedAcct.Pubkey.String())

			text.EncoderColorCyan.Print("Owner: ")
			fmt.Println(keyedAcct.Account.Owner.String())

			text.EncoderColorCyan.Print("Lamports: ")
			fmt.Println(keyedAcct.Account.Lamports)

			if err := text.NewEncoder(os.Stdout).Encode(mint, nil); err != nil {
				log.Fatalln("failed string encode", err)
			}
			fmt.Println("-------------------------------")
			fmt.Println("")
		}
		fmt.Println("\nTotal result:", len(resp))

		return nil
	},
}

func init() {
	getCmd.AddCommand(getSPLTokenCmd)
}
