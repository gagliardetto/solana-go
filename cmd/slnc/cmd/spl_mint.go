package cmd

import (
	"fmt"

	"github.com/dfuse-io/solana-go/token"

	"github.com/dfuse-io/solana-go"

	"github.com/spf13/cobra"
)

var splGetMintCmd = &cobra.Command{
	Use:   "get-mint {mint_addr}",
	Short: "Retrieves mint information",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		mintAddr, err := solana.PublicKeyFromBase58(args[0])
		if err != nil {
			return fmt.Errorf("decoding mint addr: %w", err)
		}

		client := getClient()

		mint, err := token.GetMint(ctx, client, mintAddr)
		if err != nil {
			return fmt.Errorf("unable to retrieve int information: %w", err)
		}

		fmt.Println("")
		fmt.Println("Mint Authority Option: ", mint.MintAuthorityOption)
		fmt.Println("Mint Authority: ", mint.MintAuthority)
		fmt.Println("Supply: ", mint.Supply)
		fmt.Println("Decimals: ", mint.Decimals)
		fmt.Println("Is Initialized: ", mint.IsInitialized)
		fmt.Println("Freeze Authority Option: ", mint.FreezeAuthorityOption)
		fmt.Println("Freeze Authority: ", mint.FreezeAuthority)
		fmt.Println("")

		return nil
	},
}

func init() {
	splCmd.AddCommand(splGetMintCmd)
}
