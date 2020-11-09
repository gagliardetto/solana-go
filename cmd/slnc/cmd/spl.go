package cmd

import "github.com/spf13/cobra"

var splCmd = &cobra.Command{
	Use:   "spl",
	Short: "SPL Tokens related Instructions",
}

func init() {
	RootCmd.AddCommand(splCmd)
}
