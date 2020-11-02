package cmd

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

func getRemoteFileRef(alloc *sdk.Allocation, remotePath string) (
	fr *sdk.ListResult) {

	var dir, _ = filepath.Split(remotePath)

	var lr, err = alloc.ListDir(dir)
	if err != nil {
		log.Fatalf("error listing directory %s: %v", dir, err)
	}

	for _, ch := range lr.Children {
		if ch.Path == remotePath {
			return ch
		}
	}

	log.Fatal("file not found on blobbers: ", remotePath)
	return
}

func getRemoteFileAttributes(alloc *sdk.Allocation, remotePath string) (
	attrs fileref.Attributes) {

	return getRemoteFileRef(alloc, remotePath).Attributes
}

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
		if fflags.Changed("localpath") == false {
			PrintError("Error: localpath flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		localpath := cmd.Flag("localpath").Value.String()
		thumbnailpath := cmd.Flag("thumbnailpath").Value.String()
		encrypt, _ := cmd.Flags().GetBool("encrypt")
		commit, _ := cmd.Flags().GetBool("commit")

		// get original file attributes
		var attrs = getRemoteFileAttributes(allocationObj, remotepath)

		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		if len(thumbnailpath) > 0 {
			if encrypt {
				err = allocationObj.EncryptAndUpdateFileWithThumbnail(localpath,
					remotepath, thumbnailpath, attrs, statusBar)
			} else {
				err = allocationObj.UpdateFileWithThumbnail(localpath,
					remotepath, thumbnailpath, attrs, statusBar)
			}

		} else {
			if encrypt {
				err = allocationObj.EncryptAndUpdateFile(localpath, remotepath,
					attrs, statusBar)
			} else {
				err = allocationObj.UpdateFile(localpath, remotepath,
					attrs, statusBar)
			}
		}
		if err != nil {
			PrintError("Update failed.", err)
			os.Exit(1)
		}

		wg.Wait()
		if !statusBar.success {
			os.Exit(1)
		}

		if commit {
			statusBar.wg.Add(1)
			commitMetaTxn(remotepath, "Update", "", "", allocationObj, nil, statusBar)
			statusBar.wg.Wait()
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	updateCmd.PersistentFlags().String("remotepath", "", "Remote path to upload")
	updateCmd.PersistentFlags().String("localpath", "", "Local path of file to upload")
	updateCmd.PersistentFlags().String("thumbnailpath", "", "Local thumbnail path of file to upload")
	updateCmd.Flags().Bool("encrypt", false, "pass this option to encrypt and upload the file")
	updateCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	updateCmd.MarkFlagRequired("allocation")
	updateCmd.MarkFlagRequired("localpath")
	updateCmd.MarkFlagRequired("remotepath")
}
