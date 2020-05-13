package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/twisterghost/prose/loader"
	"github.com/twisterghost/prose/prose"
)

func deletePost(delId string) {
	prosefile := loader.LoadProsefile()
	prosefile.Entries = prose.RemoveEntryById(prosefile.Entries, delId)

	loader.WriteProsefile(loader.SerializeProsefile(prosefile))

	fmt.Println("Post deleted:", delId)
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a prosefile entry",
	Long:  `Delete a prosefile entry of the given ID`,
	Run: func(cmd *cobra.Command, args []string) {
		deletePost(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
