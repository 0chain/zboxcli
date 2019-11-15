package cmd

import (
	"fmt"
	"os"

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
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		if fflags.Changed("remotepath") == false {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}

		if fflags.Changed("destname") == false {
			PrintError("Error: destname flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		destname := cmd.Flag("destname").Value.String()
		err = allocationObj.RenameObject(remotepath, destname)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}
		fmt.Println(remotepath + " renamed")
		return
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
