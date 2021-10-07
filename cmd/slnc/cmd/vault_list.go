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
	"github.com/spf13/cobra"
)

// vaultListCmd represents the list command
var vaultListCmd = &cobra.Command{
	Use:   "list",
	Short: "List public keys inside a Solana vault.",
	Long: `List public keys inside a Solana vault.

The wallet file contains a lits of public keys for easy reference, but
you cannot trust that these public keys have their counterpart in the
wallet, unless you check with the "list" command.
`,
	Run: func(cmd *cobra.Command, args []string) {
		vault := mustGetWallet()

		vault.PrintPublicKeys()
	},
}

func init() {
	vaultCmd.AddCommand(vaultListCmd)
}
