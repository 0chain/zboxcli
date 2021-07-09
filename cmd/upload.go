package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/0chain/gosdk/core/common"
	thrown "github.com/0chain/gosdk/core/common/errors"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zboxcore/zboxutil"
	"github.com/0chain/zboxcli/util"

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
		stream, _ := cmd.Flags().GetBool("stream")

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

		live, _ := cmd.Flags().GetBool("live")

		if live {
			chunkSize, _ := cmd.Flags().GetInt("chunksize")
			err = startLiveUpload(cmd, allocationObj, localpath, remotepath, encrypt, chunkSize, attrs)
		} else if stream {
			chunkSize, _ := cmd.Flags().GetInt("chunksize")
			err = startStreamUpload(cmd, allocationObj, localpath, thumbnailpath, remotepath, encrypt, chunkSize, attrs, statusBar)

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

func startStreamUpload(cmd *cobra.Command, allocationObj *sdk.Allocation, localPath, thumbnailPath, remotePath string, encrypt bool, chunkSize int, attrs fileref.Attributes, statusBar sdk.StatusCallback) error {

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

	streamUpload := sdk.CreateStreamUpload(allocationObj, fileMeta, fileReader,
		sdk.WithThumbnailFile(thumbnailPath),
		sdk.WithChunkSize(chunkSize),
		sdk.WithEncrypt(encrypt),
		sdk.WithStatusCallback(statusBar))

	return streamUpload.Start()
}

func startLiveUpload(cmd *cobra.Command, allocationObj *sdk.Allocation, localPath string, remotePath string, encrypt bool, chunkSize int, attrs fileref.Attributes) error {

	delay, _ := cmd.Flags().GetInt("delay")
	clipsSize, _ := cmd.Flags().GetInt("clipssize")

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
		sdk.WithLiveDelay(delay),
		sdk.WithLiveClipsSize(clipsSize))

	return liveUpload.Start()
}

func startLiveUploadWithYoutubeDL(cmd *cobra.Command, allocationObj *sdk.Allocation, feedURL string, remotePath string, encrypt bool, chunkSize int, attrs fileref.Attributes) error {

	format, _ := cmd.Flags().GetString("format")
	proxy, _ := cmd.Flags().GetString("proxy")
	delay, _ := cmd.Flags().GetInt("delay")
	clipsSize, _ := cmd.Flags().GetInt("clipssize")

	dir := util.GetConfigDir() + string(os.PathSeparator) + "youtubedl"

	os.MkdirAll(dir, 0744)

	reader, err := sdk.CreateYoutubeDL(feedURL, format, dir, proxy, delay)
	if err != nil {
		return err
	}

	defer reader.Close()

	mimeType := "video/mp4"

	remotePath = zboxutil.RemoteClean(remotePath)
	isabs := zboxutil.IsRemoteAbs(remotePath)
	if !isabs {
		err = thrown.New("invalid_path", "Path should be valid and absolute")
		return err
	}
	remotePath = zboxutil.GetFullRemotePath(reader.GetFileName(0), remotePath)

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
		sdk.WithLiveDelay(delay),
		sdk.WithLiveClipsSize(clipsSize))

	return liveUpload.Start()
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
	uploadCmd.Flags().Bool("stream", false, "pass this option to enable stream upload for large file")

	uploadCmd.Flags().Int("chunksize", sdk.CHUNK_SIZE, "how much bytes in a chunk for upload")

	uploadCmd.Flags().Bool("live", false, "pass this option to enable upload for live streaming")
	uploadCmd.Flags().Int("delay", 5, "how much seconds has a clips.default is 5 sencods. only works with --live")
	// uploadCmd.Flags().String("format", "best", "quality format of video. best is default. only works with --live")
	// uploadCmd.Flags().String("proxy", "", "Use the specified HTTP/HTTPS/SOCKS proxy. only works with --live")

	uploadCmd.MarkFlagRequired("allocation")
	uploadCmd.MarkFlagRequired("remotepath")
	uploadCmd.MarkFlagRequired("localpath")
}
