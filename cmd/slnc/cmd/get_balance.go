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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var getBalanceCmd = &cobra.Command{
	Use:   "balance {account_addr}",
	Short: "Retrieve an account's balance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		resp, err := client.GetBalance(ctx, args[0], "")
		if err != nil {
			return err
		}

		if resp.Value == 0 {
			errorCheck("not found", errors.New("account not found"))
		}

		fmt.Println(resp.Value, "lamports")

		return nil
	},
}

func init() {
	getCmd.AddCommand(getBalanceCmd)
}
