package cmd

import (
	"fmt"
	"leap/data"
	"os"
	"text/template"

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

		templateContent := `{{range $index, $bookmark := .}}{{if $index}}
{{end}}{{.Key}} - {{.Value}}{{end}}
`

		tmpl, err := template.New("bookmarks").Parse(templateContent)
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			return
		}

		err = tmpl.Execute(os.Stdout, bookmarks)
		if err != nil {
			fmt.Printf("Error rendering template: %v\n", err)
		}
	},
}

func init() {
	bookmarkCmd.AddCommand(listCmd)
}
