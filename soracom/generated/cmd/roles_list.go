package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// RolesListCmdOperatorId holds value of 'operator_id' option
var RolesListCmdOperatorId string

func init() {
	RolesListCmd.Flags().StringVar(&RolesListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	RolesCmd.AddCommand(RolesListCmd)
}

// RolesListCmd defines 'list' subcommand
var RolesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/roles:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/roles:get:description`),
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

		param, err := collectRolesListCmdParams()
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

func collectRolesListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForRolesListCmd("/operators/{operator_id}/roles"),
		query:  buildQueryForRolesListCmd(),
	}, nil
}

func buildPathForRolesListCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", RolesListCmdOperatorId, -1)

	return path
}

func buildQueryForRolesListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
