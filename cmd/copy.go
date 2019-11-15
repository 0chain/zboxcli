package cmd

import (
	"fmt"

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
		remotepath := cmd.Flag("remotepath").Value.String()
		destpath := cmd.Flag("destpath").Value.String()
		err = allocationObj.CopyObject(remotepath, destpath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(remotepath + " copied")
		return
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	copyCmd.PersistentFlags().String("remotepath", "", "Remote path of object to copy")
	copyCmd.PersistentFlags().String("destpath", "", "Destination path for the object")
	copyCmd.MarkFlagRequired("allocation")
	copyCmd.MarkFlagRequired("remotepath")
	copyCmd.MarkFlagRequired("destpath")
}
