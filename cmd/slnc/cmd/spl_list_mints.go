package cmd

import (
	"fmt"

	"github.com/dfuse-io/solana-go/token"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var splListMintsCmd = &cobra.Command{
	Use:   "mints",
	Short: "Lists mints",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: implement a different network argument,
		// later. Ultimately, get on chain. We have a database here!

		mints, err := token.KnownMints("mainnet")
		if err != nil {
			return fmt.Errorf("listing mints: %w", err)
		}
		out := []string{"Symbol | Mint address | Token name"}

		for _, m := range mints {
			out = append(out, fmt.Sprintf("%s | %s | %s", m.TokenSymbol, m.MintAddress, m.TokenName))
		}

		fmt.Println(columnize.Format(out, nil))

		return nil
	},
}

func init() {
	splListCmd.AddCommand(splListMintsCmd)
}
