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
	noFzf       bool
)

var rootCmd = &cobra.Command{
	Use:     "bookworm",
	Short:   "Bookworm can manage your bookmarks from the command line.",
	PreRunE: prGetCfg,
	RunE: func(cmd *cobra.Command, args []string) error {
		if Bw.Cfg.FzfIntegration && !noFzf {
			return Bw.FzfOpen("")
		} else {
			m := TeaModel()
			p := tea.NewProgram(m)
			if _, err := p.Run(); err != nil {
				if verboseMode {
					fmt.Println("Failed to run bubbletea!")
				}
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.Flags().BoolVar(&noFzf, "no-fzf", false, "skip the fzf integration")
	rootCmd.PersistentFlags().BoolVarP(&verboseMode, "verbose", "v", false, "Enable verbose output")
	rootCmd.SilenceErrors = true
	if verboseMode {
		rootCmd.SilenceUsage = true
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed to execute the root command!")
		if !verboseMode {
			fmt.Println("Use the --verbose flag for more information ")
		}
		if verboseMode {
			fmt.Println(err)
		}
	}
}
