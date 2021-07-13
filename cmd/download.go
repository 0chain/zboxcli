package cmd

import (
	"os"
	"sync"

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
		if fflags.Changed("remotepath") == false && fflags.Changed("authticket") == false {
			PrintError("Error: remotepath / authticket flag is missing")
			os.Exit(1)
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		authticket := cmd.Flag("authticket").Value.String()
		lookuphash := cmd.Flag("lookuphash").Value.String()
		thumbnail, _ := cmd.Flags().GetBool("thumbnail")
		commit, _ := cmd.Flags().GetBool("commit")
		rxPay, _ := cmd.Flags().GetBool("rx_pay")
		if len(remotepath) == 0 && len(authticket) == 0 {
			PrintError("Error: remotepath / authticket flag is missing")
			os.Exit(1)
		}

		localpath := cmd.Flag("localpath").Value.String()
		allocationID := cmd.Flag("allocation").Value.String()

		live, _ := cmd.Flags().GetBool("live")

		if live {
			delay, _ := cmd.Flags().GetInt("delay")

			m3u8, err := createM3u8Downloader(localpath, remotepath, authticket, allocationID, lookuphash, rxPay, delay)

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
		endBlock, _ := cmd.Flags().GetInt64("endblock")

		sdk.SetNumBlockDownloads(numBlocks)
		wg := &sync.WaitGroup{}
		statusBar := &StatusBar{wg: wg}
		wg.Add(1)
		var errE, err error
		var allocationObj *sdk.Allocation
		if len(remotepath) > 0 {
			if fflags.Changed("allocation") == false { // check if the flag "path" is set
				PrintError("Error: allocation flag is missing") // If not, we'll let the user know
				os.Exit(1)                                      // and return
			}

			allocationObj, err = sdk.GetAllocation(allocationID)

			if err != nil {
				PrintError("Error fetching the allocation", err)
				os.Exit(1)
			}
			if thumbnail {
				errE = allocationObj.DownloadThumbnail(localpath, remotepath, statusBar)
			} else {
				if startBlock != 0 || endBlock != 0 {
					errE = allocationObj.DownloadFileByBlock(localpath, remotepath, startBlock, endBlock, numBlocks, statusBar)
				} else {
					errE = allocationObj.DownloadFile(localpath, remotepath, statusBar)
				}
			}
		} else if len(authticket) > 0 {
			allocationObj, err = sdk.GetAllocationFromAuthTicket(authticket)
			if err != nil {
				PrintError("Error fetching the allocation", err)
				os.Exit(1)
			}
			at := sdk.InitAuthTicket(authticket)
			filename, err := at.GetFileName()
			if err != nil {
				PrintError("Error getting the filename from authticket", err)
				os.Exit(1)
			}
			if len(lookuphash) == 0 {
				lookuphash, err = at.GetLookupHash()
				if err != nil {
					PrintError("Error getting the lookuphash from authticket", err)
					os.Exit(1)
				}
			}

			if thumbnail {
				errE = allocationObj.DownloadThumbnailFromAuthTicket(localpath,
					authticket, lookuphash, filename, rxPay, statusBar)
			} else {
				if startBlock != 0 || endBlock != 0 {
					errE = allocationObj.DownloadFromAuthTicketByBlocks(
						localpath, authticket, startBlock, endBlock, numBlocks,
						lookuphash, filename, rxPay, statusBar)
				} else {
					errE = allocationObj.DownloadFromAuthTicket(localpath,
						authticket, lookuphash, filename, rxPay, statusBar)
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
			os.Exit(1)
		}
		if commit {
			statusBar.wg.Add(1)
			commitMetaTxn(remotepath, "Download", authticket, lookuphash, allocationObj, nil, statusBar)
			statusBar.wg.Wait()
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	downloadCmd.PersistentFlags().String("remotepath", "", "Remote path to download")
	downloadCmd.PersistentFlags().String("localpath", "", "Local path of file to download")
	downloadCmd.PersistentFlags().String("authticket", "", "Auth ticket fot the file to download if you dont own it")
	downloadCmd.PersistentFlags().String("lookuphash", "", "The remote lookuphash of the object retrieved from the list")
	downloadCmd.Flags().BoolP("thumbnail", "t", false, "pass this option to download only the thumbnail")
	downloadCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	downloadCmd.Flags().Bool("rx_pay", false, "used to download by authticket; pass true to pay for download yourself")
	downloadCmd.Flags().Int64P("startblock", "s", 0, "pass this option to download from specific block number")
	downloadCmd.Flags().Int64P("endblock", "e", 0, "pass this option to download till specific block number")
	downloadCmd.Flags().IntP("blockspermarker", "b", 10, "pass this option to download multiple blocks per marker")
	downloadCmd.Flags().Bool("live", false, "enable m3u8 downloader,and build playlist on --localpath")
	downloadCmd.Flags().Int("delay", 5, "how many seconds has a clips. default is 5 sencods. only works with --live")

	downloadCmd.MarkFlagRequired("allocation")
	downloadCmd.MarkFlagRequired("localpath")
}
