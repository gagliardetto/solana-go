package cmd

import (
	"fmt"

	"github.com/dfuse-io/solana-go/serum"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var serumListMarketsCmd = &cobra.Command{
	Use:   "markets",
	Short: "Get serum markets",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		markets, err := serum.KnownMarket()
		if err != nil {
			return fmt.Errorf("unable to retrieve markets: %w", err)
		}

		out := []string{"Pairs | Market Address"}

		for _, market := range markets {
			out = append(out, fmt.Sprintf("%s | %s ", market.Name, market.Address.String()))
		}

		fmt.Println(columnize.Format(out, nil))

		return nil
	},
}

func init() {
	serumListCmd.AddCommand(serumListMarketsCmd)
}
