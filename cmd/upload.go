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

var uploadChunkNumber int = 1

// uploadCmd represents upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file to blobbers",
	Long:  `upload file to blobbers`,
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
		thumbnailpath := cmd.Flag("thumbnailpath").Value.String()
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

		chunkSize, _ := cmd.Flags().GetInt("chunksize")

		if err := startChunkedUpload(cmd, allocationObj,
			chunkedUploadArgs{
				localPath:     localpath,
				thumbnailPath: thumbnailpath,
				remotePath:    remotepath,
				encrypt:       encrypt,
				chunkSize:     chunkSize,
				chunkNumber:   uploadChunkNumber,
				attrs:         attrs,
				// isUpdate:      false,
				// isRepair:      false,
			}, statusBar); err != nil {
			PrintError("Upload failed.", err.Error())
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

type chunkedUploadArgs struct {
	localPath     string
	remotePath    string
	thumbnailPath string

	encrypt     bool
	chunkSize   int
	chunkNumber int
	isUpdate    bool
	isRepair    bool
	attrs       fileref.Attributes
}

func startChunkedUpload(cmd *cobra.Command, allocationObj *sdk.Allocation, args chunkedUploadArgs, statusBar sdk.StatusCallback) error {

	fileReader, err := os.Open(args.localPath)
	if err != nil {
		return err
	}
	defer fileReader.Close()

	fileInfo, err := fileReader.Stat()
	if err != nil {
		return err
	}

	mimeType, err := zboxutil.GetFileContentType(fileReader)
	if err != nil {
		return err
	}

	remotePath := zboxutil.RemoteClean(args.remotePath)
	isabs := zboxutil.IsRemoteAbs(remotePath)
	if !isabs {
		err = thrown.New("invalid_path", "Path should be valid and absolute")
		return err
	}
	remotePath = zboxutil.GetFullRemotePath(args.localPath, remotePath)

	_, fileName := filepath.Split(remotePath)

	fileMeta := sdk.FileMeta{
		Path:       args.localPath,
		ActualSize: fileInfo.Size(),
		MimeType:   mimeType,
		RemoteName: fileName,
		RemotePath: remotePath,
		Attributes: args.attrs,
	}

	options := []sdk.ChunkedUploadOption{
		sdk.WithThumbnailFile(args.thumbnailPath),
		sdk.WithChunkSize(int64(args.chunkSize)),
		sdk.WithEncrypt(args.encrypt),
		sdk.WithStatusCallback(statusBar),
		sdk.WithChunkNumber(args.chunkNumber),
	}

	chunkedUpload, err := sdk.CreateChunkedUpload(util.GetHomeDir(), allocationObj,
		fileMeta, fileReader,
		args.isUpdate, args.isRepair,
		options...)

	if err != nil {
		return err
	}

	return chunkedUpload.Start()
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
	uploadCmd.Flags().Int("chunksize", sdk.CHUNK_SIZE, "chunk size")
	uploadCmd.Flags().IntVarP(&uploadChunkNumber, "chunknumber", "", 1, "how many chunks should be uploaded in a http request")

	uploadCmd.MarkFlagRequired("allocation")
	uploadCmd.MarkFlagRequired("remotepath")
	uploadCmd.MarkFlagRequired("localpath")

}
