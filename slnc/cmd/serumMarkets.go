package cmd

import (
	"fmt"

	"github.com/dfuse-io/solana-go/serum"

	"github.com/spf13/cobra"
)

var serumMarketsCmd = &cobra.Command{
	Use:   "markets",
	Short: "Get serum markets",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		markets, err := serum.KnownMarket()
		if err != nil {
			return fmt.Errorf("unable to retrieve markets: %w", err)
		}

		for _, market := range markets {
			fmt.Printf("%s -> %s\n", market.Name, market.Address.String())
		}

		return nil
	},
}

func init() {
	serumCmd.AddCommand(serumMarketsCmd)
}
