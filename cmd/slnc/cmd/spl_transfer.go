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

	"github.com/spf13/cobra"
)

var splTransferCmd = &cobra.Command{
	Use:   "transfer {from} {to} {amount}",
	Short: "Create and sign an SPL transfer transaction",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		from := args[0]
		to := args[1]
		amount := args[2]

		fmt.Println(from, to, amount)

		_ = client
		_ = ctx

		return nil
	},
}

func init() {
	splCmd.AddCommand(splTransferCmd)
}
