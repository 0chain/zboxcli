package cmd

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zboxcore/zboxutil"

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
		commit, _ := cmd.Flags().GetBool("commit")
		pre_at_blobber, _ := cmd.Flags().GetBool("pre-at-blobber")
		encrypt, _ := cmd.Flags().GetBool("encrypt")
		// when pre_at_blobber is true, encrypt must also be true
		encrypt = encrypt || pre_at_blobber
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		if strings.HasPrefix(remotepath, "/Encrypted") {
			encrypt = true
		}
		var attrs fileref.Attributes
		if fflags.Changed("attr-who-pays-for-reads") {
			var (
				wp  common.WhoPays
				wps string
			)
			if wps, err = fflags.GetString("attr-who-pays-for-reads"); err != nil {
				log.Fatalf("getting 'attr-who-pays-for-reads' flag: %v", err)
			}
			if err = wp.Parse(wps); err != nil {
				log.Fatal(err)
			}
			attrs.WhoPaysForReads = wp // set given value
		}
		attrs.PreAtBlobber = pre_at_blobber

		if len(thumbnailpath) > 0 {
			if encrypt {
				err = allocationObj.EncryptAndUploadFileWithThumbnail(localpath, remotepath, thumbnailpath, attrs, statusBar)
			} else {
				err = allocationObj.UploadFileWithThumbnail(localpath, remotepath, thumbnailpath, attrs, statusBar)
			}
		} else {
			if encrypt {
				err = allocationObj.EncryptAndUploadFile(localpath, remotepath, attrs, statusBar)
			} else {
				err = allocationObj.UploadFile(localpath, remotepath, attrs, statusBar)
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
			remotepath = zboxutil.GetFullRemotePath(localpath, remotepath)
			statusBar.wg.Add(1)
			commitMetaTxn(remotepath, "Upload", "", "", allocationObj, nil, statusBar)
			statusBar.wg.Wait()
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
	uploadCmd.PersistentFlags().String("attr-who-pays-for-reads", "owner", "Who pays for reads: owner or 3rd_party")
	uploadCmd.Flags().Bool("encrypt", false, "pass this option to encrypt and upload the file")
	uploadCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	uploadCmd.Flags().Bool("pre-at-blobber", false, "pass this option to use pre key at blobber")
	uploadCmd.MarkFlagRequired("allocation")
	uploadCmd.MarkFlagRequired("localpath")
	uploadCmd.MarkFlagRequired("remotepath")
}
