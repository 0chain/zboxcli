package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

// filemetaCmd represents file meta command
var filemetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "get meta data of files from blobbers",
	Long:  `get meta data of files from blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags() // fflags is a *flag.FlagSet
		if fflags.Changed("remotepath") == false && fflags.Changed("authticket") == false {
			fmt.Println("Error: remotepath / authticket flag is missing")
			return
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		authticket := cmd.Flag("authticket").Value.String()
		lookuphash := cmd.Flag("lookuphash").Value.String()
		if len(remotepath) == 0 && (len(authticket) == 0) {
			fmt.Println("Error: remotepath / authticket / lookuphash flag is missing")
			return
		}

		if len(remotepath) > 0 {
			if fflags.Changed("allocation") == false { // check if the flag "path" is set
				fmt.Println("Error: allocation flag is missing") // If not, we'll let the user know
				return                                           // and return
			}
			allocationID := cmd.Flag("allocation").Value.String()
			allocationObj, err := sdk.GetAllocation(allocationID)
			if err != nil {
				fmt.Println("Error fetching the allocation", err)
				return
			}
			remotepath := cmd.Flag("remotepath").Value.String()
			ref, err := allocationObj.GetFileMeta(remotepath)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			header := []string{"Type", "Name", "Path", "Lookup Hash"}
			data := make([][]string, 1)
			data[0] = []string{ref.Type, ref.Name, ref.Path, ref.LookupHash}
			if ref.Type == fileref.FILE {
				headerFile := []string{"Size", "Mime Type", "Thumbnail Size", "Hash", "Thumbnail Hash"}
				dataFile := []string{strconv.FormatInt(ref.Size, 10), ref.MimeType, strconv.FormatInt(ref.ThumbnailSize, 10), ref.Hash, ref.ThumbnailHash}
				header = append(header, headerFile...)
				data[0] = append(data[0], dataFile...)
			}

			util.WriteTable(os.Stdout, header, []string{}, data)
		} else if len(authticket) > 0 {
			allocationObj, err := sdk.GetAllocationFromAuthTicket(authticket)
			if err != nil {
				fmt.Println("Error fetching the allocation", err)
				return
			}
			at := sdk.InitAuthTicket(authticket)
			if len(lookuphash) == 0 {
				lookuphash, err = at.GetLookupHash()
				if err != nil {
					fmt.Println("Error getting the lookuphash from authticket", err)
					return
				}
			}

			ref, err := allocationObj.GetFileMetaFromAuthTicket(authticket, lookuphash)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			header := []string{"Type", "Name", "Lookup Hash"}
			data := make([][]string, 1)
			data[0] = []string{ref.Type, ref.Name, ref.LookupHash}
			if ref.Type == fileref.FILE {
				headerFile := []string{"Size", "Mime Type", "Thumbnail Size", "Hash", "Thumbnail Hash"}
				dataFile := []string{strconv.FormatInt(ref.Size, 10), ref.MimeType, strconv.FormatInt(ref.ThumbnailSize, 10), ref.Hash, ref.ThumbnailHash}
				header = append(header, headerFile...)
				data[0] = append(data[0], dataFile...)
			}
			util.WriteTable(os.Stdout, header, []string{}, data)
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(filemetaCmd)
	filemetaCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	filemetaCmd.PersistentFlags().String("remotepath", "", "Remote path to list from")
	filemetaCmd.PersistentFlags().String("authticket", "", "Auth ticket fot the file to download if you dont own it")
	filemetaCmd.PersistentFlags().String("lookuphash", "", "The remote lookuphash of the object retrieved from the list")
	filemetaCmd.MarkFlagRequired("allocation")
}
