package cmd

import (
	"github.com/spf13/cobra"
)

var splListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves SPL token objects",
}

func init() {
	splCmd.AddCommand(splListCmd)
}
