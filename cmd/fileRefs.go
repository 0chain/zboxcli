package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

const (
	PageLimit = 100
	Layout    = "2006-01-02 15:04:05"
)

func checkError(err error) {
	if err != nil {
		PrintError(err)
		os.Exit(1)
	}
}

var fileRefsCmd = &cobra.Command{
	Use:   "recent-refs",
	Short: "get list of recently added refs",
	Long:  `get list of recently added refs`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		page, err := cmd.Flags().GetUint("page")
		checkError(err)
		pageLimit, err := cmd.Flags().GetUint("page_limit")
		checkError(err)
		if pageLimit > PageLimit {
			PrintError("Invalid page limit value. Should be in range (0,100]")
			os.Exit(1)
		}
		allocID, err := cmd.Flags().GetString("allocation")
		checkError(err)
		in, err := cmd.Flags().GetDuration("in")
		checkError(err)
		doJSON, err := cmd.Flags().GetBool("json")
		checkError(err)

		if len(allocID) != 64 {
			PrintError("Invalid allocation id")
			os.Exit(1)
		}

		alloc, err := sdk.GetAllocation(allocID)
		checkError(err)

		d := time.Now().Unix() - int64(in.Seconds())
		result, err := alloc.GetRecentlyAddedRefs(int(page), d, int(pageLimit))
		checkError(err)

		if doJSON {
			util.PrintJSON(result)
			return
		}

		fmt.Printf("\nRequested page:%d with page limit: %d, in last: %v ago\n",
			page, pageLimit, in)

		fmt.Printf(""+
			"\nCurrent Page: %d"+
			"\nRetrieved Refs: %d"+
			"\nNew Offset: %d\n",
			page, len(result.Refs), result.Offset,
		)

		header := []string{"Type", "Name", "Path", "Size", "Lookup Hash", "Created At"}
		data := make([][]string, len(result.Refs))
		for i, ref := range result.Refs {
			var size string
			if ref.Type != fileref.DIRECTORY {
				size = strconv.FormatInt(ref.ActualFileSize, 10)
			}

			var createdAt string
			t := time.Unix(int64(ref.CreatedAt), 0)
			createdAt = t.Local().Format(Layout)
			data[i] = []string{
				ref.Type,
				ref.Name,
				ref.Path,
				size,
				ref.LookupHash,
				createdAt,
			}

		}

		util.WriteTable(os.Stdout, header, []string{}, data)
	},
}

func init() {
	rootCmd.AddCommand(fileRefsCmd)
	fileRefsCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	fileRefsCmd.PersistentFlags().Uint("page", 0, "Page to get refs from")
	fileRefsCmd.PersistentFlags().Duration("in", 0, "Recent refs in this duration")
	fileRefsCmd.PersistentFlags().Uint("page_limit", 0, "Number of refs to return in the page")
	fileRefsCmd.Flags().Bool("json", false, "pass this option to print response as json data")
	fileRefsCmd.MarkFlagRequired("allocation")
}
