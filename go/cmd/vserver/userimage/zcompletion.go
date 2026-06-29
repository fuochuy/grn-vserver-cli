package userimage

import "github.com/vngcloud/greennode-cli/internal/vserverclient"

func init() {
	deleteCmd.RegisterFlagCompletionFunc("user-image-id", vserverclient.CompleteUserImageIDs) //nolint:errcheck
}
