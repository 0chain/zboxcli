package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// moveCmd represents move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "move an object(file/folder) to another folder on blobbers",
	Long:  `move an object to another folder on blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)
		}
		if !fflags.Changed("remotepath") {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}

		if !fflags.Changed("destpath") {
			fmt.Println("Error: destpath flag is missing")
			return
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotePath := cmd.Flag("remotepath").Value.String()
		destPath := cmd.Flag("destpath").Value.String()

		err = allocationObj.MoveObject(remotePath, destPath)
		if err != nil {
			PrintError("Error performing CopyObject", err)
			os.Exit(1)
		}

		fmt.Println(remotePath + " moved")

	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	moveCmd.PersistentFlags().String("remotepath", "", "Remote path of object to move")
	moveCmd.PersistentFlags().String("destpath", "", "Destination path for the object. Existing directory the object should be copied to")

	moveCmd.MarkFlagRequired("allocation")
	moveCmd.MarkFlagRequired("remotepath")
	moveCmd.MarkFlagRequired("destpath")
}
