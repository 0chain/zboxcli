package cmd

import (
	"encoding/json"
	"fmt"
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
		if fflags.Changed("rootpath") == false {
			PrintError("Error: rootpath flag is missing")
			os.Exit(1)
		}
		if fflags.Changed("repairpath") == false {
			PrintError("Error: repairpath flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
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

var repairSize = &cobra.Command{
	Use:   "repair-size",
	Short: "gets only size to repair file to blobbers",
	Long:  `gets only size to repair file to blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                     
		if !fflags.Changed("allocation") { 
			PrintError("Error: allocation flag is missing")
			os.Exit(1)                                     
		}

		repairPath := "/"
		var err error
		if fflags.Changed("repairpath") {
			if repairPath, err = fflags.GetString("repairpath"); err != nil {
				PrintError("Error: repairpath is not of string type", err)
				os.Exit(1)
			}
		}

		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}

		size, err := allocationObj.RepairSize(repairPath)
		if err != nil {
			PrintError("get repair size failed: ", err)
			os.Exit(1)
		}

		jsonBytes, err := json.Marshal(size)
		if err != nil {
			PrintError("error marshaling size: ", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonBytes))
	},
}

func init() {
	rootCmd.AddCommand(startRepair)
	rootCmd.AddCommand(repairSize)
	startRepair.PersistentFlags().String("allocation", "", "Allocation ID")
	startRepair.PersistentFlags().String("rootpath", "", "File path for local files ")
	startRepair.PersistentFlags().String("repairpath", "", "Path to repair")
	startRepair.MarkFlagRequired("allocation")
	startRepair.MarkFlagRequired("rootpath")
	startRepair.MarkFlagRequired("repairpath")

	repairSize.PersistentFlags().String("allocation", "", "Allocation ID")
	repairSize.PersistentFlags().String("repairpath", "", "Path to repair")
	repairSize.MarkFlagRequired("allocation")
}
