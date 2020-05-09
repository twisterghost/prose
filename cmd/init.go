package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/twisterghost/prose/lib"
	"github.com/twisterghost/prose/loader"
)

var force bool

func initProsefile() {
	if !force && loader.FileExists(loader.GetProsefilePath()) {
		fmt.Println("File already exists as", loader.GetProsefilePath(), "skipping.")
		os.Exit(0)
	}

	initialPostfile := lib.Prosefile{
		Filetype: "prosefile",
		Version:  "0.0.1",
		Entries:  []lib.Entry{},
	}

	loader.WriteProsefile(loader.SerializeProsefile(initialPostfile))
	fmt.Println("Prosefile initialized at", loader.GetProsefilePath())
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a prosefile",
	Long: `Creates a prosefile at the configured location (default $HOME/prosefile.json).

Will not overwrite existing files without the --force flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		initProsefile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwriting existing prosefile")
}
