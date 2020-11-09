package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strings"

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

		asks, askSize, err := getOrderBook(ctx, market, cli, market.MarketV2.Asks, false)
		if err != nil {
			return fmt.Errorf("unable to retrieve asks: %w", err)
		}

		bids, bidSize, err := getOrderBook(ctx, market, cli, market.MarketV2.Bids, true)
		if err != nil {
			return fmt.Errorf("unable to retrieve bids: %w", err)
		}
		totalSize := new(big.Float).Add(askSize, bidSize)

		output := []string{
			"Price | Quantity | Depth",
			"Asks",
		}
		output = append(output, outputOrderBook(asks, totalSize, true)...)
		output = append(output, "------- | --------")
		output = append(output, outputOrderBook(bids, totalSize, false)...)
		output = append(output, "Bids")
		fmt.Println(columnize.Format(output, nil))
		return nil
	},
}

type orderBookEntry struct {
	price    *big.Float
	quantity *big.Float
}

func getOrderBook(ctx context.Context, market *serum.MarketMeta, cli *rpc.Client, address solana.PublicKey, desc bool) (out []*orderBookEntry, totalSize *big.Float, err error) {
	var o serum.Orderbook
	if err := cli.GetAccountDataIn(ctx, address, &o); err != nil {
		return nil, nil, fmt.Errorf("getting orderbook: %w", err)
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

	totalSize = big.NewFloat(0)
	for _, level := range levels {
		price := market.PriceLotsToNumber(level[0])
		qty := market.BaseSizeLotsToNumber(level[1])
		totalSize = new(big.Float).Add(totalSize, qty)
		out = append(out,
			&orderBookEntry{
				price:    price,
				quantity: qty,
			},
		)
	}
	return out, totalSize, nil
}

func depth(value *big.Float) string {
	v, _ := value.Int(nil)
	return strings.Repeat("#", int(v.Int64()))
}

func outputOrderBook(entries []*orderBookEntry, totalSize *big.Float, reverse bool) (out []string) {
	total := totalSize
	if totalSize == nil {
		total = new(big.Float).SetInt64(1)
	}

	type orderBookRow struct {
		price    string
		quantity string
		depth    string
	}

	rows := []*orderBookRow{}
	cumulativeSize := big.NewFloat(0)
	for i := 0; i < len(entries); i++ {
		entry := entries[i]
		cumulativeSize = new(big.Float).Add(cumulativeSize, entry.quantity)
		sizePercent := new(big.Float).Mul(new(big.Float).Quo(cumulativeSize, total), new(big.Float).SetInt64(100))
		rows = append(rows, &orderBookRow{
			price:    entry.price.String(),
			quantity: entry.quantity.String(),
			depth:    depth(sizePercent),
		})
	}

	if reverse {
		for i := len(entries) - 1; i >= 0; i-- {
			out = append(out, fmt.Sprintf("%s | %s | %s",
				rows[i].quantity,
				rows[i].price,
				rows[i].depth,
			))
		}
		return
	}
	for i := 0; i < len(rows); i++ {
		out = append(out, fmt.Sprintf("%s | %s | %s",
			rows[i].quantity,
			rows[i].price,
			rows[i].depth,
		))
	}
	return
}
func init() {
	serumGetCmd.AddCommand(serumGetMarketCmd)
}
