package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	BillsCmd.AddCommand(BillsGetLatestCmd)
}

// BillsGetLatestCmd defines 'get-latest' subcommand
var BillsGetLatestCmd = &cobra.Command{
	Use:   "get-latest",
	Short: TRAPI("/bills/latest:get:summary"),
	Long:  TRAPI(`/bills/latest:get:description`),
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

		param, err := collectBillsGetLatestCmdParams()
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

func collectBillsGetLatestCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsGetLatestCmd("/bills/latest"),
		query:  buildQueryForBillsGetLatestCmd(),
	}, nil
}

func buildPathForBillsGetLatestCmd(path string) string {

	return path
}

func buildQueryForBillsGetLatestCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
