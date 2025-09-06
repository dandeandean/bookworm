package cmd

import (
	"os"
	"slices"

	"fmt"

	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

func prGetCfg(_ *cobra.Command, _ []string) error {
	var err error
	Bw, err = internal.Get()
	if err != nil {
		fmt.Println("Couldn't get Config, please run bookworm init")
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

func prInitCfg(cmd *cobra.Command, args []string) error {
	var err error
	Bw, err = internal.Init()
	if err != nil {
		fmt.Println("Failed to create config: ", err)
		os.Exit(1)
		return err
	}
	return nil
}

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
	out := make([]string, len(Bw.BookMarks))
	for _, b := range Bw.BookMarks {
		for _, tag := range b.Tags {
			if !slices.Contains(args, tag) {
				out = append(out, tag)
			}
		}
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}

func getNamesCmp(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nonCmp(cmd, args, toComplete)
	}
	out := make([]string, len(Bw.BookMarks))
	for k := range Bw.BookMarks {
		out = append(out, k)
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}
