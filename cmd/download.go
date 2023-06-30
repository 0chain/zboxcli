package cmd

import (
	"os"
	"strings"
	"sync"

	"github.com/0chain/gosdk/zboxcore/fileref"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// downloadCmd represents download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download file from blobbers",
	Long:  `download file from blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags() // fflags is a *flag.FlagSet
		if !(fflags.Changed("remotepath") || fflags.Changed("authticket")) {
			PrintError("Error: remotepath / authticket flag is missing")
			os.Exit(1)
		}

		remotePath := cmd.Flag("remotepath").Value.String()
		authTicket := cmd.Flag("authticket").Value.String()
		lookupHash := cmd.Flag("lookuphash").Value.String()
		verifyDownload, err := cmd.Flags().GetBool("verifydownload")
		if err != nil {
			PrintError("Error: ", err)
			os.Exit(1)
		}

		thumbnail, err := cmd.Flags().GetBool("thumbnail")
		if err != nil {
			PrintError("Error: ", err)
			os.Exit(1)
		}

		localPath := cmd.Flag("localpath").Value.String()
		allocationID := cmd.Flag("allocation").Value.String()

		live, _ := cmd.Flags().GetBool("live")

		if live {
			delay, _ := cmd.Flags().GetInt("delay")

			m3u8, err := createM3u8Downloader(localPath, remotePath, authTicket, allocationID, lookupHash, delay)

			if err != nil {
				PrintError("Error: download files and build playlist: ", err)
				os.Exit(1)
			}

			err = m3u8.Start()

			if err != nil {
				PrintError("Error: download files and build playlist: ", err)
				os.Exit(1)
			}

			return

		}

		numBlocks, _ := cmd.Flags().GetInt("blockspermarker")
		if numBlocks == 0 {
			numBlocks = 10
		}

		startBlock, _ := cmd.Flags().GetInt64("startblock")
		if startBlock < 1 {
			PrintError("Error: start block should not be less than 1")
		}
		endBlock, _ := cmd.Flags().GetInt64("endblock")

		sdk.SetNumBlockDownloads(numBlocks)
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		var errE error
		var allocationObj *sdk.Allocation

		if len(authTicket) > 0 {
			at, err := sdk.InitAuthTicket(authTicket).Unmarshall()

			if err != nil {
				PrintError(err)
				os.Exit(1)
			}

			allocationObj, err = sdk.GetAllocationFromAuthTicket(authTicket)
			if err != nil {
				PrintError("Error fetching the allocation", err)
				os.Exit(1)
			}

			var fileName string

			if at.RefType == fileref.FILE {
				fileName = at.FileName
				lookupHash = at.FilePathHash
			} else if len(lookupHash) > 0 {
				fileMeta, err := allocationObj.GetFileMetaFromAuthTicket(authTicket, lookupHash)
				if err != nil {
					PrintError("Either remotepath or lookuphash is required when using authticket of directory type")
					os.Exit(1)
				}
				fileName = fileMeta.Name
			} else if len(remotePath) > 0 {
				lookupHash = fileref.GetReferenceLookup(allocationObj.Tx, remotePath)

				pathNames := strings.Split(remotePath, "/")
				fileName = pathNames[len(pathNames)-1]
			} else {
				PrintError("Either remotepath or lookuphash is required when using authticket of directory type")
				os.Exit(1)
			}

			if thumbnail {
				errE = allocationObj.DownloadThumbnailFromAuthTicket(localPath,
					authTicket, lookupHash, fileName, verifyDownload, statusBar, true)
			} else {
				if startBlock != 0 || endBlock != 0 {
					errE = allocationObj.DownloadFromAuthTicketByBlocks(
						localPath, authTicket, startBlock, endBlock, numBlocks,
						lookupHash, fileName, verifyDownload, statusBar, true)
				} else {
					errE = allocationObj.DownloadFromAuthTicket(localPath,
						authTicket, lookupHash, fileName, verifyDownload, statusBar, true)
				}
			}
		} else if len(remotePath) > 0 {
			if fflags.Changed("allocation") == false { // check if the flag "path" is set
				PrintError("Error: allocation flag is missing") // If not, we'll let the user know
				os.Exit(1)                                      // and return
			}
			allocationID := cmd.Flag("allocation").Value.String()
			allocationObj, err = sdk.GetAllocation(allocationID)

			if err != nil {
				PrintError("Error fetching the allocation", err)
				os.Exit(1)
			}

			var blobberID string
			if fflags.Changed("blobber_id") {
				blobberID = cmd.Flag("blobber_id").Value.String()
			}

			if thumbnail {
				errE = allocationObj.DownloadThumbnail(localPath, remotePath, verifyDownload, statusBar, true)
			} else if blobberID != "" {
				errE = allocationObj.DownloadFromBlobber(blobberID, localPath, remotePath, statusBar)
			} else {
				if startBlock != 0 || endBlock != 0 {
					errE = allocationObj.DownloadFileByBlock(localPath, remotePath, startBlock, endBlock, numBlocks, verifyDownload, statusBar, true)
				} else {
					errE = allocationObj.DownloadFile(localPath, remotePath, verifyDownload, statusBar, true)
				}
			}
		}

		if errE == nil {
			wg.Wait()
		} else {
			PrintError("Download failed.", errE.Error())
			os.Exit(1)
		}
		if !statusBar.success {
			// status bar always returns failure when downloading from sigle blobber. Hence returning the zero exit status
			if fflags.Changed("blobber_id") {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	downloadCmd.PersistentFlags().String("remotepath", "", "Remote path to download")
	downloadCmd.PersistentFlags().String("localpath", "", "Local path of file to download")
	downloadCmd.PersistentFlags().String("blobber_id", "", "to download the data shard present in that blobber")
	downloadCmd.PersistentFlags().String("authticket", "", "Auth ticket fot the file to download if you dont own it")
	downloadCmd.PersistentFlags().String("lookuphash", "", "The remote lookuphash of the object retrieved from the list")
	downloadCmd.Flags().BoolP("thumbnail", "t", false, "(default false) pass this option to download only the thumbnail")

	downloadCmd.Flags().Int64P("startblock", "s", 1,
		"Pass this option to download from specific block number. It should not be less than 1")
	downloadCmd.Flags().Int64P("endblock", "e", 0, "pass this option to download till specific block number")
	downloadCmd.Flags().IntP("blockspermarker", "b", 10, "pass this option to download multiple blocks per marker")
	downloadCmd.Flags().BoolP("verifydownload", "v", false, "(default false) pass this option to verify downloaded blocks")

	downloadCmd.Flags().Bool("live", false, "(default false) start m3u8 downloader,and automatically generate media playlist(m3u8) on --localpath")
	downloadCmd.Flags().Int("delay", 5, "pass segment duration to generate media playlist(m3u8). only works with --live. default duration is 5s.")

	downloadCmd.MarkFlagRequired("allocation")
	downloadCmd.MarkFlagRequired("localpath")
}
