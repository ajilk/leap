package cmd

import (
	"leap/data"

	"github.com/spf13/cobra"
)

// wipeCmd represents the wipe command
var wipeCmd = &cobra.Command{
	Use:   "wipe",
	Short: "wipe out all bookmarks",
	Run: func(cmd *cobra.Command, args []string) {
		data.DeleteAllBookmarks()
	},
}

func init() {
	bookmarkCmd.AddCommand(wipeCmd)
}
