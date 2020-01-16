package cmd

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// uploadCmd represents upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file to blobbers",
	Long:  `upload file to blobbers`,
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
		thumbnailpath := cmd.Flag("thumbnailpath").Value.String()
		encrypt, _ := cmd.Flags().GetBool("encrypt")
		commit, _ := cmd.Flags().GetBool("commit")
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		if len(thumbnailpath) > 0 {
			if encrypt {
				err = allocationObj.EncryptAndUploadFileWithThumbnail(localpath, remotepath, thumbnailpath, statusBar)
			} else {
				err = allocationObj.UploadFileWithThumbnail(localpath, remotepath, thumbnailpath, statusBar)
			}
		} else {
			if encrypt {
				err = allocationObj.EncryptAndUploadFile(localpath, remotepath, statusBar)
			} else {
				err = allocationObj.UploadFile(localpath, remotepath, statusBar)
			}
		}
		if err != nil {
			PrintError("Upload failed.", err)
			os.Exit(1)
		}
		wg.Wait()
		if !statusBar.success {
			os.Exit(1)
		}

		if commit {
			// Because while upload user gives only the directory in remotepath
			path := remotepath + filepath.Base(localpath)
			commitMetaTxn(path, "Upload", allocationObj)
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	uploadCmd.PersistentFlags().String("remotepath", "", "Remote path to upload")
	uploadCmd.PersistentFlags().String("localpath", "", "Local path of file to upload")
	uploadCmd.PersistentFlags().String("thumbnailpath", "", "Local thumbnail path of file to upload")
	uploadCmd.Flags().Bool("encrypt", false, "pass this option to encrypt and upload the file")
	uploadCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	uploadCmd.MarkFlagRequired("allocation")
	uploadCmd.MarkFlagRequired("localpath")
	uploadCmd.MarkFlagRequired("remotepath")
}
