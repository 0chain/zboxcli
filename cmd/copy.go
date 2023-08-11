package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/constants"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// copyCmd represents copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy an object(file/folder) to another folder on blobbers",
	Long:  `copy an object to another folder on blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}

		if !fflags.Changed("remotepath") {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}

		if !fflags.Changed("destpath") {
			PrintError("Error: destpath flag is missing")
			os.Exit(1)
		}

		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotePath := cmd.Flag("remotepath").Value.String()
		destPath := cmd.Flag("destpath").Value.String()

		err = allocationObj.DoMultiOperation([]sdk.OperationRequest{
			{
				OperationType: constants.FileOperationCopy,
				RemotePath:    remotePath,
				DestPath:      destPath,
			},
		})
		if err != nil {
			PrintError("Error performing CopyObject", err)
			os.Exit(1)
		}

		fmt.Println(remotePath + " copied")
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	copyCmd.PersistentFlags().String("remotepath", "", "Remote path of object to copy")
	copyCmd.PersistentFlags().String("destpath", "", "Destination path for the object. Existing directory the object should be copied to")

	copyCmd.MarkFlagRequired("allocation")
	copyCmd.MarkFlagRequired("remotepath")
	copyCmd.MarkFlagRequired("destpath")
}
