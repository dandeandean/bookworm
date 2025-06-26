package cmd

import (
	"fmt"

	. "github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

var (
	Bw = Init()
)

func init() {
	rootCmd.AddCommand(makeCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(listCmd)
}

var openCmd = &cobra.Command{
	Use:     "open",
	Short:   "Open a Bookmark.",
	Aliases: []string{"go"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bm, ok := Bw.Cfg.BookMarks[args[0]]
		if !ok {
			fmt.Println("Couldn't Find BookMark!")
			return
		}
		err := OpenURL(bm.Link)
		if err != nil {
			panic(err)
		}
		Bw.SetLastOpened(bm)
	},
}

var makeCmd = &cobra.Command{
	Use:     "make",
	Short:   "Make New Bookmarks.",
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"mk", "new"},
	Run: func(cmd *cobra.Command, args []string) {
		if !IsValidUrl(args[1]) {
			fmt.Println(args[1] + " is not a valid URL. What are you doing?")
		}
		Bw.NewBookMark(args[0], args[1])
	},
}

var rootCmd = &cobra.Command{
	Use:   "bookworm",
	Short: "Bookworm Can Manage Your Bookmarks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bookworm!")
	},
}

var listCmd = &cobra.Command{
	Use:     "list",
	Args:    cobra.ExactArgs(0),
	Short:   "Bookworm Can Manage Your Bookmarks.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		for _, b := range Bw.Cfg.BookMarks {
			b.Println()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
