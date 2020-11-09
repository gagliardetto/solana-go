package cmd

import (
	"github.com/spf13/cobra"
)

var serumListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Serum objects",
}

func init() {
	serumCmd.AddCommand(serumListCmd)
}
