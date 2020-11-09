package cmd

import (
	"github.com/spf13/cobra"
)

var splGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves SPL token objects",
}

func init() {
	splCmd.AddCommand(splGetCmd)
}
