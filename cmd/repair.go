package cmd

import (
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// startRepair represents startRepair command
var startRepair = &cobra.Command{
	Use:   "start-repair",
	Short: "start repair file to blobbers",
	Long:  `start repair file to blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}
		if fflags.Changed("fileid") == false {
			PrintError("Error: fileid flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}
		localPath := cmd.Flag("localpath").Value.String()
		fileid, _ := cmd.Flags().GetInt64("fileid")

		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		allocUnderRepair = true
		err = allocationObj.StartRepair(localPath, fileid, statusBar)
		if err != nil {
			allocUnderRepair = false
			PrintError("Repair failed.", err)
			os.Exit(1)
		}
		wg.Wait()
		if !statusBar.success {
			os.Exit(1)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(startRepair)
	startRepair.PersistentFlags().String("allocation", "", "Allocation ID")
	startRepair.PersistentFlags().String("localpath", "", "the localpath of the file to be repaired")
	startRepair.PersistentFlags().Int64("fileid", 0, "file ID to repair")
	startRepair.MarkFlagRequired("allocation")
	startRepair.MarkFlagRequired("fileid")
}
