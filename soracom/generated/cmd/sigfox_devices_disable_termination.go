package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesDisableTerminationCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesDisableTerminationCmdDeviceId string

func init() {
	SigfoxDevicesDisableTerminationCmd.Flags().StringVar(&SigfoxDevicesDisableTerminationCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesDisableTerminationCmd)
}

// SigfoxDevicesDisableTerminationCmd defines 'disable-termination' subcommand
var SigfoxDevicesDisableTerminationCmd = &cobra.Command{
	Use:   "disable-termination",
	Short: TRAPI("/sigfox_devices/{device_id}/disable_termination:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/disable_termination:post:description`),
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

		param, err := collectSigfoxDevicesDisableTerminationCmdParams()
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

func collectSigfoxDevicesDisableTerminationCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSigfoxDevicesDisableTerminationCmd("/sigfox_devices/{device_id}/disable_termination"),
		query:  buildQueryForSigfoxDevicesDisableTerminationCmd(),
	}, nil
}

func buildPathForSigfoxDevicesDisableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesDisableTerminationCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesDisableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
