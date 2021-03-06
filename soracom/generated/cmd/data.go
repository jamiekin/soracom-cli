package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DataCmd)
}

// DataCmd defines 'data' subcommand
var DataCmd = &cobra.Command{
	Use:   "data",
	Short: TRCLI("cli.data.summary"),
	Long:  TRCLI(`cli.data.description`),
}
