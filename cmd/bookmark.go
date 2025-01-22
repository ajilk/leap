/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"leap/data"

	"github.com/spf13/cobra"
)

// bookmarkCmd represents the bookmark command
var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateBookmarksTable()
		data.InsertAndFetch()
	},
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookmarkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookmarkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
