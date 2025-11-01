package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dandeandean/bookworm/internal"
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
		action := func(bytes []byte, err error) {
			if err != nil {
				panic("Error with reading the JSON data")
			}
			if outFile == "" {
				fmt.Print(string(bytes))
				return
			}
			err = os.WriteFile(outFile, bytes, 0666)
			if err != nil {
				panic("Error writing the file to disk")
			}
		}
		switch len(args) {
		case 0:
			bytes, err := Bw.GetAllRaw()
			if err != nil {
				panic(err)
			}
			strBM := make(map[string]internal.BookMark)
			for k, v := range bytes {
				bmCask := &internal.BookMark{}
				json.Unmarshal(v, bmCask)
				strBM[k] = *bmCask
			}
			action(json.Marshal(strBM))
		case 1:
			action(Bw.GetOneRaw(args[0]))
		}
	},
}
