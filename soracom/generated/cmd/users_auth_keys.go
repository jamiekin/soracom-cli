package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	UsersCmd.AddCommand(UsersAuthKeysCmd)
}

// UsersAuthKeysCmd defines 'auth-keys' subcommand
var UsersAuthKeysCmd = &cobra.Command{
	Use:   "auth-keys",
	Short: TRCLI("cli.users.auth-keys.summary"),
	Long:  TRCLI(`cli.users.auth-keys.description`),
}
