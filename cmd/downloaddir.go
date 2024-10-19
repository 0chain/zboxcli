package cmd

import (
	"context"
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// downloadCmd represents download directory command
var downdirCmd = &cobra.Command{
	Use:   "downloaddir",
	Short: "download directory from blobbers",
	Long:  `download directory from blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags() // fflags is a *flag.FlagSet
		if !(fflags.Changed("remotepath") || fflags.Changed("authticket")) {
			PrintError("Error: remotepath / authticket flag is missing")
			os.Exit(1)
		}

		remotePath := cmd.Flag("remotepath").Value.String()
		authTicket := cmd.Flag("authticket").Value.String()
		localPath := cmd.Flag("localpath").Value.String()
		allocationID := cmd.Flag("allocation").Value.String()

		var allocationObj *sdk.Allocation
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		errE := allocationObj.DownloadDirectory(context.Background(), remotePath, localPath, authTicket, statusBar)
		if errE == nil {
			wg.Wait()
		} else {
			PrintError("Download failed.", errE.Error())
			os.Exit(1)
		}
		if !statusBar.success {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(downdirCmd)
	downdirCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	downdirCmd.PersistentFlags().String("remotepath", "", "Remote path of directory to download")
	downdirCmd.PersistentFlags().String("localpath", "", "Local path of directory to download")
	downdirCmd.PersistentFlags().String("authticket", "", "Auth ticket fot the file to download if you dont own it")
	downdirCmd.MarkFlagRequired("allocation")
	downdirCmd.MarkFlagRequired("localpath")
	downdirCmd.MarkFlagRequired("remotepath")
}
