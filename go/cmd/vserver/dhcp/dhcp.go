package dhcp

import (
	"github.com/spf13/cobra"
)

// DhcpCmd is the parent command for all DHCP option subcommands.
var DhcpCmd = &cobra.Command{
	Use:   "dhcp",
	Short: "Manage DHCP options",
	Long:  "List and delete DHCP option sets.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DhcpCmd.AddCommand(listCmd)
	DhcpCmd.AddCommand(deleteCmd)
}
