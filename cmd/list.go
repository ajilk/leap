package cmd

import (
	"fmt"
	"leap/data"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List bookmarks",
	Run: func(cmd *cobra.Command, args []string) {
		bookmarks := data.ListBookmarks()
		if len(bookmarks) == 0 {
			fmt.Println("No bookmarks found.")
			return
		}

		t := table.NewWriter()
		t.SetStyle(table.StyleLight)
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(table.Row{"id", "name", "url"})
		for _, value := range bookmarks {
			t.AppendRow(table.Row{value.ID, value.Key, value.Value})
		}

		t.Render()
	},
}

func init() {
	bookmarkCmd.AddCommand(listCmd)
}
