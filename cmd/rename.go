package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/constants"
	"github.com/0chain/gosdk/core/pathutil"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// renameCmd represents rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename an object(file/folder) on blobbers",
	Long:  `rename an object on blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		if !fflags.Changed("remotepath") {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}

		if !fflags.Changed("destname") {
			PrintError("Error: destname flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotePath := cmd.Flag("remotepath").Value.String()
		destName := cmd.Flag("destname").Value.String()
		oldName := pathutil.Dir(remotePath)
		if oldName == destName {
			fmt.Println(remotePath + " renamed")
			return
		}

		err = allocationObj.DoMultiOperation([]sdk.OperationRequest{
			{
				OperationType: constants.FileOperationRename,
				RemotePath:    remotePath,
				DestName:      destName,
			},
		})
		if err != nil {
			PrintError("Error performing RenameObject", err)
			os.Exit(1)
		}
		fmt.Println(remotePath + " renamed")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	renameCmd.PersistentFlags().String("remotepath", "", "Remote path of object to rename")
	renameCmd.PersistentFlags().String("destname", "", "New Name for the object (Only the name and not the path). Include the file extension if applicable")

	renameCmd.MarkFlagRequired("allocation")
	renameCmd.MarkFlagRequired("remotepath")
	renameCmd.MarkFlagRequired("destname")
}
