package cmd

import (
	"bytes"
	"context"
	"strings"
	"github.com/dfuse-io/solana-go/serum"
	"github.com/lunixbochs/struc"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
)

var serumOrderbookCmd = &cobra.Command{
	Use:   "orderbook {market_addr}",
	Short: "Get Serum orderbook for a given market",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pubKey, err := solana.PublicKeyFromBase58(args[0])
		if err != nil {
			return fmt.Errorf("decoding market addr: %w", err)
		}

		ctx := context.Background()

		bids, err := rpcClient.GetProgramAccounts(ctx, pubKey, &rpc.GetProgramAccountsOpts{
			Encoding: "base64",
			Filters:  []rpc.RPCFilter{
				{Memcmp: &rpc.RPCFilterMemcmp{Bytes: []byte{0x73, 0x65, 0x72, 0x75, 0x6d, 0x21}}}
			},
		})
		if err != nil {
			return fmt.Errorf("failed query: %w", err)
		}

		fmt.Println("Query done")

		for _, bid := range bids {
			fmt.Println("Big", bid.Pubkey)
			data, err := bid.Account.DataToBytes()
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
			o.Items(true, func(node *serum.SlabLeafNode) error {
				if uint64(node.Quantity) > highestQuantity {
					highestQuantity = uint64(node.Quantity)
				}
				return nil
			})

			o.Items(true, func(node *serum.SlabLeafNode) error {
				// TODO: compute the actual price and lots size?
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
		}
		return nil
	},
}

func init() {
	serumCmd.AddCommand(serumOrderbookCmd)
}
