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

// listCmd represents list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list files from blobbers",
	Long:  `list files from blobbers`,
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
			ref, err := allocationObj.ListDir(remotepath)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			header := []string{"Type", "Name", "Path", "Size", "Num Blocks", "Lookup Hash"}
			data := make([][]string, len(ref.Children))
			for idx, child := range ref.Children {
				size := strconv.FormatInt(child.Size, 10)
				if child.Type == fileref.DIRECTORY {
					size = ""
				}
				data[idx] = []string{child.Type, child.Name, child.Path, size, strconv.FormatInt(child.NumBlocks, 10), child.LookupHash}
			}
			util.WriteTable(os.Stdout, header, []string{}, data)
		} else if len(authticket) > 0 {
			allocationObj, err := sdk.GetAllocationFromAuthTicket(authticket)
			if err != nil {
				fmt.Println("Error fetching the allocation", err)
				return
			}
			at := sdk.InitAuthTicket(authticket)
			isDir, err := at.IsDir()
			if isDir && len(lookuphash) == 0 {
				lookuphash, err = at.GetLookupHash()
				if err != nil {
					fmt.Println("Error getting the lookuphash from authticket", err)
					return
				}
			}
			if !isDir {
				fmt.Println("Invalid operation. Auth ticket is not for a directory")
				return
			}

			ref, err := allocationObj.ListDirFromAuthTicket(authticket, lookuphash)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			header := []string{"Type", "Name", "Size", "Num Blocks", "Lookup Hash"}
			data := make([][]string, len(ref.Children))
			for idx, child := range ref.Children {
				size := strconv.FormatInt(child.Size, 10)
				if child.Type == fileref.DIRECTORY {
					size = ""
				}
				data[idx] = []string{child.Type, child.Name, size, strconv.FormatInt(child.NumBlocks, 10), child.LookupHash}
			}
			util.WriteTable(os.Stdout, header, []string{}, data)
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	listCmd.PersistentFlags().String("remotepath", "", "Remote path to list from")
	listCmd.PersistentFlags().String("authticket", "", "Auth ticket fot the file to download if you dont own it")
	listCmd.PersistentFlags().String("lookuphash", "", "The remote lookuphash of the object retrieved from the list")
	listCmd.MarkFlagRequired("allocation")
}
