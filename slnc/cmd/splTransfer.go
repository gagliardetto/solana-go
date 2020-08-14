package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var splTransferCmd = &cobra.Command{
	Use:   "transfer [from] [to] [amount]",
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
