package placementgroup

import "github.com/vngcloud/greennode-cli/internal/vserverclient"

func init() {
	deleteCmd.RegisterFlagCompletionFunc("placement-group-id", vserverclient.CompletePlacementGroupIDs) //nolint:errcheck
	editCmd.RegisterFlagCompletionFunc("placement-group-id", vserverclient.CompletePlacementGroupIDs)   //nolint:errcheck
}
