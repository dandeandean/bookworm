package cmd

import (
	"fmt"
	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(openCmd)
}

var openCmd = &cobra.Command{
	Use:               "open",
	Short:             "Open a Bookmark.",
	Aliases:           []string{"go"},
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getNamesCmp,
	Run: func(cmd *cobra.Command, args []string) {
		bm, ok := Bw.Cfg.BookMarks[args[0]]
		if !ok {
			fmt.Println("Couldn't Find BookMark!")
			return
		}
		err := internal.OpenURL(bm.Link)
		if err != nil {
			panic(err)
		}
		Bw.SetLastOpened(*bm)
	},
}
