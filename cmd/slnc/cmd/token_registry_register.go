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

	"github.com/spf13/viper"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/programs/system"
	"github.com/dfuse-io/solana-go/programs/tokenregistry"
	"github.com/spf13/cobra"
)

var tokenRegistryRegisterCmd = &cobra.Command{
	Use:   "register {token} {name} {symbol} {logo}",
	Short: "register meta data for a token",
	Args:  cobra.ExactArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		vault := mustGetWallet()
		client := getClient()

		tokenAddress, err := solana.PublicKeyFromBase58(args[0])
		errorCheck("invalid token address", err)

		logo, err := tokenregistry.LogoFromString(args[1])
		errorCheck("invalid logo", err)
		name, err := tokenregistry.NameFromString(args[2])
		errorCheck("invalid name", err)
		symbol, err := tokenregistry.SymbolFromString(args[3])
		errorCheck("invalid symbol", err)

		pkeyStr := viper.GetString("token-registry-register-cmd-registrar")
		if pkeyStr == "" {
			fmt.Errorf("unable to continue without a specified registrar")
		}

		registrarPubKey, err := solana.PublicKeyFromBase58(pkeyStr)
		errorCheck(fmt.Sprintf("invalid registrar key %q", pkeyStr), err)

		found := false
		for _, privateKey := range vault.KeyBag {
			if privateKey.PublicKey() == registrarPubKey {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("registrar key must be present in the vault to register a token")
		}

		fmt.Println(registrarPubKey.String())

		tokenMetaAccount := solana.NewAccount()

		lamport, err := client.GetMinimumBalanceForRentExemption(context.Background(), tokenregistry.TOKEN_META_SIZE)
		errorCheck("get minimum balance for rent exemption ", err)

		tokenRegistryProgramID := tokenregistry.ProgramID()

		createAccountInstruction := system.NewCreateAccountInstruction(uint64(lamport), tokenregistry.TOKEN_META_SIZE, tokenRegistryProgramID, registrarPubKey, tokenMetaAccount.PublicKey())
		registerTokenInstruction := tokenregistry.NewRegisterTokenInstruction(logo, name, symbol, tokenMetaAccount.PublicKey(), registrarPubKey, tokenAddress)

		trx, err := solana.TransactionWithInstructions([]solana.TransactionInstruction{createAccountInstruction, registerTokenInstruction}, &solana.Options{
			Payer: registrarPubKey,
		})
		errorCheck("unable to craft transaction", err)

		_, err = trx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
			for _, k := range vault.KeyBag {
				if k.PublicKey() == key {
					return &k
				}
			}
			return nil
		})

		errorCheck("unable to sign transaction", err)

		trxHash, err := client.SendTransaction(cmd.Context(), trx)

		fmt.Println("sent transaction hash:", trxHash, " error:", err)
		return nil
	},
}

func init() {
	tokenRegistryCmd.AddCommand(tokenRegistryRegisterCmd)
	tokenRegistryRegisterCmd.PersistentFlags().String("registrar", "9hFtYBYmBJCVguRYs9pBTWKYAFoKfjYR7zBPpEkVsmD", "The public key that will register the token")
}
