package cmd

import (
	"encoding/json"
	"fmt"
	"leap/data"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export saved bookmarks",
	Long:  `export saved bookmarks`,
	Run: func(cmd *cobra.Command, args []string) {
		bookmarks := data.ListBookmarks()
		out, _ := json.Marshal(bookmarks)
		fmt.Println(string(out))
	},
}

func init() {
	bookmarkCmd.AddCommand(exportCmd)
}
