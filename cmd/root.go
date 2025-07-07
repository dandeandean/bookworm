package cmd

import (
	"fmt"
	"slices"

	. "github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

var (
	Bw = Init()
)

var tagFilter string

func init() {
	rootCmd.AddCommand(makeCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(tagCmd)

	listCmd.PersistentFlags().StringVarP(&tagFilter, "tag", "t", "", "tag filter exact match")
	listCmd.RegisterFlagCompletionFunc("tag", getTagsCmp)
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
		err := OpenURL(bm.Link)
		if err != nil {
			panic(err)
		}
		Bw.SetLastOpened(*bm)
	},
}
var makeCmd = &cobra.Command{
	Use:               "make",
	Short:             "Make new bookmarks.",
	Args:              cobra.ExactArgs(2),
	Aliases:           []string{"mk", "new"},
	ValidArgsFunction: nonCmp,
	Run: func(cmd *cobra.Command, args []string) {
		if !IsValidUrl(args[1]) {
			fmt.Println(args[1] + " is not a valid URL. What are you doing?")
		}
		Bw.NewBookMark(args[0], args[1], []string{})
	},
}

var delCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Delete no good bookmarks.",
	Args:              cobra.ExactArgs(1),
	Aliases:           []string{"rm"},
	ValidArgsFunction: getNamesCmp,
	Run: func(cmd *cobra.Command, args []string) {
		Bw.DeleteBookMark(args[0])
	},
}

var rootCmd = &cobra.Command{
	Use:   "bookworm",
	Short: "Bookworm can manage your bookmarks from the command line.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bookworm!")
	},
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
