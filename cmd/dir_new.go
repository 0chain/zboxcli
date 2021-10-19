package cmd

import (
	"context"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var (
	dirNewAllocationID string
	dirNewName         string
)

func init() {
	rootCmd.AddCommand(createDirCmd)
	createDirCmd.Flags().String("allocation", "", "Allocation ID")
	createDirCmd.Flags().String("dirname", "", "New directory name")
	createDirCmd.MarkFlagRequired("allocation")
	createDirCmd.MarkFlagRequired("dirname")

	rootCmd.AddCommand(dirNewCmd)
	dirNewCmd.Flags().StringVarP(&dirNewAllocationID, "alloc", "a", "", "allocation id")
	dirNewCmd.Flags().StringVarP(&dirNewName, "name", "n", "", "directory name")
	dirNewCmd.MarkFlagRequired("alloc") //nolint
	dirNewCmd.MarkFlagRequired("name")  //nolint
}

// use dir-create instead
var createDirCmd = &cobra.Command{
	Use:        "createdir",
	Deprecated: "please use mkdir",
	Short:      "Create directory in allocation",
	Long:       `Create directory in allocation`,
	Args:       cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}
		if !fflags.Changed("dirname") {
			PrintError("Error: dirname flag is missing")
			os.Exit(1)
		}

		allocationID := cmd.Flag("allocation").Value.String()
		dirname := cmd.Flag("dirname").Value.String()

		err := sdk.CreateDir(context.TODO(), allocationID, dirname)

		if err != nil {
			PrintError("ERR: ", err.Error())
			return
		}

		PrintInfof("OK: created directory '%s'", dirNewName)
	},
}

var dirNewCmd = &cobra.Command{
	Use:   "dir-new",
	Short: "Create directories named on blobbers",
	Long:  `Create directories named on blobbers`,
	Run: func(cmd *cobra.Command, args []string) {
		err := sdk.CreateDir(context.TODO(), dirNewAllocationID, dirNewName)

		if err != nil {
			PrintError("ERR: ", err.Error())
			return
		}

		PrintInfof("OK: created directory '%s'", dirNewName)
	},
}
