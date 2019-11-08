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
		fmt.Println("Version info:")
		fmt.Println("\tzboxcli..: ", VersionStr)
		fmt.Println("\tgosdk....: ", zcncore.GetVersion())
		return
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
