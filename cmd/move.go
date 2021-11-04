package cmd

import (
	"fmt"
	"os"
	"sync"

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
		remotepath := cmd.Flag("remotepath").Value.String()
		destpath := cmd.Flag("destpath").Value.String()
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

		err = allocationObj.MoveObject(remotepath, destpath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(remotepath + " moved")
		if commit {
			fmt.Println("Commiting changes to blockchain ...")
			if isFile {
				wg := &sync.WaitGroup{}
				statusBar := &StatusBar{wg: wg}
				wg.Add(1)
				commitMetaTxn(remotepath, "Move", "", "", allocationObj, fileMeta, statusBar)
				wg.Wait()
			} else {
				commitFolderTxn("Move", remotepath, destpath, allocationObj)
			}
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	moveCmd.PersistentFlags().String("remotepath", "", "Remote path of object to move")
	moveCmd.PersistentFlags().String("destpath", "", "Destination path for the object. Existing directory the object should be copied to")
	moveCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	moveCmd.MarkFlagRequired("allocation")
	moveCmd.MarkFlagRequired("remotepath")
	moveCmd.MarkFlagRequired("destpath")
}
