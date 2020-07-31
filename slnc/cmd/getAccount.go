package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var getAccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Retrieve info about an account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		resp, err := client.GetAccountInfo(ctx, args[0], "")
		if err != nil {
			return err
		}

		if resp.Value == nil {
			errorCheck("not found", errors.New("account not found"))
		}

		data, err := json.MarshalIndent(resp.Value, "", "  ")
		errorCheck("json marshal", err)

		fmt.Println(string(data))

		return nil
	},
}

func init() {
	getCmd.AddCommand(getAccountCmd)
}
