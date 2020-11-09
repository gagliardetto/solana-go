package cmd

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
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

		ioutil.WriteFile("/tmp/bids-go.txt", []byte(hex.EncodeToString(bids.Value.MustDataToBytes())), 0644)

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

		limit := 20
		levels := [][]*big.Int{}
		o.Items(true, func(node *serum.SlabLeafNode) error {
			quantity := big.NewInt(int64(node.Quantity))
			price := node.GetPrice()
			if len(levels) > 0 && levels[len(levels)-1][0].Cmp(price) == 0 {
				current := levels[len(levels)-1][1]
				levels[len(levels)-1][1] = new(big.Int).Add(current, quantity)
			} else if len(levels) == limit {
				return fmt.Errorf("done")
			} else {
				levels = append(levels, []*big.Int{price, quantity})
			}
			return nil
		})

		for _, level := range levels {
			price := market.PriceLotsToNumber(level[0])
			qty := market.BaseSizeLotsToNumber(level[1])
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
