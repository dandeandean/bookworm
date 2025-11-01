package cmd

import (
	"encoding/json"
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
	Args:              cobra.MinimumNArgs(0),
	Aliases:           []string{"jsonify"},
	PreRunE:           prGetCfg,
	ValidArgsFunction: getNamesCmp,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			bytes, err := Bw.GetAllRaw()
			if err != nil {
				panic(err)
			}
			encoded, err := json.Marshal(bytes)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(encoded))
		case 1:
			bytes, err := Bw.GetOneRaw(args[0])
			if err != nil {
				panic(err)
			}
			fmt.Print(string(bytes))
		}
	},
}
