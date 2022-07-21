package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

const PageLimit = 100

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
		fmt.Printf("\nHello world\n\n")
		if pageLimit < 0 || pageLimit > PageLimit {
			PrintError("Invalid page limit value. Should be in range (0,100]")
			os.Exit(1)
		}
		allocID, err := cmd.Flags().GetString("allocation")
		checkError(err)
		fromDate, err := cmd.Flags().GetDuration("from_date")
		checkError(err)
		doJSON, err := cmd.Flags().GetBool("json")
		checkError(err)

		if len(allocID) != 64 {
			PrintError("Invalid allocation id")
			os.Exit(1)
		}

		alloc, err := sdk.GetAllocation(allocID)
		checkError(err)

		d := time.Now().Unix() - int64(fromDate.Seconds())
		result, err := alloc.GetRecentlyAddedRefs(int(page), d, int(pageLimit))
		checkError(err)

		if doJSON {
			util.PrintJSON(result)
			return
		}

		fmt.Printf("\nRequested page:%d with page limit: %d, from date: %v ago\n",
			page, pageLimit, fromDate)

		fmt.Printf(""+
			"\nTotal Pages: %d"+
			"\nCurrent Page: %d"+
			"\nRetrieved Refs: %d"+
			"\nNew Offset: %d\n",
			result.TotalPages, page, len(result.Refs), result.Offset,
		)

	},
}

func init() {
	rootCmd.AddCommand(fileRefsCmd)
	fileRefsCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	fileRefsCmd.PersistentFlags().Uint("page", 0, "Page to get refs from")
	fileRefsCmd.PersistentFlags().Duration("from_date", 0, "Date to consider refs was added recently")
	fileRefsCmd.PersistentFlags().Uint("page_limit", 0, "Number of refs to return in the page")
	fileRefsCmd.Flags().Bool("json", false, "pass this option to print response as json data")
	fileRefsCmd.MarkFlagRequired("allocation")
}
