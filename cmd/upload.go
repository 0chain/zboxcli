package cmd

import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	thrown "github.com/0chain/errors"
	"github.com/0chain/gosdk/core/pathutil"
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
		if !(fflags.Changed("multiuploadjson") || (fflags.Changed("remotepath") && fflags.Changed("localpath"))) {
			PrintError("Error: multiuploadjson or remotepath/localpath flag is missing")
			os.Exit(1)
		}

		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}

		var multiuploadJSON string
		if fflags.Changed("multiuploadjson") {
			multiuploadJSON = cmd.Flag("multiuploadjson").Value.String()
		}

		remotePath := cmd.Flag("remotepath").Value.String()
		localPath := cmd.Flag("localpath").Value.String()
		thumbnailPath := cmd.Flag("thumbnailpath").Value.String()
		encrypt, _ := cmd.Flags().GetBool("encrypt")
		webStreaming, _ := cmd.Flags().GetBool("web-streaming")

		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		if strings.HasPrefix(remotePath, "/Encrypted") {
			encrypt = true
		}

		if multiuploadJSON != "" {
			err = multiUpload(allocationObj, localPath, multiuploadJSON, statusBar)
		} else {
			err = singleUpload(allocationObj, localPath, remotePath, thumbnailPath, encrypt, webStreaming, false, uploadChunkNumber, statusBar)
		}
		if err != nil {
			PrintError("Upload failed.", err.Error())
			os.Exit(1)
		}
		wg.Wait()
		if !statusBar.success {
			os.Exit(1)
		}
	},
}

type chunkedUploadArgs struct {
	localPath     string
	remotePath    string
	thumbnailPath string

	encrypt      bool
	webStreaming bool
	chunkNumber  int
	isUpdate     bool
	isRepair     bool
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

	remotePath, fileName, err := fullPathAndFileNameForUpload(args.localPath, args.remotePath)
	if err != nil {
		return err
	}

	fileMeta := sdk.FileMeta{
		Path:       args.localPath,
		ActualSize: fileInfo.Size(),
		MimeType:   mimeType,
		RemoteName: fileName,
		RemotePath: remotePath,
	}

	options := []sdk.ChunkedUploadOption{
		sdk.WithThumbnailFile(args.thumbnailPath),
		sdk.WithEncrypt(args.encrypt),
		sdk.WithStatusCallback(statusBar),
		sdk.WithChunkNumber(args.chunkNumber),
	}

	chunkedUpload, err := sdk.CreateChunkedUpload(util.GetHomeDir(), allocationObj,
		fileMeta, fileReader,
		args.isUpdate, args.isRepair, args.webStreaming,
		zboxutil.NewConnectionId(),
		options...)

	if err != nil {
		return err
	}

	return chunkedUpload.Start()
}

type MultiUploadOption struct {
	FilePath      string `json:"filePath,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	RemotePath    string `json:"remotePath,omitempty"`
	ThumbnailPath string `json:"thumbnailPath,omitempty"`
	Encrypt       bool   `json:"encrypt,omitempty"`
	ChunkNumber   int    `json:"chunkNumber,omitempty"`
	IsUpdate      bool   `json:"isUpdate,omitempty"`
}

func multiUpload(allocationObj *sdk.Allocation, workdir, jsonMultiUploadOptions string, statusBar *StatusBar) error {
	file, err := os.Open(jsonMultiUploadOptions)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var options []MultiUploadOption

	err = decoder.Decode(&options)
	if err != nil {
		return err
	}

	return multiUploadWithOptions(allocationObj, workdir, options, statusBar)
}

func singleUpload(allocationObj *sdk.Allocation, localPath, remotePath, thumbnailPath string, encrypt, webStreaming, isUpdate bool, chunkNumber int, statusBar *StatusBar) error {
	fullRemotePath, fileName, err := fullPathAndFileNameForUpload(localPath, remotePath)
	if err != nil {
		return err
	}
	remotePath = pathutil.Dir(fullRemotePath) + "/"
	options := []MultiUploadOption{
		{
			FilePath:      localPath,
			FileName:      fileName,
			RemotePath:    remotePath,
			ThumbnailPath: thumbnailPath,
			Encrypt:       encrypt,
			ChunkNumber:   chunkNumber,
			IsUpdate:      isUpdate,
		},
	}

	workdir := util.GetHomeDir()

	return multiUploadWithOptions(allocationObj, workdir, options, statusBar)
}

func multiUploadWithOptions(allocationObj *sdk.Allocation, workdir string, options []MultiUploadOption, statusBar *StatusBar) error {
	totalUploads := len(options)
	filePaths := make([]string, totalUploads)
	fileNames := make([]string, totalUploads)
	remotePaths := make([]string, totalUploads)
	thumbnailPaths := make([]string, totalUploads)
	chunkNumbers := make([]int, totalUploads)
	encrypts := make([]bool, totalUploads)
	isUpdates := make([]bool, totalUploads)
	for idx, option := range options {
		statusBar.wg.Add(1)
		filePaths[idx] = option.FilePath
		fileNames[idx] = option.FileName
		thumbnailPaths[idx] = option.ThumbnailPath
		remotePaths[idx] = option.RemotePath
		chunkNumbers[idx] = option.ChunkNumber
		encrypts[idx] = option.Encrypt

	}

	return allocationObj.StartMultiUpload(workdir, filePaths, fileNames, thumbnailPaths, encrypts, chunkNumbers, remotePaths, isUpdates[0], statusBar)
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	uploadCmd.PersistentFlags().String("remotepath", "", "Remote path to upload")
	uploadCmd.PersistentFlags().String("localpath", "", "Local path of file to upload")
	uploadCmd.PersistentFlags().String("thumbnailpath", "", "Local thumbnail path of file to upload")
	uploadCmd.PersistentFlags().String("multiuploadjson", "", "A JSON file containing multiupload options")
	uploadCmd.PersistentFlags().String("attr-who-pays-for-reads", "owner", "Who pays for reads: owner or 3rd_party")
	uploadCmd.Flags().Bool("encrypt", false, "(default false) pass this option to encrypt and upload the file")
	uploadCmd.Flags().Bool("web-streaming", false, "(default false) pass this option to enable web streaming support")
	uploadCmd.Flags().IntVarP(&uploadChunkNumber, "chunknumber", "", 1, "how many chunks should be uploaded in a http request")

	uploadCmd.MarkFlagRequired("allocation")
	uploadCmd.MarkFlagRequired("remotepath")
	uploadCmd.MarkFlagRequired("localpath")

}

func fullPathAndFileNameForUpload(localPath, remotePath string) (string, string, error) {
	isUploadToDir := strings.HasSuffix(remotePath, "/")
	remotePath = zboxutil.RemoteClean(remotePath)
	if !zboxutil.IsRemoteAbs(remotePath) {
		return "", "", thrown.New("invalid_path", "Path should be valid and absolute")
	}

	// re-add trailing slash to indicate intending to upload to directory
	if isUploadToDir && !strings.HasSuffix(remotePath, "/") {
		remotePath += "/"
	}

	fullRemotePath := zboxutil.GetFullRemotePath(localPath, remotePath)
	_, fileName := pathutil.Split(fullRemotePath)

	return fullRemotePath, fileName, nil
}
