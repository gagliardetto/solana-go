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

		resp, err := client.GetAccountInfo(ctx, args[0])
		if err != nil {
			return err
		}

		if resp.Value == nil {
			errorCheck("not found", errors.New("account not found"))
		}

		acct := resp.Value

		data, err := json.MarshalIndent(acct, "", "  ")
		errorCheck("json marshal", err)

		fmt.Println(string(data))

		obj, err := decode(acct.Owner, acct.Data)
		if err != nil {
			return err
		}

		if obj != nil {
			cnt, err := json.MarshalIndent(obj, "", "  ")
			if err != nil {
				return err
			}
			fmt.Printf("Data %T: %s\n", obj, string(cnt))
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getAccountCmd)
}
