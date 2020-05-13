package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"

	"github.com/twisterghost/prose/loader"
	"github.com/twisterghost/prose/prose"
)

var Title string

func addEntry(path string) {

	if !loader.FileExists(path) {
		fmt.Println("File does not exist:", path)
		os.Exit(1)
	}

	author := viper.GetString("author")

	message, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading input file at", path)
		fmt.Println(err)
		os.Exit(1)
	}

	prosefile := loader.LoadProsefile()

	newEntry := prose.NewBasicEntry(string(message), author)
	newEntry.Title = Title
	prosefile.Entries = append(prosefile.Entries, newEntry)

	loader.WriteProsefile(prosefile)

	fmt.Println("Entry added:", string(message))
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a file to your entries",
	Long: `Reads in the given file and adds the contents to your postfile.

Optionally provide a title with --title (-t) and author with --author.

You can configure a default author in your ~/.postfile.yaml config`,
	Run: func(cmd *cobra.Command, args []string) {
		addEntry(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Title, "title", "t", "", "Title of the entry")
}
