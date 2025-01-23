package cmd

import (
	"leap/data"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize leap database",
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateBookmarksTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
