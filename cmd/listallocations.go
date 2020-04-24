package cmd

import (
	"os"
	"strconv"
	"time"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

var listallocationsCmd = &cobra.Command{
	Use:   "listallocations",
	Short: "List allocations for the client",
	Long:  `List allocations for the client`,
	Run: func(cmd *cobra.Command, args []string) {
		allocations, err := sdk.GetAllocations()
		if err != nil {
			PrintError("Error getting allocations list." + err.Error())
			os.Exit(1)
		}
		header := []string{"ID", "Size", "Expiration", "Datashards",
			"Parityshards", "Finalized", "Canceled", "R. Price", "W. Price"}
		data := make([][]string, len(allocations))
		for idx, allocation := range allocations {
			size := strconv.FormatInt(allocation.Size, 10)
			expStr := strconv.FormatInt(allocation.Expiration, 10)
			exp, err := strconv.ParseInt(expStr, 10, 64)
			if err == nil {
				tm := time.Unix(exp, 0)
				expStr = tm.String()
			}
			d := strconv.FormatInt(int64(allocation.DataShards), 10)
			p := strconv.FormatInt(int64(allocation.ParityShards), 10)

			// TODO (sfxdx): data shards, parity shards

			var rp, wp common.Balance
			for _, d := range allocation.BlobberDetails {
				rp += d.Terms.ReadPrice
				wp += d.Terms.WritePrice
			}

			data[idx] = []string{
				allocation.ID, size, expStr, d, p,
				strconv.FormatBool(allocation.Finalized),
				strconv.FormatBool(allocation.Canceled),
				rp.String(), wp.String(),
			}
		}
		util.WriteTable(os.Stdout, header, []string{}, data)
		return
	},
}

func init() {
	rootCmd.AddCommand(listallocationsCmd)
}
