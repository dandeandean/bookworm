package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var outFile string

func init() {
	exportCmd.PersistentFlags().StringVarP(&outFile, "out", "o", "", "write to output file")
	exportCmd.RegisterFlagCompletionFunc("out", nonCmp)
	rootCmd.AddCommand(exportCmd)

}

var exportCmd = &cobra.Command{
	Use:               "export",
	Short:             "Export to your favorite serialization format, so long as it's json.",
	Args:              cobra.ExactArgs(1),
	Aliases:           []string{"jsonify"},
	PreRunE:           prGetCfg,
	ValidArgsFunction: getNamesCmp,
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := Bw.GetOneRaw(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Print(string(bytes))
	},
}

