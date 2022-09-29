package cmd

import (
	"fmt"

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
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			fmt.Println("Error: allocation flag is missing") // If not, we'll let the user know
			return                                           // and return
		}
		if fflags.Changed("remotepath") == false {
			fmt.Println("Error: remotepath flag is missing")
			return
		}

		if fflags.Changed("destpath") == false {
			fmt.Println("Error: destpath flag is missing")
			return
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			fmt.Println("Error fetching the allocation", err)
			return
		}
		remotePath := cmd.Flag("remotepath").Value.String()
		destPath := cmd.Flag("destpath").Value.String()

		err = allocationObj.MoveObject(remotePath, destPath)
		if err != nil {
			fmt.Println(err.Error())
			return
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
