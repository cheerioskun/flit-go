package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var OutputFormat string
var OutputFile string
var rootCmd = &cobra.Command{
	Use:   "flit",
	Short: "Flit is a Go tool to parse and export your fitness log entries into other formats like JSON",
	Long: `Flit is a Go tool to parse and export your fitness log entries into other formats like JSON.
				Full Documentation at github.com/cheerioskun/flit-go/ `,
	Run: func(cmd *cobra.Command, args []string) {},
}

var convertCmd = &cobra.Command{
	Use:   "convert <input_file> FLAGS",
	Short: "Convert log entries in input_file to format specified by the output flag",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse and convert here
		bytes := ConvertFileToFmt(args[0], OutputFormat)
		if OutputFile == "" {
			// Print to stdout
			fmt.Println(string(bytes))
		} else {
			err := os.WriteFile(OutputFile, bytes, 0666)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	convertCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "json", "specify the format for exporting")
	convertCmd.PersistentFlags().StringVarP(&OutputFile, "file", "f", "", "specify where to store the converted log")
	rootCmd.AddCommand(convertCmd)

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
