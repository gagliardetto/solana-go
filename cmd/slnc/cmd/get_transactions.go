// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
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

	"github.com/gagliardetto/solana-go"
	_ "github.com/gagliardetto/solana-go/programs/serum"
	_ "github.com/gagliardetto/solana-go/programs/system"
	_ "github.com/gagliardetto/solana-go/programs/token"
	_ "github.com/gagliardetto/solana-go/programs/tokenregistry"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/text"
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
		if err != nil {
			return fmt.Errorf("invalid account address %q: %w", address, err)
		}

		limit := uint64(1)
		csList, err := client.GetConfirmedSignaturesForAddress2(ctx, pubKey, &rpc.GetConfirmedSignaturesForAddress2Opts{
			Limit: &limit,
			// Before: "",
			// Until:  "",
		})
		if err != nil {
			return fmt.Errorf("unable to retrieve confirmed transaction signatures for account: %w", err)
		}

		for _, cs := range csList {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			text.EncoderColorCyan.Print("Transaction: ")
			fmt.Println(cs.Signature)

			text.EncoderColorGreen.Print("Slot: ")
			fmt.Println(cs.Slot)
			text.EncoderColorGreen.Print("Memo: ")
			fmt.Println(cs.Memo)

			ct, err := client.GetConfirmedTransaction(ctx, cs.Signature)
			if err != nil {
				return fmt.Errorf("unable to get confirmed transaction with signature %q: %w", cs.Signature, err)
			}

			if ct.Meta.Err != nil {
				return fmt.Errorf("unable to get confirmed transaction with signature %q: %s", cs.Signature, ct.Meta.Err)
			}

			_, err = ct.MustGetTransaction().EncodeTree(text.NewTreeEncoder(os.Stdout, text.Bold("INSTRUCTIONS")))
			if err != nil {
				panic(err)
			}
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getTransactionsCmd)
}
