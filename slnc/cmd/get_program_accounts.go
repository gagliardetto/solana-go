package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dfuse-io/solana-go"

	"github.com/spf13/cobra"
)

var getProgramAccountsCmd = &cobra.Command{
	Use:   "program-accounts {program_addr}",
	Short: "Retrieve info about an account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		resp, err := client.GetProgramAccounts(ctx, solana.MustPublicKeyFromBase58(args[0]), nil)
		if err != nil {
			return err
		}

		if resp == nil {
			errorCheck("not found", errors.New("program account not found"))
		}

		for _, keyedAcct := range resp {
			acct := keyedAcct.Account
			//fmt.Println("Data len:", len(acct.Data), keyedAcct.Pubkey)

			obj, err := decode(acct.Owner, acct.MustDataToBytes())
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
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getProgramAccountsCmd)
}
