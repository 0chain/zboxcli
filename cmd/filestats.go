package cmd

import (
	"context"
	"os"

	"github.com/0chain/gosdk/zcnbridge/transaction"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

// statsCmd represents list command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "stats for file from blobbers",
	Long:  `stats for file from blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		// if fflags.Changed("allocation") == false { // check if the flag "path" is set
		// 	PrintError("Error: allocation flag is missing") // If not, we'll let the user know
		// 	os.Exit(1)                                      // and return
		// }
		// if fflags.Changed("remotepath") == false {
		// 	PrintError("Error: remotepath flag is missing")
		// 	os.Exit(1)
		// }
		// allocationID := cmd.Flag("allocation").Value.String()
		// doJSON, _ := cmd.Flags().GetBool("json")

		// allocationObj, err := sdk.GetAllocation(allocationID)
		// if err != nil {
		// 	PrintError("Error fetching the allocation", err)
		// 	os.Exit(1)
		// }
		// remotepath := cmd.Flag("remotepath").Value.String()
		// ref, err := allocationObj.GetFileStats(remotepath)
		// if err != nil {
		// 	PrintError(err.Error())
		// 	os.Exit(1)
		// }
		// if doJSON {
		// 	util.PrintJSON(ref)
		// 	return
		// }

		payload := &zcncore.AuthorizerHealthCheckPayload{
			ID: "0",
		}

		ctx := context.TODO()
		trx, err := transaction.AuthorizerHealthCheck(ctx, payload)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}

		err = trx.Verify(ctx)
		PrintError(err.Error())

		// header := []string{"Blobber", "Name", "Path", "Size", "Uploads", "Block Downloads", "Challenges", "Blockchain Aware"}
		// data := make([][]string, 0)
		// idx := 0
		// for k, v := range ref {
		// 	if v != nil {
		// 		size := strconv.FormatInt(v.Size, 10)
		// 		rowData := []string{k, v.Name, v.Path, size, strconv.FormatInt(v.NumUpdates, 10), strconv.FormatInt(v.NumBlockDownloads, 10), strconv.FormatInt(v.SuccessChallenges, 10), strconv.FormatBool(len(v.WriteMarkerRedeemTxn) > 0)}
		// 		data = append(data, rowData)
		// 		idx++
		// 	}
		// }

		// util.WriteTable(os.Stdout, header, []string{}, data)
		return
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	statsCmd.PersistentFlags().String("remotepath", "", "Remote path to list from")
	statsCmd.MarkFlagRequired("allocation")
	statsCmd.MarkFlagRequired("remotepath")
	statsCmd.Flags().Bool("json", false, "pass this option to print response as json data")
}
