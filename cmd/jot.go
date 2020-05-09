package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"

	"github.com/spf13/viper"

	"github.com/twisterghost/prose/lib"
	"github.com/twisterghost/prose/loader"
)

func jot(args []string) {
	message := strings.Join(args, " ")
	author := viper.GetString("author")

	prosefile := loader.LoadProsefile()
	newEntry := lib.NewBasicEntry(message, author)

	prosefile.Entries = append(prosefile.Entries, newEntry)

	loader.WriteProsefile(loader.SerializeProsefile(prosefile))

	fmt.Println("Message saved:", message)
}

// jotCmd represents the jot command
var jotCmd = &cobra.Command{
	Use:   "jot <words...>",
	Short: "Add a simple (titleless) entry.",
	Long:  `Add a simple entry with just a body message to the prosefile.`,
	Run: func(cmd *cobra.Command, args []string) {
		jot(args)
	},
}

func init() {
	rootCmd.AddCommand(jotCmd)
}
