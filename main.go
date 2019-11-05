package main

import (
	"github.com/0chain/zboxcli/cmd"
)

var VersionStr string

func main() {
	cmd.VersionStr = VersionStr
	cmd.Execute()
}
