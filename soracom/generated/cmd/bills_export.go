package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// BillsExportCmdExportMode holds value of 'export_mode' option
var BillsExportCmdExportMode string

// BillsExportCmdYyyyMM holds value of 'yyyyMM' option
var BillsExportCmdYyyyMM string

func init() {
	BillsExportCmd.Flags().StringVar(&BillsExportCmdExportMode, "export-mode", "", TRAPI("export_mode (async, sync)"))

	BillsExportCmd.Flags().StringVar(&BillsExportCmdYyyyMM, "yyyy-mm", "", TRAPI("yyyyMM"))

	BillsCmd.AddCommand(BillsExportCmd)
}

// BillsExportCmd defines 'export' subcommand
var BillsExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/bills/{yyyyMM}/export:post:summary"),
	Long:  TRAPI(`/bills/{yyyyMM}/export:post:description`),
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

		param, err := collectBillsExportCmdParams()
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

func collectBillsExportCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForBillsExportCmd("/bills/{yyyyMM}/export"),
		query:  buildQueryForBillsExportCmd(),
	}, nil
}

func buildPathForBillsExportCmd(path string) string {

	path = strings.Replace(path, "{"+"yyyyMM"+"}", BillsExportCmdYyyyMM, -1)

	return path
}

func buildQueryForBillsExportCmd() string {
	result := []string{}

	if BillsExportCmdExportMode != "" {
		result = append(result, sprintf("%s=%s", "export_mode", BillsExportCmdExportMode))
	}

	return strings.Join(result, "&")
}
