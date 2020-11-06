package cmd

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/dfuse-io/solana-go/serum"
	"github.com/lunixbochs/struc"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var serumMarketCmd = &cobra.Command{
	Use:   "market {market_addr}",
	Short: "Get Serum orderbook for a given market",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		marketAddr, err := solana.PublicKeyFromBase58(args[0])
		if err != nil {
			return fmt.Errorf("decoding market addr: %w", err)
		}

		ctx := context.Background()

		cli := serum.NewSerumClient("http://api.mainnet-beta.solana.com:80/rpc")
		market, err := cli.FetchMarket(ctx, marketAddr)
		if err != nil {
			return fmt.Errorf("fetch market: %w", err)
		}

		rpcClient := rpc.NewClient("http://api.mainnet-beta.solana.com:80/rpc")

		// Print first segment of the order book

		bids, err := rpcClient.GetAccountInfo(ctx, market.MarketV2.Bids)
		if err != nil {
			return fmt.Errorf("failed query: %w", err)
		}

		fmt.Println("Query done")

		data, err := bids.Value.DataToBytes()
		if err != nil {
			return fmt.Errorf("decoding bid account data: %w", err)
		}
		var o serum.Orderbook
		if err := struc.Unpack(bytes.NewReader(data), &o); err != nil {
			return fmt.Errorf("decoding bid orderbook data: %w", err)
		}

		output := []string{
			"Price | Quantity | Depth",
		}

		var highestQuantity uint64
		var sumQty uint64
		o.Items(true, func(node *serum.SlabLeafNode) error {
			if uint64(node.Quantity) > highestQuantity {
				highestQuantity = uint64(node.Quantity)
			}
			sumQty += uint64(node.Quantity)
			return nil
		})

		o.Items(true, func(node *serum.SlabLeafNode) error {
			// TODO: compute the actual price and lots size?
			//price := serum.ComputePrice(node.KeyPrice, market.BaseMint)
			price := node.KeyPrice
			qty := node.Quantity
			percent := uint64(qty) * 100 / highestQuantity
			output = append(output,
				fmt.Sprintf("%d | %d | %s",
					price,
					qty,
					strings.Repeat("-", int(percent/10))+strings.Repeat("#", int(10-(percent/10))),
				),
			)

			return nil
		})

		fmt.Println(columnize.Format(output, nil))

		// Repeat for second segment of the order book

		return nil
	},
}

func init() {
	serumCmd.AddCommand(serumMarketCmd)
}
