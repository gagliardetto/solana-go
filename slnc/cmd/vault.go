package cmd

import (
	"github.com/spf13/cobra"
)

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "The solana-go Vault is a secure Solana key vault",
}

func init() {
	RootCmd.AddCommand(vaultCmd)
}
