package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var getSlotCmd = &cobra.Command{
	Use:   "slot",
	Short: "Retrieve the current slot the node is processing",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		ctx := context.Background()

		resp, err := client.GetSlot(ctx, "")
		if err != nil {
			return err
		}

		fmt.Println(resp)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getSlotCmd)
}
