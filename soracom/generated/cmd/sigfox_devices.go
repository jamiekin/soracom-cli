package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(SigfoxDevicesCmd)
}

// SigfoxDevicesCmd defines 'sigfox-devices' subcommand
var SigfoxDevicesCmd = &cobra.Command{
	Use:   "sigfox-devices",
	Short: TRCLI("cli.sigfox-devices.summary"),
	Long:  TRCLI(`cli.sigfox-devices.description`),
}
