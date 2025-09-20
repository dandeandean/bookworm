package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

var (
	// Global BookWorm
	Bw          *internal.BookWorm
	tagFilter   string
	verboseMode bool
)

var rootCmd = &cobra.Command{
	Use:     "bookworm",
	Short:   "Bookworm can manage your bookmarks from the command line.",
	PreRunE: prGetCfg,
	Run: func(cmd *cobra.Command, args []string) {
		m := TeaModel()
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verboseMode, "verbose", "v", false, "Enable verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Something Horrible Happened!", err)
	}
}
