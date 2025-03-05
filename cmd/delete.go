package cmd

import (
	"fmt"
	"leap/data"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete bookmark",
	Long:    "delete bookmark",
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		data.DeleteBookmark(id)

		fmt.Printf("Deleted: %d\n", id)
	},
}

func init() {
	bookmarkCmd.AddCommand(deleteCmd)
}
