package placementgroup

import (
	"github.com/spf13/cobra"
)

// PlacementGroupCmd is the parent command for all placement group subcommands.
var PlacementGroupCmd = &cobra.Command{
	Use:   "placement-group",
	Short: "Manage placement groups (server groups)",
	Long:  "List and delete placement groups.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	PlacementGroupCmd.AddCommand(listCmd)
	PlacementGroupCmd.AddCommand(deleteCmd)
}
