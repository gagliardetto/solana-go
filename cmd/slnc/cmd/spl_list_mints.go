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

	"github.com/dfuse-io/solana-go/programs/token"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var splListMintsCmd = &cobra.Command{
	Use:   "mints",
	Short: "Lists mints",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: implement a different network argument,
		// later. Ultimately, get on chain. We have a database here!

		mints, err := token.KnownMints("mainnet")
		if err != nil {
			return fmt.Errorf("listing mints: %w", err)
		}
		out := []string{"Symbol | Mint address | Token name"}

		for _, m := range mints {
			out = append(out, fmt.Sprintf("%s | %s | %s", m.TokenSymbol, m.MintAddress, m.TokenName))
		}

		fmt.Println(columnize.Format(out, nil))

		return nil
	},
}

func init() {
	splListCmd.AddCommand(splListMintsCmd)
}
