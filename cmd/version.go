package cmd

import (
	"fmt"

	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version information",
	Long:  `Prints version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zboxcli version:", VersionStr)
		fmt.Println("gosdk version:  ", zcncore.GetVersion())
		return
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
