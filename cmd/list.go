package cmd

import (
	"fmt"
	"leap/data"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List bookmarks",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch bookmarks
		bookmarks := data.ListBookmarks()
		if len(bookmarks) == 0 {
			fmt.Println("No bookmarks found.")
			return
		}

    templateContent := `{{range $index, $bookmark := .}}{{if $index}}
{{end}}{{.Key}} - {{.Value}}{{end}}
`

		// Parse the template
		tmpl, err := template.New("bookmarks").Parse(templateContent)
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			return
		}

		// Execute the template
		err = tmpl.Execute(os.Stdout, bookmarks)
		if err != nil {
			fmt.Printf("Error rendering template: %v\n", err)
		}

		// // Get the path of the template file
		// execPath, err := os.Executable()
		// if err != nil {
		// 	fmt.Printf("Error: %v\n", err)
		// 	return
		// }
		// templatePath := filepath.Join(filepath.Dir(execPath), "templates", "bookmarks.go.tmpl")
		//
		// // Check if the template file exists
		// if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		// 	fmt.Printf("Template file not found: %s\n", templatePath)
		// 	return
		// }
		//
		// // Parse the template file
		// fmt.Println(templatePath)
		// t, err := template.New("leap").ParseFiles(templatePath)
		// if err != nil {
		// 	fmt.Printf("Error parsing template: %v\n", err)
		// 	return
		// }
		//
		// // Render the template with bookmarks
		// err = t.Execute(os.Stdout, bookmarks)
		// if err != nil {
		// 	fmt.Printf("Error rendering template: %v\n", err)
		// }
	},
}

func init() {
	bookmarkCmd.AddCommand(listCmd)
}
