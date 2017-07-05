package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesUnobserveResourceCmdDeviceId holds value of 'deviceId' option
var DevicesUnobserveResourceCmdDeviceId string

// DevicesUnobserveResourceCmdInstance holds value of 'instance' option
var DevicesUnobserveResourceCmdInstance string

// DevicesUnobserveResourceCmdObject holds value of 'object' option
var DevicesUnobserveResourceCmdObject string

// DevicesUnobserveResourceCmdResource holds value of 'resource' option
var DevicesUnobserveResourceCmdResource string

func init() {
	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdDeviceId, "device-id", "", TR("devices.unobserve_resource.parameters.deviceId.description"))

	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdInstance, "instance", "", TR("devices.unobserve_resource.parameters.instance.description"))

	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdObject, "object", "", TR("devices.unobserve_resource.parameters.object.description"))

	DevicesUnobserveResourceCmd.Flags().StringVar(&DevicesUnobserveResourceCmdResource, "resource", "", TR("devices.unobserve_resource.parameters.resource.description"))

	DevicesCmd.AddCommand(DevicesUnobserveResourceCmd)
}

// DevicesUnobserveResourceCmd defines 'unobserve-resource' subcommand
var DevicesUnobserveResourceCmd = &cobra.Command{
	Use:   "unobserve-resource",
	Short: TR("devices.unobserve_resource.summary"),
	Long:  TR(`devices.unobserve_resource.description`),
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

		param, err := collectDevicesUnobserveResourceCmdParams()
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

func collectDevicesUnobserveResourceCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesUnobserveResourceCmd("/devices/{deviceId}/{object}/{instance}/{resource}/unobserve"),
		query:  buildQueryForDevicesUnobserveResourceCmd(),
	}, nil
}

func buildPathForDevicesUnobserveResourceCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesUnobserveResourceCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"instance"+"}", DevicesUnobserveResourceCmdInstance, -1)

	path = strings.Replace(path, "{"+"object"+"}", DevicesUnobserveResourceCmdObject, -1)

	path = strings.Replace(path, "{"+"resource"+"}", DevicesUnobserveResourceCmdResource, -1)

	return path
}

func buildQueryForDevicesUnobserveResourceCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
