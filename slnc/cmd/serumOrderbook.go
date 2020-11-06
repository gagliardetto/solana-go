package cmd

import (
	"github.com/spf13/cobra"
)

var serumOrderbookCmd = &cobra.Command{
	Use:   "orderbook {market_addr}",
	Short: "Get serum orderbook for a market",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	serumCmd.AddCommand(serumMarketsCmd)
}
