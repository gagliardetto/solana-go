package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var getConfirmedBlockCmd = &cobra.Command{
	Use:   "confirmed-block {block_num}",
	Short: "Retrieve a confirmed block, with all of its transactions",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		slot, err := strconv.ParseInt(args[0], 10, 64)
		errorCheck("parsing slot number in first argument", err)

		resp, err := client.GetConfirmedBlock(ctx, uint64(slot), "")
		if err != nil {
			return err
		}

		cnt, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(cnt))

		// for idx, trx := range resp.Transactions {
		// 	unpacked, err := solana.TransactionFromData(trx.Transaction)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	cnt, _ := json.MarshalIndent(unpacked, "", "  ")
		// 	fmt.Println(idx, string(cnt))
		// }

		return nil
	},
}

func init() {
	getCmd.AddCommand(getConfirmedBlockCmd)
}
