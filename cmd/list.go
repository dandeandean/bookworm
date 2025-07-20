package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&tagFilter, "tag", "t", "", "tag filter exact match")
	listCmd.RegisterFlagCompletionFunc("tag", getTagsCmp)
}

// Inspo `gh repo list`
var listCmd = &cobra.Command{
	Use:               "list",
	Args:              cobra.ExactArgs(0),
	Short:             "List your bookmarks.",
	Aliases:           []string{"ls"},
	ValidArgsFunction: nonCmp,
	Run: func(cmd *cobra.Command, args []string) {
		// for _, b := range Bw.Cfg.BookMarks {
		// 	if tagFilter == "" || slices.Contains(b.Tags, tagFilter) {
		// 		b.Println()
		// 	}
		// }
		choices := make([]string, 0, len(Bw.Cfg.BookMarks))
		for k := range Bw.Cfg.BookMarks {
			choices = append(choices, k)
		}
		m := modelFrom(choices)
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}
