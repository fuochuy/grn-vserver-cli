package networkinterface

import "github.com/vngcloud/greennode-cli/internal/vserverclient"

func init() {
	editCmd.RegisterFlagCompletionFunc("network-interface-id", vserverclient.CompleteNetworkInterfaceIDs)       //nolint:errcheck
	deleteCmd.RegisterFlagCompletionFunc("network-interface-id", vserverclient.CompleteNetworkInterfaceIDs)     //nolint:errcheck
	updateTagsCmd.RegisterFlagCompletionFunc("network-interface-id", vserverclient.CompleteNetworkInterfaceIDs) //nolint:errcheck
	createCmd.RegisterFlagCompletionFunc("zone-id", vserverclient.CompleteZoneIDs)                              //nolint:errcheck
}
