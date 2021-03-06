package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ShippingAddressesListCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesListCmdOperatorId string

func init() {
	ShippingAddressesListCmd.Flags().StringVar(&ShippingAddressesListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	ShippingAddressesCmd.AddCommand(ShippingAddressesListCmd)
}

// ShippingAddressesListCmd defines 'list' subcommand
var ShippingAddressesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/shipping_addresses:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/shipping_addresses:get:description`),
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

		param, err := collectShippingAddressesListCmdParams()
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

func collectShippingAddressesListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForShippingAddressesListCmd("/operators/{operator_id}/shipping_addresses"),
		query:  buildQueryForShippingAddressesListCmd(),
	}, nil
}

func buildPathForShippingAddressesListCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", ShippingAddressesListCmdOperatorId, -1)

	return path
}

func buildQueryForShippingAddressesListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
