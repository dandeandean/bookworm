package cmd

import (
	"fmt"
	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use:               "make",
	Short:             "Make new bookmarks.",
	Args:              cobra.ExactArgs(2),
	Aliases:           []string{"mk", "new"},
	PreRunE:           prGetCfg,
	ValidArgsFunction: nonCmp,
	Run: func(cmd *cobra.Command, args []string) {
		if !internal.IsValidUrl(args[1]) {
			fmt.Println(args[1] + " is not a valid URL. What are you doing?")
		}
		err := Bw.NewBookMark(args[0], args[1], []string{})
		if err != nil {
			fmt.Println(err)
		}
	},
}
