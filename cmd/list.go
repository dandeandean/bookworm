package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&tagFilter, "tag", "t", "", "tag filter exact match")
	listCmd.RegisterFlagCompletionFunc("tag", getTagsCmp)
}

var listCmd = &cobra.Command{
	Use:               "list",
	Args:              cobra.ExactArgs(0),
	Short:             "List your bookmarks.",
	Aliases:           []string{"ls"},
	PreRunE:           prGetCfg,
	ValidArgsFunction: nonCmp,
	Run: func(cmd *cobra.Command, args []string) {
		for _, b := range Bw.ListBookMarks(tagFilter) {
			b.Println()
		}
	},
}
