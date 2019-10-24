package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete file from blobbers",
	Long:  `delete file from blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}
		if fflags.Changed("remotepath") == false {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		err = allocationObj.DeleteFile(remotepath)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}
		fmt.Println(remotepath + " deleted")
		return
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	deleteCmd.PersistentFlags().String("remotepath", "", "Remote path of file to delete")
	deleteCmd.MarkFlagRequired("allocation")
	deleteCmd.MarkFlagRequired("remotepath")
}
