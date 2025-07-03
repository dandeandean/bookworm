package cmd

import (
	"slices"

	"github.com/spf13/cobra"
)

func nonCmp(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
	return []string{}, cobra.ShellCompDirectiveNoFileComp
}

func getNamesThenTagsCmp(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) == 0 {
		return getNamesCmp(cmd, args, toComplete)
	}
	return getTagsCmp(cmd, args, toComplete)
}

func getTagsCmp(_ *cobra.Command, args []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
	out := make([]string, len(Bw.Cfg.BookMarks))
	for _, b := range Bw.Cfg.BookMarks {
		for _, t := range b.Tags {
			if !slices.Contains(args, t) {
				out = append(out, t)
			}
		}
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}

func getNamesCmp(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	out := make([]string, len(Bw.Cfg.BookMarks))
	for _, b := range Bw.Cfg.BookMarks {
		out = append(out, b.Name)
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}
