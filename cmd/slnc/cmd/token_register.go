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

	bin "github.com/dfuse-io/binary"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/programs/system"
	"github.com/dfuse-io/solana-go/programs/tokenregistry"
	"github.com/spf13/cobra"
)

var tokenRegisterCmd = &cobra.Command{
	Use:   "token register {token} {name} {symbol} {logo}",
	Short: "register meta data for a token",
	Args:  cobra.ExactArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()

		logo, err := tokenregistry.LogoFromString(args[0])
		errorCheck("invalid logo", err)
		name, err := tokenregistry.NameFromString(args[1])
		errorCheck("invalid name", err)
		symbol, err := tokenregistry.SymbolFromString(args[2])
		errorCheck("invalid symbol", err)

		tokenAddress, err := solana.PublicKeyFromBase58(args[3])
		errorCheck("invalid token address", err)

		tokenMetaAccount := solana.NewAccount()

		keys := []solana.PublicKey{
			system.PROGRAM_ID,
			tokenregistry.ProgramID(),
			tokenMetaAccount.PublicKey(),
			tokenAddress,
		}

		fromAddressIndex := uint8(1)
		tokenMetaAddressIndex := uint8(2)
		tokenAddressIndex := uint8(3)
		size := 145

		lamport, err := client.GetMinimumBalanceForRentExemption(context.Background(), size)
		errorCheck("get minimum balance for rent exemption ", err)

		metaDataAccountInstruction, err := system.NewCreateAccount(
			bin.Uint64(lamport), bin.Uint64(size), system.PROGRAM_ID, 0, tokenMetaAddressIndex, fromAddressIndex,
		)
		errorCheck("new create account instruction", err)

		registerToken, err := tokenregistry.NewRegisterToken(logo, name, symbol, 1, tokenMetaAddressIndex, tokenMetaAddressIndex, tokenAddressIndex)
		errorCheck("new register token instruction", err)

		instructions := []solana.CompiledInstruction{
			*metaDataAccountInstruction,
			*registerToken,
		}

		trx := solana.Transaction{
			Signatures: nil,
			Message: solana.Message{
				Header:          solana.MessageHeader{},
				AccountKeys:     keys,
				RecentBlockhash: solana.PublicKey{},
				Instructions:    instructions,
			},
		}
		_ = trx
		return nil
	},
}

func init() {
	splCmd.AddCommand(tokenRegisterCmd)
}
