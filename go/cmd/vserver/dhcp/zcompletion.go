package dhcp

import "github.com/vngcloud/greennode-cli/internal/vserverclient"

func init() {
	deleteCmd.RegisterFlagCompletionFunc("dhcp-option-id", vserverclient.CompleteDhcpOptionIDs) //nolint:errcheck
}
