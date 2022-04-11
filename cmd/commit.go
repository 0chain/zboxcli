package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// copyCmd represents copy command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit a file changes to chain ",
	Long:  `commit a file changes to chain`,
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
		if fflags.Changed("operation") == false {
			fmt.Println("Error: operation flag is missing")
			return
		}

		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := storageSdk.GetAllocation(allocationID)
		if err != nil {
			fmt.Println("Error fetching the allocation", err)
			return
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		operation := cmd.Flag("operation").Value.String()

		newvalue := cmd.Flag("newvalue").Value.String()
		filemeta := cmd.Flag("filemeta").Value.String()
		authticket := cmd.Flag("authticket").Value.String()
		lookuphash := cmd.Flag("lookuphash").Value.String()

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

		var fileMetaData *sdk.ConsolidatedFileMeta
		if len(filemeta) > 0 {
			err := json.Unmarshal([]byte(filemeta), fileMetaData)
			if err != nil {
				PrintError("failed to convert fileMeta." + err.Error())
				os.Exit(1)
			}
		}

		if isFile {
			wg := &sync.WaitGroup{}
			statusBar := &StatusBar{wg: wg}
			wg.Add(1)
			commitMetaTxn(remotepath, operation, authticket, lookuphash, allocationObj, fileMetaData, statusBar)
			wg.Wait()
		} else {
			commitFolderTxn(operation, remotepath, newvalue, allocationObj)
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	commitCmd.PersistentFlags().String("remotepath", "", "Remote path of object to commit")
	commitCmd.PersistentFlags().String("operation", "", "Operation name for the commit changes")
	commitCmd.PersistentFlags().String("newvalue", "", "New value for the folder operation if applicable")
	commitCmd.PersistentFlags().String("filemeta", "", "provide file meta for commit if applicable")
	commitCmd.PersistentFlags().String("authticket", "", "Auth ticket for the file to commit")
	commitCmd.PersistentFlags().String("lookuphash", "", "The remote lookuphash of the object to commit")

	commitCmd.MarkFlagRequired("allocation")
	commitCmd.MarkFlagRequired("remotepath")
	commitCmd.MarkFlagRequired("operation")
}
