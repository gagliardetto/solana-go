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
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/dfuse-io/solana-go/rpc"

	bin "github.com/dfuse-io/binary"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/programs/system"
	"github.com/dfuse-io/solana-go/programs/tokenregistry"
	"github.com/spf13/cobra"
)

var tokenRegisterCmd = &cobra.Command{
	Use:   "token register {token} {name} {symbol} {logo}",
	Short: "register meta data for a token",
	Args:  cobra.ExactArgs(5),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()

		tokenAddress, err := solana.PublicKeyFromBase58(args[1])
		errorCheck("invalid token address", err)

		logo, err := tokenregistry.LogoFromString(args[2])
		errorCheck("invalid logo", err)
		name, err := tokenregistry.NameFromString(args[3])
		errorCheck("invalid name", err)
		symbol, err := tokenregistry.SymbolFromString(args[4])
		errorCheck("invalid symbol", err)

		tokenMetaAccount := solana.NewAccount()

		programIDAddress := tokenregistry.ProgramID()
		keys := []solana.PublicKey{
			programIDAddress,
			tokenMetaAccount.PublicKey(),
			system.PROGRAM_ID,
			tokenAddress,
		}

		programAccountIndex := uint8(0)
		fromAddressIndex := uint8(0)
		tokenMetaAddressIndex := uint8(1)
		systemIDIndex := uint8(2)
		tokenAddressIndex := uint8(3)

		size := 145

		lamport, err := client.GetMinimumBalanceForRentExemption(context.Background(), size)
		errorCheck("get minimum balance for rent exemption ", err)

		metaDataAccountInstruction, err := system.NewCreateAccount(
			bin.Uint64(lamport), bin.Uint64(size), system.PROGRAM_ID, systemIDIndex, tokenMetaAddressIndex, fromAddressIndex,
		)
		errorCheck("new create account instruction", err)

		registerToken, err := tokenregistry.NewRegisterToken(logo, name, symbol, programAccountIndex, tokenMetaAddressIndex, tokenMetaAddressIndex, tokenAddressIndex)
		errorCheck("new register token instruction", err)

		instructions := []solana.CompiledInstruction{
			*metaDataAccountInstruction,
			*registerToken,
		}

		blockHashResult, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentRecent)
		errorCheck("get block recent block hash", err)

		message := solana.Message{
			Header: solana.MessageHeader{
				NumRequiredSignatures:       1,
				NumReadonlySignedAccounts:   1,
				NumReadonlyunsignedAccounts: 2,
			},
			AccountKeys:     keys,
			RecentBlockhash: blockHashResult.Value.Blockhash,
			Instructions:    instructions,
		}

		var dataToSign []byte
		buf := bytes.NewBuffer(dataToSign)
		err = bin.NewEncoder(buf).Encode(message)
		errorCheck("message encoding", err)

		v := mustGetWallet()
		var signature solana.Signature
		var signed bool
		for _, privateKey := range v.KeyBag {
			if privateKey.PublicKey() == programIDAddress {
				signed = true
				signature, err = privateKey.Sign(dataToSign)
				errorCheck("signe message", err)
			}
		}

		if !signed {
			fmt.Errorf("unable to find matching private key for signing")
			os.Exit(1)
		}

		trx := &solana.Transaction{
			Signatures: []solana.Signature{signature},
			Message:    message,
		}

		trxHash, err := client.SendTransaction(context.Background(), trx)
		fmt.Println("sent transaction", trxHash, err)
		return nil
	},
}

func init() {
	splCmd.AddCommand(tokenRegisterCmd)
}
