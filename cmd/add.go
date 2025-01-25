package cmd

import (
	"fmt"
	"leap/data"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new bookmark",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Error: You must provide a <key> and <value>")
			cmd.Help()
			return
		}

		key := args[0]
		value := args[1]

		data.InsertBookmark(key, value)
		fmt.Printf("Bookmark saved: %s -> %s\n", key, value)
	},
}

func init() {
	bookmarkCmd.AddCommand(addCmd)
}
