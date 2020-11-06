package cmd

import (
	"fmt"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/spf13/cobra"
)

var serumOrderbookCmd = &cobra.Command{
	Use:   "orderbook {market_addr}",
	Short: "Get serum orderbook for a market",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pubKey, err := solana.PublicKeyFromBase58(args[0])
		if err != nil {
			return err
		}

		rpcClient := rpc.NewClient("http://api.mainnet-beta.solana.com:80/rpc")

		marketAddr := args[0]
		fmt.Println("")
		fmt.Println(marketAddr)
		fmt.Println("")
		return nil
	},
}

func init() {
	serumCmd.AddCommand(serumMarketsCmd)
}
