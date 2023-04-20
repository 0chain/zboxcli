package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/0chain/zboxcli/util"
	"github.com/icza/bitio"
	"github.com/spf13/cobra"
)

var VersionStr string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version information",
	Long:  `Prints version information`,
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")

		if doJSON {
			j := make(map[string]string)
			j["zbox"] = VersionStr
			j["gosdk"] = getVersion("github.com/0chain/gosdk")
			util.PrintJSON(j)
			return
		}

		fmt.Println("Version info:")
		fmt.Println("\tzbox....: ", VersionStr)
		fmt.Println("\tgosdk...: ", getVersion("github.com/0chain/gosdk"))
	},
}

func getVersion(path string) string {
	_ = bitio.NewReader
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to read build info")
		return ""
	}

	for _, dep := range bi.Deps {
		if dep.Path == path {
			if dep.Replace != nil && dep.Replace.Version != "" {
				return dep.Replace.Version
			}

			return dep.Version
		}
	}

	return ""
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().Bool("json (boolean)", false, "pass this option to print response as json data")
}
