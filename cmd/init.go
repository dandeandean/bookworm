package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize Bookworm.",
	PreRunE: prInitCfg,
	Run: func(cmd *cobra.Command, args []string) {
		if Bw == nil {
			panic("BW object is nil")
		}
		fmt.Println("Successfully Initialized")
	},
}
