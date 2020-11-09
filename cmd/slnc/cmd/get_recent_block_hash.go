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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var getRecentBlockhashCmd = &cobra.Command{
	Use:   "recent-blockhash",
	Short: "Retrieve a recent blockhash, needed for crafting transactions",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		resp, err := client.GetRecentBlockhash(ctx, "")
		if err != nil {
			return err
		}

		cnt, _ := json.MarshalIndent(resp.Value, "", "  ")
		fmt.Println(string(cnt))

		return nil
	},
}

func init() {
	getCmd.AddCommand(getRecentBlockhashCmd)
}
