package cmd

import (
	"context"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var (
	mkdirAllocationID string
	mkdirDirName      string
)

func init() {
	rootCmd.AddCommand(createDirCmd)
	createDirCmd.Flags().String("allocation", "", "Allocation ID")
	createDirCmd.Flags().String("dirname", "", "New directory name")
	createDirCmd.MarkFlagRequired("allocation")
	createDirCmd.MarkFlagRequired("dirname")

	rootCmd.AddCommand(dirNewCmd)

	dirNewCmd.Flags().StringVarP(&mkdirAllocationID, "alloc", "a", "", "allocation id")
	dirNewCmd.Flags().StringVarP(&mkdirDirName, "dir", "d", "", "directory name")
	dirNewCmd.MarkFlagRequired("alloc") //nolint
	dirNewCmd.MarkFlagRequired("dir")   //nolint
}

// use dir-new instead
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
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}
		dirname := cmd.Flag("dirname").Value.String()

		if err != nil {
			PrintError("CreateDir failed.", err)
			os.Exit(1)
		}
		err = allocationObj.CreateDir(dirname)

		if err != nil {
			PrintError("CreateDir failed.", err)
			os.Exit(1)
		}

		return
	},
}

var dirNewCmd = &cobra.Command{
	Use:   "dir-new",
	Short: "Create directories named on blobbers",
	Long:  `Create directories named on blobbers`,
	Run: func(cmd *cobra.Command, args []string) {

		err := sdk.CreateDirectory(context.TODO(), mkdirAllocationID, mkdirDirName)
		if err != nil {
			PrintError(err.Error())
			return
		}

		PrintInfof("OK: created directory '%s'", mkdirDirName)
		return
	},
}
