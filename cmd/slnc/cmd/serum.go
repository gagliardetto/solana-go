package cmd

import "github.com/spf13/cobra"

var serumCmd = &cobra.Command{
	Use:          "serum",
	Short:        "serum commands",
	SilenceUsage: false,
}

func init() {
	RootCmd.AddCommand(serumCmd)
}
