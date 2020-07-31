package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var getBalanceCmd = &cobra.Command{
	Use:   "balance",
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
