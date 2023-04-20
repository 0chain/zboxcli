package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
			PrintError("Error: remotepath / authticket flag is missing")
			os.Exit(1)
		}

		remotepath := cmd.Flag("remotepath").Value.String()
		authticket := cmd.Flag("authticket").Value.String()
		lookuphash := cmd.Flag("lookuphash").Value.String()
		doJSON, _ := cmd.Flags().GetBool("json")
		if len(remotepath) == 0 && (len(authticket) == 0) {
			PrintError("Error: remotepath / authticket / lookuphash flag is missing")
			os.Exit(1)
		}

		if len(remotepath) > 0 {
			if fflags.Changed("allocation") == false { // check if the flag "path" is set
				PrintError("Error: allocation flag is missing") // If not, we'll let the user know
				os.Exit(1)                                      // and os.Exit(1)
			}

			allocationID := cmd.Flag("allocation").Value.String()
			allocationObj, err := sdk.GetAllocation(allocationID)
			if err != nil {
				PrintError("Error fetching the allocation", err)
				os.Exit(1)
			}

			remotepath := cmd.Flag("remotepath").Value.String()
			ref, err := allocationObj.ListDir(remotepath)
			if err != nil {
				PrintError(err.Error())
				os.Exit(1)
			}

			printListDirResult(doJSON, ref)
		} else if len(authticket) > 0 {
			allocationObj, err := sdk.GetAllocationFromAuthTicket(authticket)
			if err != nil {
				PrintError("Error fetching the allocation", err)
				os.Exit(1)
			}

			at := sdk.InitAuthTicket(authticket)
			lookuphash, err = at.GetLookupHash()
			if err != nil {
				PrintError("Error getting the lookuphash from authticket", err)
				os.Exit(1)
			}

			ref, err := allocationObj.ListDirFromAuthTicket(authticket, lookuphash)
			if err != nil {
				PrintError(err.Error())
				os.Exit(1)
			}

			printListDirResult(doJSON, ref)
		}

		return
	},
}

var listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "list all files from blobbers",
	Long:  `list all files from blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}

		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		ref, err := allocationObj.GetRemoteFileMap(nil)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}

		type fileResp struct {
			sdk.FileInfo
			Name string `json:"name"`
			Path string `json:"path"`
		}

		fileResps := make([]fileResp, 0)
		for path, data := range ref {
			paths := strings.SplitAfter(path, "/")
			var resp = fileResp{
				Name:     paths[len(paths)-1],
				Path:     path,
				FileInfo: data,
			}
			fileResps = append(fileResps, resp)
		}

		util.PrintJSON(fileResps)
		return
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	listCmd.PersistentFlags().String("remotepath", "", "Remote path to list from")
	listCmd.PersistentFlags().String("authticket", "", "Auth ticket fot the file to download if you dont own it")
	listCmd.PersistentFlags().String("lookuphash", "", "The remote lookuphash of the object retrieved from the list")
	listCmd.Flags().Bool("json (boolean)", false, "pass this option to print response as json data")
	listCmd.MarkFlagRequired("allocation")

	rootCmd.AddCommand(listAllCmd)
	listAllCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	listAllCmd.MarkFlagRequired("allocation")
}

func printListDirResult(outJson bool, ref *sdk.ListResult) {
	if outJson {
		util.PrintJSON(ref.Children)
		return
	}

	header := []string{"Type", "Name", "Path", "Size", "Actual Size", "Num Blocks", "Lookup Hash", "Is Encrypted"}
	data := make([][]string, len(ref.Children))
	for idx, child := range ref.Children {
		size := strconv.FormatInt(child.Size, 10)
		if child.Type == fileref.DIRECTORY {
			size = ""
		}
		isEncrypted := ""
		if child.Type == fileref.FILE {
			if len(child.EncryptionKey) > 0 {
				isEncrypted = "YES"
			} else {
				isEncrypted = "NO"
			}
		}
		data[idx] = []string{
			child.Type,
			child.Name,
			child.Path,
			size,
			fmt.Sprint(child.ActualSize),
			strconv.FormatInt(child.NumBlocks, 10),
			child.LookupHash,
			isEncrypted,
		}
	}

	util.WriteTable(os.Stdout, header, []string{}, data)
}
