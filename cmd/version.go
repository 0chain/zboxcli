package cmd

import (
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

		PrintInfo("Version info:")
		PrintInfo("\tzbox....: ", VersionStr)
		PrintInfo("\tgosdk...: ", getVersion("github.com/0chain/gosdk"))
	},
}

func getVersion(path string) string {
	_ = bitio.NewReader
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		PrintInfo("Failed to read build info")
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
	versionCmd.Flags().Bool("json", false, "(default false) pass this option to print response as json data")
}
