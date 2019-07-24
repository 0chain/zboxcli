package cmd

import (
	"fmt"

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
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			fmt.Println("Error: allocation flag is missing") // If not, we'll let the user know
			return                                           // and return
		}
		if fflags.Changed("remotepath") == false {
			fmt.Println("Error: remotepath flag is missing")
			return
		}

		if fflags.Changed("destname") == false {
			fmt.Println("Error: destname flag is missing")
			return
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			fmt.Println("Error fetching the allocation", err)
			return
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		destname := cmd.Flag("destname").Value.String()
		err = allocationObj.RenameObject(remotepath, destname)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(remotepath + " renamed")
		return
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	renameCmd.PersistentFlags().String("remotepath", "", "Remote path of file to delete")
	renameCmd.PersistentFlags().String("destname", "", "New Name for the object")
	renameCmd.MarkFlagRequired("allocation")
	renameCmd.MarkFlagRequired("remotepath")
	renameCmd.MarkFlagRequired("destname")
}
