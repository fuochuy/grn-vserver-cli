package sshkey

import (
	"github.com/spf13/cobra"
	"github.com/vngcloud/greennode-cli/internal/vserverclient"
)

func init() {
	deleteCmd.RegisterFlagCompletionFunc("sshkey-id", vserverclient.CompleteSSHKeyIDs) //nolint:errcheck

	// Complete --public-key-file with .pub files (falls back to all files).
	importCmd.RegisterFlagCompletionFunc("public-key-file", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) { //nolint:errcheck
		return []string{"pub"}, cobra.ShellCompDirectiveFilterFileExt
	})
}
