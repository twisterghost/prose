package cmd

import (
	"fmt"
	"os"

	"github.com/twisterghost/prose/loader"

	"github.com/spf13/cobra"
)

func format() {
	prosefile := loader.LoadProsefile()
	err := prosefile.Format()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	loader.WriteProsefile(loader.SerializeProsefile(prosefile))
	fmt.Println("Prosefile at", loader.GetProsefilePath(), "formatted.")
}

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Load, format and save your prosefile back to its original location",
	Long: `Lods your prosefile (default ~/prosefile.json), ensures proper formatting
and saves the contents back to the file. Optionally prettifies content.`,
	Run: func(cmd *cobra.Command, args []string) {
		format()
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
