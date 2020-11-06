package cmd

import (
	"fmt"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
)

var serumMarketsCmd = &cobra.Command{
	Use:   "markets",
	Short: "Get serum markets",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		rpcClient := rpc.NewClient("http://api.mainnet-beta.solana.com:80/rpc")
		programPubKey := solana.MustPublicKeyFromBase58("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o")

		d, err := base58.Decode("zTupHRte") // this is base58 of 0x736572756d03
		if err != nil {
			return err
		}
		results, err := rpcClient.GetProgramAccounts(cmd.Context(), programPubKey, &rpc.GetProgramAccountsOpts{
			Encoding: "base64",
			Filters: []rpc.RPCFilter{
				{
					Memcmp: &rpc.RPCFilterMemcmp{
						Offset: 0,
						Bytes:  d,
					},
				},
			},
		})
		if err != nil {
			return fmt.Errorf("unable to retrieve accounts: %w", err)
		}

		for _, account := range results {
			fmt.Println(account.Pubkey.String())
		}

		return nil
	},
}

func init() {
	serumCmd.AddCommand(serumMarketsCmd)
}
