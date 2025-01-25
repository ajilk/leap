package cmd

import (
	"github.com/spf13/cobra"
)

var bookmarkCmd = &cobra.Command{
	Use:     "bookmark",
	Short:   "Bookmark <key> <value>",
	Long:    `Bookmark command allows you to save and retrieve bookmarks with a key-value pair.`,
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		helpFlag, _ := cmd.Flags().GetBool("help")
		if len(args) == 0 || helpFlag {
			cmd.Help() // Display the help message
		}
	},
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)

	bookmarkCmd.Flags().BoolP("help", "h", false, "Display help for bookmark command")
}
