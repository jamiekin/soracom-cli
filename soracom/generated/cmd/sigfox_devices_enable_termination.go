package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesEnableTerminationCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesEnableTerminationCmdDeviceId string

func init() {
	SigfoxDevicesEnableTerminationCmd.Flags().StringVar(&SigfoxDevicesEnableTerminationCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesEnableTerminationCmd)
}

// SigfoxDevicesEnableTerminationCmd defines 'enable-termination' subcommand
var SigfoxDevicesEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/sigfox_devices/{device_id}/enable_termination:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/enable_termination:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectSigfoxDevicesEnableTerminationCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectSigfoxDevicesEnableTerminationCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSigfoxDevicesEnableTerminationCmd("/sigfox_devices/{device_id}/enable_termination"),
		query:  buildQueryForSigfoxDevicesEnableTerminationCmd(),
	}, nil
}

func buildPathForSigfoxDevicesEnableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesEnableTerminationCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesEnableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
