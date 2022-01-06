package util

import (
	"fmt"
	"os"
)

// GetConfigDir get config directory , default is ~/.zcn/
func GetConfigDir() string {

	configDir := GetHomeDir() + string(os.PathSeparator) + ".zcn"

	if err := os.MkdirAll(configDir, 0744); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return configDir
}

// GetHomeDir Find home directory.
func GetHomeDir() string {
	// Find home directory.
	idr, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return idr
}
