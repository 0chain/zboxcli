package cmd

import (
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// uploadCmd represents upload command
var repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "repair file to blobbers",
	Long:  `repair file to blobbers`,
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
		if fflags.Changed("localpath") == false {
			PrintError("Error: localpath flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		localpath := cmd.Flag("localpath").Value.String()
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		err = allocationObj.RepairFile(localpath, remotepath, statusBar)
		if err != nil {
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
	rootCmd.AddCommand(repairCmd)
	repairCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	repairCmd.PersistentFlags().String("remotepath", "", "Remote path to repair")
	repairCmd.PersistentFlags().String("localpath", "", "Local path of file to repair")
	repairCmd.MarkFlagRequired("allocation")
	repairCmd.MarkFlagRequired("localpath")
	repairCmd.MarkFlagRequired("remotepath")
}
