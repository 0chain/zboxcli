package cmd

import (
	"fmt"
	"os"
	"sync"

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
		commit, _ := cmd.Flags().GetBool("commit")
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()

		fileMeta, err := allocationObj.GetFileMeta(remotepath)
		if err != nil {
			PrintError("Failed to fetch metadata for the given file", err.Error())
			os.Exit(1)
		}

		err = allocationObj.DeleteFile(remotepath)
		if err != nil {
			PrintError("Delete failed.", err.Error())
			os.Exit(1)
		}
		fmt.Println(remotepath + " deleted")
		if commit {
			wg := &sync.WaitGroup{}
			statusBar := &StatusBar{wg: wg}
			wg.Add(1)
			commitMetaTxn(remotepath, "Delete", "", "", allocationObj, fileMeta, statusBar)
			wg.Wait()
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	deleteCmd.PersistentFlags().String("remotepath", "", "Remote path of the object to delete")
	deleteCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	deleteCmd.MarkFlagRequired("allocation")
	deleteCmd.MarkFlagRequired("remotepath")
}
