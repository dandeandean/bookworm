package cmd

import (
	"github.com/spf13/cobra"
	"slices"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&tagFilter, "tag", "t", "", "tag filter exact match")
	listCmd.RegisterFlagCompletionFunc("tag", getTagsCmp)
}

// Inspo `gh repo list`
var listCmd = &cobra.Command{
	Use:               "list",
	Args:              cobra.ExactArgs(0),
	Short:             "List your bookmarks.",
	Aliases:           []string{"ls"},
	ValidArgsFunction: nonCmp,
	Run: func(cmd *cobra.Command, args []string) {
		for _, b := range Bw.Cfg.BookMarks {
			if tagFilter == "" || slices.Contains(b.Tags, tagFilter) {
				b.Println()
			}
		}
	},
}
