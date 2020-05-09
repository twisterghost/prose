package cmd

import (
	"fmt"

	"github.com/twisterghost/prose/lib"
	"github.com/twisterghost/prose/loader"

	"github.com/spf13/cobra"
)

func send() {
	prosefile := loader.LoadProsefile()

	for index, entry := range prosefile.Entries {
		prosefile.Entries[index] = lib.FormatEntry(entry)
	}

	fmt.Printf(loader.SerializeProsefile(prosefile))
}

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Load, format and send entries to stdout",
	Long: `Loads your prosefile (default ~/prosefile.json), ensures proper formatting
and prints the contents to stdout to be piped to a modifier or renderer.`,
	Run: func(cmd *cobra.Command, args []string) {
		send()
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
