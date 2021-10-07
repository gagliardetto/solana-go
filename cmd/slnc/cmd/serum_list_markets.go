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
	"fmt"

	"github.com/gagliardetto/solana-go/programs/serum"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var serumListMarketsCmd = &cobra.Command{
	Use:   "markets",
	Short: "Get serum markets",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		markets, err := serum.KnownMarket()
		if err != nil {
			return fmt.Errorf("unable to retrieve markets: %w", err)
		}

		out := []string{"Pairs | Market Address"}

		for _, market := range markets {
			out = append(out, fmt.Sprintf("%s | %s ", market.Name, market.Address.String()))
		}

		fmt.Println(columnize.Format(out, nil))

		return nil
	},
}

func init() {
	serumListCmd.AddCommand(serumListMarketsCmd)
}
