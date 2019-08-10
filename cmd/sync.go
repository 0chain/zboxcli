package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func printTable(files []sdk.FileDiff) {
	header := []string{"Operation", "Path"}
	data := make([][]string, len(files))
	for idx, child := range files {
		data[idx] = []string{child.Op, child.Path}
	}
	util.WriteTable(os.Stdout, header, []string{}, data)
}

// syncCmd represents sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync files from blobbers",
	Long:  `Sync all files from blobbers to a localpath`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags() // fflags is a *flag.FlagSet
		if fflags.Changed("localpath") == false {
			fmt.Println("Error: localpath flag is missing")
			return
		}

		localpath := cmd.Flag("localpath").Value.String()

		if len(localpath) == 0 {
			fmt.Println("Error: localpath flag is missing")
			return
		}

		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			fmt.Println("Error: allocation flag is missing") // If not, we'll let the user know
			return                                           // and return
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			fmt.Println("Error fetching the allocation", err)
			return
		}

		// Create filter
		filter := []string{".DS_Store", ".git"}
		lDiff, err := allocationObj.GetAllocationDiff("", localpath, filter)
		if err != nil {
			fmt.Println("Error gettind diff.", err)
			return
		}
		if len(lDiff) > 0 {
			printTable(lDiff)
		} else {
			fmt.Println("Already up to date")
		}
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		for _, f := range lDiff {
			lPath := localpath + "/" + f.Path
			switch f.Op {
			case "Download":
				wg.Add(1)
				err = allocationObj.DownloadFile(lPath, f.Path, statusBar)
			case "Upload":
				wg.Add(1)
				err = allocationObj.UploadFile(lPath, f.Path, statusBar)
			case "Update":
				wg.Add(1)
				err = allocationObj.UpdateFile(lPath, f.Path, statusBar)
			}
			if err == nil {
				wg.Wait()
			} else {
				fmt.Println(err.Error())
			}
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	syncCmd.PersistentFlags().String("localpath", "", "Local dir path to sync")
	syncCmd.MarkFlagRequired("allocation")
	syncCmd.MarkFlagRequired("localpath")
}
