package cmd

import (
	"github.com/spf13/cobra"
)

var serumGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Serum objects",
}

func init() {
	serumCmd.AddCommand(serumGetCmd)
}
