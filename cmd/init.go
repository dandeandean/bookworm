package cmd

import (
	"fmt"
	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Bookworm.",
	Run: func(cmd *cobra.Command, args []string) {
		Bw = internal.Init()
		fmt.Println("todo")
		os.Exit(1)
	},
}
