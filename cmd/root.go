package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

// Global BookWorm object
var Bw *internal.BookWorm

func init() {
	var err error
	Bw, err = internal.Get()
	if err != nil {
		fmt.Println("Couldn't get Config, please run bookworm init")
		os.Exit(2)
	}
}

var tagFilter string

var rootCmd = &cobra.Command{
	Use:   "bookworm",
	Short: "Bookworm can manage your bookmarks from the command line.",
	Run: func(cmd *cobra.Command, args []string) {
		m := TeaModel()
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
