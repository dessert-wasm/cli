package main

import (
	"dessert-cli/cmd"
	"os"

	"github.com/spf13/afero"
)

func main() {
	cmd.Fs = afero.NewOsFs()
	root := cmd.CreateRoot()
	_, err := root.ExecuteC()

	if err != nil {
		cmd.Logger.Errorln(err)
		os.Exit(1)
	}
}
