package cmd

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/dfuse-io/solana-go/serum"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var serumGetMarketCmd = &cobra.Command{
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

		output := []string{
			"Price | Quantity | Depth",
			"Asks",
		}

		asks, err := getOrderBook(ctx, market, cli, market.MarketV2.Asks, true)
		if err != nil {
			return fmt.Errorf("unable to retrieve asks: %w", err)
		}
		for _, a := range asks {
			output = append(output, fmt.Sprintf("%s | %s",
				a.price.String(),
				a.quantity.String(),
			))
		}

		output = append(output, "Bids")
		bids, err := getOrderBook(ctx, market, cli, market.MarketV2.Bids, true)
		if err != nil {
			return fmt.Errorf("unable to retrieve bids: %w", err)
		}
		for _, b := range bids {
			output = append(output, fmt.Sprintf("%s | %s",
				b.price.String(),
				b.quantity.String(),
			))
		}

		fmt.Println(columnize.Format(output, nil))
		return nil
	},
}

type orderBookEntry struct {
	price    *big.Float
	quantity *big.Float
}

func getOrderBook(ctx context.Context, market *serum.MarketMeta, cli *rpc.Client, address solana.PublicKey, desc bool) (out []*orderBookEntry, err error) {
	var o serum.Orderbook
	if err := cli.GetAccountDataIn(ctx, address, &o); err != nil {
		return nil, fmt.Errorf("getting orderbook: %w", err)
	}

	limit := 20
	levels := [][]*big.Int{}
	o.Items(desc, func(node *serum.SlabLeafNode) error {
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
		out = append(out,
			&orderBookEntry{
				price:    price,
				quantity: qty,
			},
		)
	}
	return out, nil
}

func init() {
	serumGetCmd.AddCommand(serumGetMarketCmd)
}
