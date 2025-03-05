package cmd

import (
	"slices"

	"github.com/spf13/cobra"
)

var bookmarkCmd = &cobra.Command{
	Use:     "bookmark",
	Short:   "Bookmark <key> <value>",
	Long:    `Bookmark command allows you to save and retrieve bookmarks with a key-value pair.`,
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		helpFlag, _ := cmd.Flags().GetBool("help")
		if helpFlag {
			cmd.Help() // Display the help message
		}

		if len(args) >= 2 {
			id := args[0]
			action := args[1]

			switch {
			case action == deleteCmd.Name() || slices.Contains(deleteCmd.Aliases, action):
				deleteCmd.Run(deleteCmd, []string{id})
			}
		} else {
			listCmd.Run(listCmd, []string{})
		}
	},
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)

	bookmarkCmd.Flags().BoolP("help", "h", false, "Display help for bookmark command")
}
