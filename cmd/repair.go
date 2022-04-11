package cmd

import (
	"os"
	"sync"

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
		if fflags.Changed("rootpath") == false {
			PrintError("Error: rootpath flag is missing")
			os.Exit(1)
		}
		if fflags.Changed("repairpath") == false {
			PrintError("Error: repairpath flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := storageSdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}
		localRootPath := cmd.Flag("rootpath").Value.String()
		repairPath := cmd.Flag("repairpath").Value.String()

		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		allocUnderRepair = true
		err = allocationObj.StartRepair(localRootPath, repairPath, statusBar)
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
	startRepair.PersistentFlags().String("rootpath", "", "File path for local files ")
	startRepair.PersistentFlags().String("repairpath", "", "Path to repair")
	startRepair.MarkFlagRequired("allocation")
	startRepair.MarkFlagRequired("rootpath")
	startRepair.MarkFlagRequired("repairpath")
}
