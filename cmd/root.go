package cmd

import (
	"fmt"
	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

var (
	Bw = internal.Init()
)

var tagFilter string

var rootCmd = &cobra.Command{
	Use:   "bookworm",
	Short: "Bookworm can manage your bookmarks from the command line.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bookworm!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
