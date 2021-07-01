package util

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

// GetConfigDir get config directory , default is ~/.zcn/
func GetConfigDir() string {

	var configDir string
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configDir = home + string(os.PathSeparator) + ".zcn"

	os.MkdirAll(configDir, 0744)

	return configDir
}
