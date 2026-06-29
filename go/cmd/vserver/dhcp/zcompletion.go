package dhcp

import "github.com/vngcloud/greennode-cli/internal/vserverclient"

func init() {
	deleteCmd.RegisterFlagCompletionFunc("dhcp-option-id", vserverclient.CompleteDhcpOptionIDs)       //nolint:errcheck
	getCmd.RegisterFlagCompletionFunc("dhcp-option-id", vserverclient.CompleteDhcpOptionIDs)          //nolint:errcheck
	listVpcsCmd.RegisterFlagCompletionFunc("dhcp-option-id", vserverclient.CompleteDhcpOptionIDs)     //nolint:errcheck
	associateVpcCmd.RegisterFlagCompletionFunc("vpc-id", vserverclient.CompleteVPCIDs)                //nolint:errcheck
	associateVpcCmd.RegisterFlagCompletionFunc("dhcp-option-id", vserverclient.CompleteDhcpOptionIDs) //nolint:errcheck
}
