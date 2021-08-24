package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	thrown "github.com/0chain/errors"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zboxcore/zboxutil"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

var createDirCmd = &cobra.Command{
	Use:   "createdir",
	Short: "Create directory",
	Long:  `Create directory`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}
		if !fflags.Changed("dirname") {
			PrintError("Error: dirname flag is missing")
			os.Exit(1)
		}

		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}
		dirname := cmd.Flag("dirname").Value.String()

		if err != nil {
			PrintError("CreateDir failed.", err)
			os.Exit(1)
		}
		err = allocationObj.CreateDir(dirname)

		if err != nil {
			PrintError("CreateDir failed.", err)
			os.Exit(1)
		}

		return
	},
}

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
		stream, _ := cmd.Flags().GetBool("stream")

		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		if strings.HasPrefix(remotepath, "/Encrypted") {
			encrypt = true
		}

		var attrs fileref.Attributes // depreciated

		live, _ := cmd.Flags().GetBool("live")
		sync, _ := cmd.Flags().GetBool("sync")
		chunkSize, _ := cmd.Flags().GetInt("chunksize")

		if live {
			// capture video and audio from local default camera and micrlphone, and upload it to zcn
			err = startLiveUpload(cmd, allocationObj, localpath, remotepath, encrypt, chunkSize, attrs)
		} else if sync {
			// download video from remote live feed(eg youtube), and sync it to zcn
			err = startSyncUpload(cmd, allocationObj, localpath, remotepath, encrypt, chunkSize, attrs)
		} else if stream {

			err = startChunkedUpload(cmd, allocationObj, localpath, thumbnailpath, remotepath, encrypt, chunkSize, attrs, statusBar)

		} else {

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

func startChunkedUpload(cmd *cobra.Command, allocationObj *sdk.Allocation, localPath, thumbnailPath, remotePath string, encrypt bool, chunkSize int, attrs fileref.Attributes, statusBar sdk.StatusCallback) error {

	fileReader, err := os.Open(localPath)
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

	remotePath = zboxutil.RemoteClean(remotePath)
	isabs := zboxutil.IsRemoteAbs(remotePath)
	if !isabs {
		err = thrown.New("invalid_path", "Path should be valid and absolute")
		return err
	}
	remotePath = zboxutil.GetFullRemotePath(localPath, remotePath)

	_, fileName := filepath.Split(remotePath)

	fileMeta := sdk.FileMeta{
		Path:       localPath,
		ActualSize: fileInfo.Size(),
		MimeType:   mimeType,
		RemoteName: fileName,
		RemotePath: remotePath,
		Attributes: attrs,
	}

	ChunkedUpload := sdk.CreateChunkedUpload(allocationObj, fileMeta, fileReader,
		sdk.WithThumbnailFile(thumbnailPath),
		sdk.WithChunkSize(chunkSize),
		sdk.WithEncrypt(encrypt),
		sdk.WithStatusCallback(statusBar))

	return ChunkedUpload.Start()
}

func startLiveUpload(cmd *cobra.Command, allocationObj *sdk.Allocation, localPath string, remotePath string, encrypt bool, chunkSize int, attrs fileref.Attributes) error {

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

	liveUpload := sdk.CreateLiveUpload(allocationObj, liveMeta, reader,
		sdk.WithLiveChunkSize(chunkSize),
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

func startSyncUpload(cmd *cobra.Command, allocationObj *sdk.Allocation, localPath, remotePath string, encrypt bool, chunkSize int, attrs fileref.Attributes) error {

	downloadArgs, _ := cmd.Flags().GetString("downloader-args")
	ffmpegArgs, _ := cmd.Flags().GetString("ffmpeg-args")
	delay, _ := cmd.Flags().GetInt("delay")
	feed, _ := cmd.Flags().GetString("feed")

	if len(feed) == 0 {
		return thrown.New("invalid_path", "feed should be valid")
	}

	reader, err := sdk.CreateYoutubeDL(localPath, feed, util.SplitArgs(downloadArgs), util.SplitArgs(ffmpegArgs), delay)
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

	syncUpload := sdk.CreateLiveUpload(allocationObj, liveMeta, reader,
		sdk.WithLiveChunkSize(chunkSize),
		sdk.WithLiveEncrypt(encrypt),
		sdk.WithLiveStatusCallback(func() sdk.StatusCallback {
			wg := &sync.WaitGroup{}
			statusBar := &StatusBar{wg: wg}
			wg.Add(1)

			return statusBar
		}),
		sdk.WithLiveDelay(delay))

	return syncUpload.Start()
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(createDirCmd)
	uploadCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	uploadCmd.PersistentFlags().String("remotepath", "", "Remote path to upload")
	uploadCmd.PersistentFlags().String("localpath", "", "Local path of file to upload")
	uploadCmd.PersistentFlags().String("thumbnailpath", "", "Local thumbnail path of file to upload")
	uploadCmd.Flags().Bool("encrypt", false, "pass this option to encrypt and upload the file")
	uploadCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")

	uploadCmd.Flags().Bool("stream", false, "pass this option to enable stream upload for large file")

	uploadCmd.Flags().Int("chunksize", sdk.CHUNK_SIZE, "how much bytes in a chunk for upload")

	uploadCmd.Flags().Bool("live", false, "pass this option to enable upload for live streaming")
	uploadCmd.Flags().Int("delay", 5, "how much seconds has a clips. default is 5 sencods. only works with --live")

	uploadCmd.Flags().Bool("sync", false, "sync live stream from remote live feed(eg. youtube), and upload to zcn")
	uploadCmd.Flags().String("feed", "", "remote live stream feed (eg youtube live feed url). only works with --sync")
	uploadCmd.Flags().String("downloader-args", "-q", "give args to youtube-dl to download video. default is -q. only works with --sync")
	uploadCmd.Flags().String("ffmpeg-args", "", "give args to ffmpeg to build output. only works with --sync")
	uploadCmd.MarkFlagRequired("allocation")
	uploadCmd.MarkFlagRequired("remotepath")
	uploadCmd.MarkFlagRequired("localpath")

	createDirCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	createDirCmd.PersistentFlags().String("dirname", "", "New directory name")
	createDirCmd.MarkFlagRequired("allocation")
	createDirCmd.MarkFlagRequired("dirname")

}
