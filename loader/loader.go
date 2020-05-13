package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/twisterghost/prose/prose"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func GetProsefilePath() string {
	prosefilePath := viper.GetString("file")

	// If no path was given, look for prosefile.json in the home directory
	if prosefilePath == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error finding home directory")
			fmt.Print(err)
			os.Exit(1)
		}
		prosefilePath = filepath.Join(home, "prosefile.json")
	}

	return prosefilePath
}

func LoadProsefile() prose.Prosefile {
	prosefilePath := GetProsefilePath()
	data, err := ioutil.ReadFile(prosefilePath)
	if err != nil {
		fmt.Println("Error when reading prosefile at", prosefilePath)
		fmt.Println(err)
		os.Exit(1)
	}

	var obj prose.Prosefile
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("Error when parsing prosefile json.")
		fmt.Println("Reading from file:", prosefilePath)
		fmt.Println(err)
		os.Exit(1)
	}

	return obj
}

// TODO: Report errors instead of printing them here
func WriteProsefile(prosefile prose.Prosefile) {
	pretty := viper.GetBool("pretty")
	serialized := prosefile.Serialize(pretty)
	bytes := []byte(serialized)
	err := ioutil.WriteFile(GetProsefilePath(), bytes, 0644)
	if err != nil {
		fmt.Println("Error writing prosefile at", GetProsefilePath())
		fmt.Println(err)
	}
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
