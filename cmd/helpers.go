package cmd

import "github.com/spf13/cobra"

func getNamesCmp(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	out := make([]string, len(Bw.Cfg.BookMarks))
	for _, b := range Bw.Cfg.BookMarks {
		out = append(out, b.Name)
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}
