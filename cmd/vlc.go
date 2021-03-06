package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{

	Use:   "pack",
	Short: "Pack file using variable-lenght code",
	Run:   pack,
}

const packedExtension = "vlc"

func pack(_ *cobra.Command, args []string) {
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	// packed := Encode(data)
	packed := "" + string(data) //TODO remove

	err = ioutil.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	// path to file myfile.txt -> myfile.vlc
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
