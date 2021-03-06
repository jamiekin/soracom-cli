package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgCreateMirroringPeerCmdDescription holds value of 'description' option
var VpgCreateMirroringPeerCmdDescription string

// VpgCreateMirroringPeerCmdId holds value of 'id' option
var VpgCreateMirroringPeerCmdId string

// VpgCreateMirroringPeerCmdIpAddress holds value of 'ipAddress' option
var VpgCreateMirroringPeerCmdIpAddress string

// VpgCreateMirroringPeerCmdProtocol holds value of 'protocol' option
var VpgCreateMirroringPeerCmdProtocol string

// VpgCreateMirroringPeerCmdEnabled holds value of 'enabled' option
var VpgCreateMirroringPeerCmdEnabled bool

// VpgCreateMirroringPeerCmdBody holds contents of request body to be sent
var VpgCreateMirroringPeerCmdBody string

func init() {
	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdDescription, "description", "", TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdId, "id", "", TRAPI("VPG ID"))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdIpAddress, "ip-address", "", TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdProtocol, "protocol", "", TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().BoolVar(&VpgCreateMirroringPeerCmdEnabled, "enabled", false, TRAPI(""))

	VpgCreateMirroringPeerCmd.Flags().StringVar(&VpgCreateMirroringPeerCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgCreateMirroringPeerCmd)
}

// VpgCreateMirroringPeerCmd defines 'create-mirroring-peer' subcommand
var VpgCreateMirroringPeerCmd = &cobra.Command{
	Use:   "create-mirroring-peer",
	Short: TRAPI("/virtual_private_gateways/{id}/junction/mirroring/peers:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{id}/junction/mirroring/peers:post:description`),
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

		param, err := collectVpgCreateMirroringPeerCmdParams()
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

func collectVpgCreateMirroringPeerCmdParams() (*apiParams, error) {

	body, err := buildBodyForVpgCreateMirroringPeerCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgCreateMirroringPeerCmd("/virtual_private_gateways/{id}/junction/mirroring/peers"),
		query:       buildQueryForVpgCreateMirroringPeerCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgCreateMirroringPeerCmd(path string) string {

	path = strings.Replace(path, "{"+"id"+"}", VpgCreateMirroringPeerCmdId, -1)

	return path
}

func buildQueryForVpgCreateMirroringPeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgCreateMirroringPeerCmd() (string, error) {
	if VpgCreateMirroringPeerCmdBody != "" {
		if strings.HasPrefix(VpgCreateMirroringPeerCmdBody, "@") {
			fname := strings.TrimPrefix(VpgCreateMirroringPeerCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgCreateMirroringPeerCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgCreateMirroringPeerCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if VpgCreateMirroringPeerCmdDescription != "" {
		result["description"] = VpgCreateMirroringPeerCmdDescription
	}

	if VpgCreateMirroringPeerCmdIpAddress != "" {
		result["ipAddress"] = VpgCreateMirroringPeerCmdIpAddress
	}

	if VpgCreateMirroringPeerCmdProtocol != "" {
		result["protocol"] = VpgCreateMirroringPeerCmdProtocol
	}

	if VpgCreateMirroringPeerCmdEnabled != false {
		result["enabled"] = VpgCreateMirroringPeerCmdEnabled
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
