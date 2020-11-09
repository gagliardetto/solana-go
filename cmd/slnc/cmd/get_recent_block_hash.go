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
