package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SubscribersListCmdLastEvaluatedKey string

// SubscribersListCmdSerialNumberFilter holds value of 'serial_number_filter' option
var SubscribersListCmdSerialNumberFilter string

// SubscribersListCmdSpeedClassFilter holds value of 'speed_class_filter' option
var SubscribersListCmdSpeedClassFilter string

// SubscribersListCmdStatusFilter holds value of 'status_filter' option
var SubscribersListCmdStatusFilter string

// SubscribersListCmdTagName holds value of 'tag_name' option
var SubscribersListCmdTagName string

// SubscribersListCmdTagValue holds value of 'tag_value' option
var SubscribersListCmdTagValue string

// SubscribersListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var SubscribersListCmdTagValueMatchMode string

// SubscribersListCmdLimit holds value of 'limit' option
var SubscribersListCmdLimit int64

func init() {
	SubscribersListCmd.Flags().StringVar(&SubscribersListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The IMSI of the last subscriber retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next subscriber onward."))

	SubscribersListCmd.Flags().StringVar(&SubscribersListCmdSerialNumberFilter, "serial-number-filter", "", TRAPI("Serial number for filtering the search. Can specify multiple values delimited by `|`. Returns subscribers with serial number starting with the specified value(s)."))

	SubscribersListCmd.Flags().StringVar(&SubscribersListCmdSpeedClassFilter, "speed-class-filter", "", TRAPI("Speed class for filtering the search. Can specify multiple values delimited by `|`. Valid values include: `s1.minimum`, `s1.slow`, `s1.standard`, `s1.fast`"))

	SubscribersListCmd.Flags().StringVar(&SubscribersListCmdStatusFilter, "status-filter", "", TRAPI("Status for filtering the search. Can specify multiple values delimited by `|`. Valid values include: `active`, `inactive`, `ready`, `instock`, `shipped`, `suspended`, and `terminated`."))

	SubscribersListCmd.Flags().StringVar(&SubscribersListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search (exact match)."))

	SubscribersListCmd.Flags().StringVar(&SubscribersListCmdTagValue, "tag-value", "", TRAPI("Tag search string for filtering the search. Required when `tag_name` has been specified."))

	SubscribersListCmd.Flags().StringVar(&SubscribersListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag match mode."))

	SubscribersListCmd.Flags().Int64Var(&SubscribersListCmdLimit, "limit", 0, TRAPI("Maximum number of subscribers to retrieve."))

	SubscribersCmd.AddCommand(SubscribersListCmd)
}

// SubscribersListCmd defines 'list' subcommand
var SubscribersListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/subscribers:get:summary"),
	Long:  TRAPI(`/subscribers:get:description`),
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

		param, err := collectSubscribersListCmdParams()
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

func collectSubscribersListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersListCmd("/subscribers"),
		query:  buildQueryForSubscribersListCmd(),
	}, nil
}

func buildPathForSubscribersListCmd(path string) string {

	return path
}

func buildQueryForSubscribersListCmd() string {
	result := []string{}

	if SubscribersListCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", SubscribersListCmdLastEvaluatedKey))
	}

	if SubscribersListCmdSerialNumberFilter != "" {
		result = append(result, sprintf("%s=%s", "serial_number_filter", SubscribersListCmdSerialNumberFilter))
	}

	if SubscribersListCmdSpeedClassFilter != "" {
		result = append(result, sprintf("%s=%s", "speed_class_filter", SubscribersListCmdSpeedClassFilter))
	}

	if SubscribersListCmdStatusFilter != "" {
		result = append(result, sprintf("%s=%s", "status_filter", SubscribersListCmdStatusFilter))
	}

	if SubscribersListCmdTagName != "" {
		result = append(result, sprintf("%s=%s", "tag_name", SubscribersListCmdTagName))
	}

	if SubscribersListCmdTagValue != "" {
		result = append(result, sprintf("%s=%s", "tag_value", SubscribersListCmdTagValue))
	}

	if SubscribersListCmdTagValueMatchMode != "" {
		result = append(result, sprintf("%s=%s", "tag_value_match_mode", SubscribersListCmdTagValueMatchMode))
	}

	if SubscribersListCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", SubscribersListCmdLimit))
	}

	return strings.Join(result, "&")
}
