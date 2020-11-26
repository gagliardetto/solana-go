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
	"encoding/hex"
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

		tokenMetaDataAccount := solana.NewAccount()

		alexKey := solana.MustPublicKeyFromBase58("9hFtYBYmBJCVguRYs9pBTWKYAFoKfjYR7zBPpEkVsmD")

		//todo: remove
		airDrop, err := client.RequestAirdrop(context.Background(), &alexKey, 10_000_000_000, rpc.CommitmentMax)
		errorCheck("air drop", err)
		fmt.Println("air drop hash:", airDrop)

		tokenRegistryProgramIDAddress := tokenregistry.ProgramID()
		keys := []solana.PublicKey{
			alexKey,
			tokenMetaDataAccount.PublicKey(),
			tokenRegistryProgramIDAddress,
			tokenAddress,
			system.PROGRAM_ID,
		}
		alexKeyIndex := uint8(0)
		tokenMetaDataAddressIndex := uint8(1)
		tokenRegistryProgramAccountIndex := uint8(2)
		tokenAddressIndex := uint8(3)
		systemIDIndex := uint8(4)

		size := 145
		lamport, err := client.GetMinimumBalanceForRentExemption(context.Background(), size)
		errorCheck("get minimum balance for rent exemption ", err)

		fmt.Println("minimum lamport for rent exemption:", lamport)

		from := alexKeyIndex
		to := tokenMetaDataAddressIndex
		metaDataAccountCreationInstruction, err := system.NewCreateAccount(
			bin.Uint64(lamport), bin.Uint64(size), tokenRegistryProgramIDAddress, systemIDIndex, from, to,
		)
		errorCheck("new create account instruction", err)

		buf := new(bytes.Buffer)
		if err := bin.NewEncoder(buf).Encode(metaDataAccountCreationInstruction); err != nil {
			panic(err)
		}
		fmt.Println("create account instruction hex:", hex.EncodeToString(buf.Bytes()))

		programIdIndex := tokenRegistryProgramAccountIndex
		mintMetaIdx := tokenMetaDataAddressIndex
		ownerIdx := alexKeyIndex
		tokenIdx := tokenAddressIndex

		/// 0. `[writable]` The register data's account to initialize
		/// 1. `[signer]` The registry's owner
		/// 2. `[]` The mint address to link with this registration
		registerToken, err := tokenregistry.NewRegisterToken(logo, name, symbol, programIdIndex, mintMetaIdx, ownerIdx, tokenIdx)
		errorCheck("new register token instruction", err)

		_ = registerToken

		instructions := []solana.CompiledInstruction{
			*metaDataAccountCreationInstruction,
			*registerToken,
		}

		blockHashResult, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentMax)
		errorCheck("get block recent block hash", err)

		message := solana.Message{
			Header: solana.MessageHeader{
				NumRequiredSignatures:       2,
				NumReadonlySignedAccounts:   0,
				NumReadonlyunsignedAccounts: 2,
			},
			AccountKeys:     keys,
			RecentBlockhash: blockHashResult.Value.Blockhash,
			Instructions:    instructions,
		}

		buf = new(bytes.Buffer)
		err = bin.NewEncoder(buf).Encode(message)
		errorCheck("message encoding", err)
		dataToSign := buf.Bytes()
		fmt.Println("Data to sign:", buf)

		v := mustGetWallet()
		var signature solana.Signature
		var signed bool
		for _, privateKey := range v.KeyBag {
			if privateKey.PublicKey() == alexKey {
				signature, err = privateKey.Sign(dataToSign)
				errorCheck("signe message", err)
				signed = true
			}
		}
		fmt.Println("signing completed")

		if !signed {
			fmt.Println("unable to find matching private key for signing")
			os.Exit(1)
		}

		fmt.Println("signature:", signature.String())

		tokenMetaAccountSignature, err := tokenMetaDataAccount.PrivateKey.Sign(dataToSign)
		errorCheck("tokenMetaAccountSignature", err)

		trx := &solana.Transaction{
			Signatures: []solana.Signature{signature, tokenMetaAccountSignature},
			Message:    message,
		}

		trxHash, err := client.SendTransaction(context.Background(), trx)
		//trxHash, err := client.SimulateTransaction(context.Background(), trx)
		fmt.Println("sent transaction hash:", trxHash, " error:", err)
		return nil
	},
}

func init() {
	splCmd.AddCommand(tokenRegisterCmd)
}
