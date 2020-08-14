package cmd

import "github.com/spf13/cobra"

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "System Instructions",
}

func init() {
	RootCmd.AddCommand(systemCmd)
}
