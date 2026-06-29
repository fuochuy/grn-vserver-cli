package server

import (
	"github.com/spf13/cobra"
	"github.com/vngcloud/greennode-cli/internal/vserverclient"
)

func init() {
	// --server-id on all commands that target an existing server
	getCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs)    //nolint:errcheck
	startCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs)  //nolint:errcheck
	stopCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs)   //nolint:errcheck
	rebootCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs) //nolint:errcheck
	deleteCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs) //nolint:errcheck
	resizeCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs) //nolint:errcheck

	// update-secgroup: target server + the security groups to attach
	updateSecgroupCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs)        //nolint:errcheck
	updateSecgroupCmd.RegisterFlagCompletionFunc("security-group", vserverclient.CompleteSecgroupIDs) //nolint:errcheck

	// create-image: target server; tag-value: which tag key to inspect
	createImageCmd.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs) //nolint:errcheck
	tagValueCmd.RegisterFlagCompletionFunc("key", vserverclient.CompleteTagKeys)            //nolint:errcheck

	// attach/detach floating IP: target server, the floating IP, and the network interface
	for _, c := range []*cobra.Command{attachFloatingIPCmd, detachFloatingIPCmd} {
		c.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs)                      //nolint:errcheck
		c.RegisterFlagCompletionFunc("floating-ip-id", vserverclient.CompleteFloatingIPIDs)             //nolint:errcheck
		c.RegisterFlagCompletionFunc("network-interface-id", vserverclient.CompleteNetworkInterfaceIDs) //nolint:errcheck
	}

	// attach/detach internal & external network interfaces: all target a server
	for _, c := range []*cobra.Command{attachInternalInterfaceCmd, detachInternalInterfaceCmd, attachExternalInterfaceCmd, detachExternalInterfaceCmd} {
		c.RegisterFlagCompletionFunc("server-id", vserverclient.CompleteServerIDs) //nolint:errcheck
	}
	// external interfaces are elastic network interfaces; complete from that list
	attachExternalInterfaceCmd.RegisterFlagCompletionFunc("network-interface-id", vserverclient.CompleteNetworkInterfaceIDs) //nolint:errcheck
	detachExternalInterfaceCmd.RegisterFlagCompletionFunc("network-interface-id", vserverclient.CompleteNetworkInterfaceIDs) //nolint:errcheck

	// create: zone, network, subnet, image, volume types, security group
	createCmd.RegisterFlagCompletionFunc("zone-id", vserverclient.CompleteZoneIDs)                 //nolint:errcheck
	createCmd.RegisterFlagCompletionFunc("network-id", vserverclient.CompleteVPCIDs)               //nolint:errcheck
	createCmd.RegisterFlagCompletionFunc("subnet-id", vserverclient.CompleteSubnetIDs)             //nolint:errcheck
	createCmd.RegisterFlagCompletionFunc("image-id", vserverclient.CompleteImageIDs)               //nolint:errcheck
	createCmd.RegisterFlagCompletionFunc("root-disk-type-id", vserverclient.CompleteVolumeTypeIDs) //nolint:errcheck
	createCmd.RegisterFlagCompletionFunc("data-disk-type-id", vserverclient.CompleteVolumeTypeIDs) //nolint:errcheck
	createCmd.RegisterFlagCompletionFunc("security-group", vserverclient.CompleteSecgroupIDs)      //nolint:errcheck
}
