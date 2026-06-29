package floatingip

import "github.com/vngcloud/greennode-cli/internal/vserverclient"

func init() {
	deleteCmd.RegisterFlagCompletionFunc("floating-ip-id", vserverclient.CompleteFloatingIPIDs) //nolint:errcheck
}
