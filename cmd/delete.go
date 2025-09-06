package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(delCmd)
}

var delCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Delete no good bookmarks.",
	Args:              cobra.ExactArgs(1),
	Aliases:           []string{"rm"},
	PreRunE:           prGetCfg,
	ValidArgsFunction: getNamesCmp,
	Run: func(cmd *cobra.Command, args []string) {
		Bw.DeleteBookMark(args[0])
	},
}
