package cmd

import (
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// updateCmd represents update file command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update file to blobbers",
	Long:  `update file to blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()
		if fflags.Changed("allocation") == false {
			PrintError("Error: allocation flag is missing")
			os.Exit(1)
		}
		if fflags.Changed("remotepath") == false {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}
		if !fflags.Changed("localpath") {
			PrintError("Error: localpath flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotePath := cmd.Flag("remotepath").Value.String()

		if remotePath == "/Encrypted" {
			PrintError("Error: can not update Encrypted Folder")
			os.Exit(1)
		}

		localPath := cmd.Flag("localpath").Value.String()
		thumbnailPath := cmd.Flag("thumbnailpath").Value.String()
		encrypt, _ := cmd.Flags().GetBool("encrypt")

		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)

		err = startChunkedUpload(cmd, allocationObj, chunkedUploadArgs{
			localPath:     localPath,
			remotePath:    remotePath,
			thumbnailPath: thumbnailPath,
			encrypt:       encrypt,
			chunkNumber:   updateChunkNumber,
			isUpdate:      true,
			// isRepair:      false,
		}, statusBar)

		if err != nil {
			PrintError("Update failed.", err)
			os.Exit(1)
		}

		wg.Wait()
		if !statusBar.success {
			os.Exit(1)
		}
	},
}

var updateChunkNumber int

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	updateCmd.PersistentFlags().String("remotepath", "", "Remote path to upload")
	updateCmd.PersistentFlags().String("localpath", "", "Local path of file to upload")
	updateCmd.PersistentFlags().String("thumbnailpath", "", "Local thumbnail path of file to upload")
	updateCmd.Flags().Bool("encrypt (boolean)", false, "pass this option to encrypt and upload the file")

	updateCmd.Flags().IntVarP(&updateChunkNumber, "chunknumber", "", 1, "how many chunks should be uploaded in a http request")

	updateCmd.MarkFlagRequired("allocation")
	updateCmd.MarkFlagRequired("localpath")
	updateCmd.MarkFlagRequired("remotepath")
}
