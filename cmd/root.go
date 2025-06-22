package cmd

import (
	"fmt"

	. "github.com/dandeandean/bookworm/config"
	"github.com/spf13/cobra"
)

var (
	Cfg = GetConfig()
)

func init() {
	rootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Make New Bookmarks.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		Cfg.NewBookMark(args[0], args[1])
	},
}

var rootCmd = &cobra.Command{
	Use:   "bookworm",
	Short: "Bookworm Can Manage Your Bookmarks.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, b := range Cfg.BookMarks {
			b.Println()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
