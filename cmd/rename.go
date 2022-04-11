package cmd

import (
	"fmt"
	"os"
	"sync"

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
		allocationObj, err := storageSdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		destname := cmd.Flag("destname").Value.String()
		commit, _ := cmd.Flags().GetBool("commit")

		statsMap, err := allocationObj.GetFileStats(remotepath)
		if err != nil {
			PrintError("Error in getting information about the object." + err.Error())
			os.Exit(1)
		}
		isFile := false
		for _, v := range statsMap {
			if v != nil {
				isFile = true
				break
			}
		}

		var fileMeta *sdk.ConsolidatedFileMeta
		if isFile && commit {
			fileMeta, err = allocationObj.GetFileMeta(remotepath)
			if err != nil {
				PrintError("Failed to fetch metadata for the given file", err.Error())
				os.Exit(1)
			}
		}
		err = allocationObj.RenameObject(remotepath, destname)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}
		fmt.Println(remotepath + " renamed")

		if commit {
			fmt.Println("Commiting changes to blockchain ...")
			if isFile {
				wg := &sync.WaitGroup{}
				statusBar := &StatusBar{wg: wg}
				wg.Add(1)
				commitMetaTxn(remotepath, "Rename", "", "", allocationObj, fileMeta, statusBar)
				wg.Wait()
			} else {
				commitFolderTxn("Rename", remotepath, destname, allocationObj)
			}
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	renameCmd.PersistentFlags().String("remotepath", "", "Remote path of object to rename")
	renameCmd.PersistentFlags().String("destname", "", "New Name for the object (Only the name and not the path). Include the file extension if applicable")
	renameCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	renameCmd.MarkFlagRequired("allocation")
	renameCmd.MarkFlagRequired("remotepath")
	renameCmd.MarkFlagRequired("destname")
}
