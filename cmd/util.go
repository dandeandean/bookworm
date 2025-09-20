package cmd

import (
	"fmt"
	"os"
	"slices"

	"github.com/dandeandean/bookworm/internal"
	"github.com/spf13/cobra"
)

func prGetCfg(_ *cobra.Command, _ []string) error {
	if Bw != nil {
		return nil
	}
	var err error
	Bw, err = internal.Get()
	if err != nil {
		fmt.Println("Couldn't get Config, please run bookworm init!")
		os.Exit(2)
		return err
	}
	return nil
}

func prInitCfg(cmd *cobra.Command, args []string) error {
	var err error
	Bw, err = internal.Init()
	if err != nil {
		fmt.Println("Failed to create config!")
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

func getTagsCmp(cmd *cobra.Command, args []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
	prGetCfg(cmd, args)
	bms := Bw.ListBookMarks("")
	out := make([]string, len(bms))
	for _, b := range bms {
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
	prGetCfg(cmd, args)
	bms := Bw.ListBookMarks("")
	out := make([]string, 0)
	for _, k := range bms {
		out = append(out, k.Name)
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}
