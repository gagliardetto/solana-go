package cmd

import (
	"bytes"
	"fmt"
	"math/big"

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
		ctx := cmd.Context()

		marketAddr, err := solana.PublicKeyFromBase58(args[0])
		if err != nil {
			return fmt.Errorf("decoding market addr: %w", err)
		}

		cli := getClient()
		market, err := serum.FetchMarket(ctx, cli, marketAddr)
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

		//var highestQuantity uint64
		//var sumQty uint64
		//o.Items(true, func(node *serum.SlabLeafNode) error {
		//	if uint64(node.Quantity) > highestQuantity {
		//		highestQuantity = uint64(node.Quantity)
		//	}
		//	sumQty += uint64(node.Quantity)
		//	return nil
		//})

		limit := 20
		levels := [][]uint64{}
		o.Items(true, func(node *serum.SlabLeafNode) error {
			fmt.Println("leaf: ", node.KeyPrice, node.Quantity)
			if len(levels) > 0 && levels[len(levels)-1][0] == uint64(node.KeyPrice) {
				levels[len(levels)-1][1] += uint64(node.Quantity)
			} else if len(levels) == limit {
				return fmt.Errorf("done")
			} else {
				levels = append(levels, []uint64{uint64(node.KeyPrice), uint64(node.Quantity)})
			}
			return nil
		})

		for _, level := range levels {
			price := market.PriceLotsToNumber(big.NewInt(int64(level[0])))
			qty := market.BaseSizeLotsToNumber(big.NewInt(int64(level[1])))
			output = append(output,
				fmt.Sprintf("%s | %s",
					price.String(),
					qty.String(),
				),
			)
		}

		// TODO: compute the actual price and lots size?
		//price := serum.ComputePrice(node.KeyPrice, market.BaseMint)

		fmt.Println(columnize.Format(output, nil))

		// Repeat for second segment of the order book

		return nil
	},
}

func init() {
	serumCmd.AddCommand(serumMarketCmd)
}
