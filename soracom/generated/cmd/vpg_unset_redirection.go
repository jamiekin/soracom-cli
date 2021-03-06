package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgUnsetRedirectionCmdId holds value of 'id' option
var VpgUnsetRedirectionCmdId string

func init() {
	VpgUnsetRedirectionCmd.Flags().StringVar(&VpgUnsetRedirectionCmdId, "id", "", TRAPI("VPG ID"))

	VpgCmd.AddCommand(VpgUnsetRedirectionCmd)
}

// VpgUnsetRedirectionCmd defines 'unset-redirection' subcommand
var VpgUnsetRedirectionCmd = &cobra.Command{
	Use:   "unset-redirection",
	Short: TRAPI("/virtual_private_gateways/{id}/junction/unset_redirection:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{id}/junction/unset_redirection:post:description`),
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

		param, err := collectVpgUnsetRedirectionCmdParams()
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

func collectVpgUnsetRedirectionCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForVpgUnsetRedirectionCmd("/virtual_private_gateways/{id}/junction/unset_redirection"),
		query:  buildQueryForVpgUnsetRedirectionCmd(),
	}, nil
}

func buildPathForVpgUnsetRedirectionCmd(path string) string {

	path = strings.Replace(path, "{"+"id"+"}", VpgUnsetRedirectionCmdId, -1)

	return path
}

func buildQueryForVpgUnsetRedirectionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
