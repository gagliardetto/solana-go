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
	"fmt"
	"os"

	"github.com/dfuse-io/solana-go"
	_ "github.com/dfuse-io/solana-go/programs/serum"
	_ "github.com/dfuse-io/solana-go/programs/system"
	_ "github.com/dfuse-io/solana-go/programs/token"
	_ "github.com/dfuse-io/solana-go/programs/tokenregistry"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/dfuse-io/solana-go/text"
	"github.com/spf13/cobra"
)

var getTransactionsCmd = &cobra.Command{
	Use:   "transactions {account}",
	Short: "Retrieve transaction for a specific account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		address := args[0]
		pubKey, err := solana.PublicKeyFromBase58(address)
		errorCheck("public key", err)

		csList, err := client.GetConfirmedSignaturesForAddress2(ctx, pubKey, &rpc.GetConfirmedSignaturesForAddress2Opts{
			Limit:  1,
			Before: "",
			Until:  "",
		})

		errorCheck("getting confirm transaction:", err)

		for _, cs := range csList {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			text.EncoderColorCyan.Print("Transaction: ")
			fmt.Println(cs.Signature)

			text.EncoderColorGreen.Print("Slot: ")
			fmt.Println(cs.Slot)
			text.EncoderColorGreen.Print("Memo: ")
			fmt.Println(cs.Memo)

			ct, err := client.GetConfirmedTransaction(ctx, cs.Signature)
			errorCheck("confirm transaction", err)
			if ct.Meta.Err != nil {
				fmt.Println("ERROR:", ct.Meta.Err)
				//	for k, _ := range ct.Meta.Err
			}

			fmt.Print("\nInstructions:\n-------------\n\n")
			for _, i := range ct.Transaction.Message.Instructions {

				//Missing Initial account instruction ??????

				id, err := ct.Transaction.ResolveProgramIDIndex(i.ProgramIDIndex)
				errorCheck("resolving programID", err)
				decoder := solana.InstructionDecoderRegistry[id.String()]
				if decoder == nil {
					fmt.Println("raw instruction:")
					fmt.Printf("Program: %s Data: %s\n", id.String(), i.Data)
					fmt.Println("Accounts:")
					for _, accIndex := range i.Accounts {
						key := ct.Transaction.Message.AccountKeys[accIndex]

						fmt.Printf("%s Is Writable: %t Is Signer: %t\n", key.String(), ct.Transaction.IsWritable(key), ct.Transaction.IsSigner(key))
					}
					fmt.Printf("\n\n")
					continue
				}

				decoded, err := decoder(ct.Transaction.AccountMetaList(), &i)
				errorCheck("bin decode", err)
				err = text.NewEncoder(os.Stdout).Encode(decoded, nil)
				errorCheck("textEncoding", err)
			}
			text.EncoderColorCyan.Print("\n\nEnd of transaction\n\n")
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getTransactionsCmd)
}
