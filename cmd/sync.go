package cmd

import (
	"fmt"
	"os"
	"strings"
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
	fmt.Println("")
}

func saveCache(allocationObj *sdk.Allocation, path string, exclPath []string) {
	if len(path) > 0 {
		err := allocationObj.SaveRemoteSnapshot(path, exclPath)
		if err != nil {
			fmt.Println("Failed to save local cache. %v", err)
			return
		}
		fmt.Println("Local cache saved.")
	}
}

// syncCmd represents sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync files to/from blobbers",
	Long:  `Sync all files to/from blobbers from/to a localpath`,
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

		localcache := ""
		if fflags.Changed("localcache") {
			localcache = cmd.Flag("localcache").Value.String()
		}
		exclPath := []string{}
		if fflags.Changed("excludepath") {
			exclPath, _ = cmd.Flags().GetStringArray("excludepath")
		}

		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			fmt.Println("Error fetching the allocation", err)
			return
		}

		// Create filter
		filter := []string{".DS_Store", ".git"}
		lDiff, err := allocationObj.GetAllocationDiff(localcache, localpath, filter, exclPath)
		if err != nil {
			fmt.Println("Error getting diff.", err)
			return
		}
		if len(lDiff) > 0 {
			printTable(lDiff)
		} else {
			fmt.Println("Already up to date")
			saveCache(allocationObj, localcache, exclPath)
			return
		}
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		for _, f := range lDiff {
			localpath = strings.TrimRight(localpath, "/")
			lPath := localpath + f.Path
			switch f.Op {
			case sdk.Download:
				wg.Add(1)
				err = allocationObj.DownloadFile(lPath, f.Path, statusBar)
			case sdk.Upload:
				wg.Add(1)
				err = allocationObj.UploadFile(lPath, f.Path, statusBar)
			case sdk.Update:
				wg.Add(1)
				err = allocationObj.UpdateFile(lPath, f.Path, statusBar)
			case sdk.Delete:
				// TODO: User confirm??
				fmt.Printf("Deleting remote %s...\n", f.Path)
				err = allocationObj.DeleteFile(f.Path)
				if err != nil {
					fmt.Println("Error deleting remote file,", err.Error())
				}
				continue
			case sdk.LocalDelete:
				// TODO: User confirm??
				fmt.Printf("Deleting local %s...\n", lPath)
				err = os.Remove(lPath)
				if err != nil {
					fmt.Println("Error deleting local file.", err.Error())
				}
				continue
			}
			if err == nil {
				wg.Wait()
			} else {
				fmt.Println(err.Error())
			}
		}
		fmt.Println("\nSync Complete")
		saveCache(allocationObj, localcache, exclPath)
		return
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	syncCmd.PersistentFlags().String("localpath", "", "Local dir path to sync")
	syncCmd.PersistentFlags().String("localcache", "", `Local cache of remote snapshot.
If file exists, this will be used for comparison with remote.
After sync complete, remote snapshot will be updated to the same file for next use.`)
	syncCmd.PersistentFlags().StringArray("excludepath", []string{}, "Remote folder paths exclude to sync")
	syncCmd.MarkFlagRequired("allocation")
	syncCmd.MarkFlagRequired("localpath")
}
