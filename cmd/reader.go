package cmd

import (
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// readerCmd represents reader command that downloads file from io.Reader interface provided in
// gosdk
var readerCmd = &cobra.Command{
	Use:   "read",
	Short: "read and store file from blobbers",
	Long: `This command will use io.ReadSeekCloser interface provided by Allocation object in gosdk to read file
	from blobbers`,
	Args: cobra.MinimumNArgs(0),
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

		contentMode := sdk.DOWNLOAD_CONTENT_FULL
		if thumbnail {
			contentMode = sdk.DOWNLOAD_CONTENT_THUMB
		}

		localPath := cmd.Flag("localpath").Value.String()
		allocationID := cmd.Flag("allocation").Value.String()

		numBlocks, _ := cmd.Flags().GetInt("blockspermarker")
		wg := &sync.WaitGroup{}
		wg.Add(1)
		var allocationObj *sdk.Allocation
		if authTicket != "" {
			allocationObj, err = sdk.GetAllocationFromAuthTicket(authTicket)
		} else {
			if allocationID == "" {
				PrintError("Both authtoken and allocationID are empty")
				os.Exit(1)
			}
			allocationObj, err = sdk.GetAllocation(allocationID)
		}

		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}

		err = allocationObj.DownloadFromReader(
			remotePath, localPath, lookupHash, authTicket, contentMode, verifyDownload, uint(numBlocks))
		if err != nil {
			PrintError("Error: ", err)
			os.Exit(1)
		}
		PrintError("Download successful")
	},
}

func init() {
	rootCmd.AddCommand(readerCmd)
	readerCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	readerCmd.PersistentFlags().String("remotepath", "", "Remote path to download")
	readerCmd.PersistentFlags().String("localpath", "", "Local path of file to download")
	readerCmd.PersistentFlags().String("authticket", "", "Auth ticket fot the file to download if you dont own it")
	readerCmd.PersistentFlags().String("lookuphash", "", "The remote lookuphash of the object retrieved from the list")
	readerCmd.Flags().BoolP("thumbnail", "t", false, "pass this option to download only the thumbnail")

	readerCmd.Flags().IntP("blockspermarker", "b", 10, "pass this option to download multiple blocks per marker")
	readerCmd.Flags().BoolP("verifydownload", "v", false, "pass this option to verify downloaded blocks")

	readerCmd.MarkFlagRequired("allocation")
	readerCmd.MarkFlagRequired("localpath")
}
