package cmd

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch information from a cluster",
}

func init() {
	RootCmd.AddCommand(getCmd)
}
