package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tagCmd)
}

var tagCmd = &cobra.Command{
	Use:               "tag name tags...",
	Args:              cobra.MinimumNArgs(2),
	Short:             "Tag boomarks with tags.",
	ValidArgsFunction: getNamesThenTagsCmp,
	Run: func(cmd *cobra.Command, args []string) {
		err := Bw.SetTags(args[0], args[1:])
		if err != nil {
			fmt.Println(err)
		}
	},
}
