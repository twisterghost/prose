package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var prosefileLocation string
var author string
var pretty bool

var rootCmd = &cobra.Command{
	Use:   "prose [command]",
	Short: "A simple way to manage your writing.",
	Long: `Write, store and manage written word.
Pipe through modifiers and renderers to fit your need.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prose.yaml)")
	rootCmd.PersistentFlags().StringVar(&prosefileLocation, "file", "", "prosefile location (default is $HOME/prosefile.json)")
	rootCmd.PersistentFlags().StringVar(&author, "author", "", "the name of the post author (default is blank)")
	rootCmd.PersistentFlags().BoolVar(&pretty, "pretty", false, "Format all json as human readable")
	viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("pretty", rootCmd.PersistentFlags().Lookup("pretty"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".prosefile" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".prose")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadInConfig()
}
