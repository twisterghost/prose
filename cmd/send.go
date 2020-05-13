package cmd

import (
	"fmt"
	"os"

	"github.com/twisterghost/prose/loader"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func send() {
	prosefile := loader.LoadProsefile()

	err := prosefile.Format()

	if err != nil {
		fmt.Println("Error formatting file")
		fmt.Println(err)
		os.Exit(1)
	}

	pretty := viper.GetBool("pretty")
	fmt.Printf(prosefile.Serialize(pretty))
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
