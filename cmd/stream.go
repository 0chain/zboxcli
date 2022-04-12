package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	thrown "github.com/0chain/errors"
	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zboxcore/zboxutil"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

// streamCmd represents upload command with --live flag
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "capture video and audio streaming form local devices, and upload",
	Long:  "capture video and audio streaming form local devices, and upload",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}
		if !fflags.Changed("remotepath") {
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
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		localpath := cmd.Flag("localpath").Value.String()
		encrypt, _ := cmd.Flags().GetBool("encrypt")
		commit, _ := cmd.Flags().GetBool("commit")

		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		if strings.HasPrefix(remotepath, "/Encrypted") {
			encrypt = true
		}
		var attrs fileref.Attributes
		if fflags.Changed("attr-who-pays-for-reads") {
			var wp common.WhoPays
			var wps string
			if wps, err = fflags.GetString("attr-who-pays-for-reads"); err != nil {
				log.Fatalf("getting 'attr-who-pays-for-reads' flag: %v", err)
			}
			if err = wp.Parse(wps); err != nil {
				log.Fatal(err)
			}
			attrs.WhoPaysForReads = wp // set given value
		}

		// capture video and audio from local default camera and micrlphone, and upload it to zcn
		err = startLiveUpload(cmd, allocationObj, localpath, remotepath, encrypt, streamChunkNumber, attrs)

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
	},
}

func startLiveUpload(cmd *cobra.Command, allocationObj *sdk.Allocation, localPath string, remotePath string, encrypt bool, chunkNumber int, attrs fileref.Attributes) error {

	delay, _ := cmd.Flags().GetInt("delay")

	reader, err := sdk.CreateFfmpegRecorder(localPath, delay)
	if err != nil {
		return err
	}

	defer reader.Close()

	mimeType, err := reader.GetFileContentType()
	if err != nil {
		return err
	}

	remotePath = zboxutil.RemoteClean(remotePath)
	isabs := zboxutil.IsRemoteAbs(remotePath)
	if !isabs {
		err = thrown.New("invalid_path", "Path should be valid and absolute")
		return err
	}
	remotePath = zboxutil.GetFullRemotePath(localPath, remotePath)

	_, fileName := filepath.Split(remotePath)

	liveMeta := sdk.LiveMeta{
		MimeType:   mimeType,
		RemoteName: fileName,
		RemotePath: remotePath,
		Attributes: attrs,
	}

	liveUpload := sdk.CreateLiveUpload(util.GetHomeDir(), allocationObj, liveMeta, reader,
		sdk.WithLiveChunkNumber(chunkNumber),
		sdk.WithLiveEncrypt(encrypt),
		sdk.WithLiveStatusCallback(func() sdk.StatusCallback {
			wg := &sync.WaitGroup{}
			statusBar := &StatusBar{wg: wg}
			wg.Add(1)

			return statusBar
		}),
		sdk.WithLiveDelay(delay))

	return liveUpload.Start()
}

var streamChunkNumber int

func init() {

	// stream Command
	rootCmd.AddCommand(streamCmd)
	streamCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	streamCmd.PersistentFlags().String("remotepath", "", "Remote path to upload")
	streamCmd.PersistentFlags().String("localpath", "", "Local path of file to upload")
	streamCmd.PersistentFlags().String("thumbnailpath", "", "Local thumbnail path of file to upload")
	streamCmd.PersistentFlags().String("attr-who-pays-for-reads", "owner", "Who pays for reads: owner or 3rd_party")
	streamCmd.Flags().Bool("encrypt", false, "pass this option to encrypt and upload the file")
	streamCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")

	streamCmd.Flags().IntVarP(&streamChunkNumber, "chunknumber", "", 1, "how many chunks should be uploaded in a http request")

	streamCmd.Flags().Int("delay", 5, "set segment duration to seconds.")

	streamCmd.MarkFlagRequired("allocation")
	streamCmd.MarkFlagRequired("remotepath")
	streamCmd.MarkFlagRequired("localpath")

}
