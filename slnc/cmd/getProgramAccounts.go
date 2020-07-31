package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/dfuse-io/solana-go/token"
	"github.com/lunixbochs/struc"
	"github.com/spf13/cobra"
)

var getProgramAccountsCmd = &cobra.Command{
	Use:   "program-accounts",
	Short: "Retrieve info about an account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		resp, err := client.GetProgramAccounts(ctx, args[0], nil)
		if err != nil {
			return err
		}

		if resp == nil {
			errorCheck("not found", errors.New("program account not found"))
		}

		for _, keyedAcct := range *resp {
			acct := keyedAcct.Account
			fmt.Println("Data len:", len(acct.Data), keyedAcct.Pubkey)

			switch len(acct.Data) {
			case 120:
				var tokenAcct token.Account
				if err := struc.Unpack(bytes.NewReader(acct.Data), &tokenAcct); err != nil {
					log.Fatalln("failed unpack", err)
				}

				cnt, _ := json.MarshalIndent(tokenAcct, "", "  ")
				fmt.Println(string(cnt))
			case 40:
				var mint token.Mint
				if err := struc.Unpack(bytes.NewReader(acct.Data), &mint); err != nil {
					log.Fatalln("failed unpack", err)
				}

				cnt, _ := json.MarshalIndent(mint, "", "  ")
				fmt.Println(string(cnt))
			default:
				fmt.Println("Unknown data length")
			}
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getProgramAccountsCmd)
}
