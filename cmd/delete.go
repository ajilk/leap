package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete bookmark",
	Long:    "delete bookmark",
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("delete called: {%s}\n", args[0])
	},
}

func init() {
	bookmarkCmd.AddCommand(deleteCmd)
}
