package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/twisterghost/prose/loader"
	"github.com/twisterghost/prose/prose"
)

var force bool

func initProsefile() {
	if !force && loader.FileExists(loader.GetProsefilePath()) {
		fmt.Println("File already exists as", loader.GetProsefilePath(), "skipping.")
		os.Exit(0)
	}

	initialPostfile := prose.Prosefile{
		Filetype: "prosefile",
		Version:  "0.0.1",
		Entries:  []prose.Entry{},
	}

	loader.WriteProsefile(initialPostfile)
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
