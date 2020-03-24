package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	. "github.com/0chain/gosdk/zboxcore/logger"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

// getallocationCmd represents the get allocation info command
var getallocationCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the allocation info",
	Long:  `Gets the allocation info`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			Logger.Error("Error fetching the allocation", err)
			PrintError("Error fetching/verifying the allocation")
			os.Exit(1)
		}
		stats := allocationObj.GetStats()
		statsBytes, _ := json.Marshal(stats)
		fmt.Println(string(statsBytes))
		fmt.Printf("ID : %v\n", allocationObj.ID)
		fmt.Printf("Data Shards  : %v\n", allocationObj.DataShards)
		fmt.Printf("Parity Shards  : %v\n", allocationObj.ParityShards)
		fmt.Printf("Expiration  : %v\n", time.Unix(allocationObj.Expiration, 0))
		fmt.Printf("Blobbers : \n")
		for _, blobber := range allocationObj.Blobbers {
			fmt.Printf("\t%v\n", blobber.Baseurl)
		}
		fmt.Printf("Requested read price range  : [%f, %f] token / read\n",
			zcncore.ConvertToToken(allocationObj.ReadPriceRange.Min),
			zcncore.ConvertToToken(allocationObj.ReadPriceRange.Max))
		fmt.Printf("Requested write price range : [%f, %f] token / GB\n",
			zcncore.ConvertToToken(allocationObj.WritePriceRange.Min),
			zcncore.ConvertToToken(allocationObj.WritePriceRange.Max))
		fmt.Printf("Challenge Completion Time   : %v\n", allocationObj.ChallengeCompletionTime)

		fmt.Printf("Stats : \n")
		fmt.Printf("\tTotal Size : %v\n", allocationObj.Size)
		fmt.Printf("\tUsed Size : %v\n", allocationObj.Stats.UsedSize)
		fmt.Printf("\tNumber of Writes : %v\n", allocationObj.Stats.NumWrites)
		fmt.Printf("\tTotal Challenges : %v\n", allocationObj.Stats.TotalChallenges)
		fmt.Printf("\tPassed Challenges : %v\n", allocationObj.Stats.SuccessChallenges)
		fmt.Printf("\tFailed Challenges : %v\n", allocationObj.Stats.FailedChallenges)
		fmt.Printf("\tOpen Challenges : %v\n", allocationObj.Stats.OpenChallenges)
		fmt.Printf("\tLast Challenge redeemed : %v\n", allocationObj.Stats.LastestClosedChallengeTxn)
		return
	},
}

func init() {
	rootCmd.AddCommand(getallocationCmd)
	getallocationCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	getallocationCmd.MarkFlagRequired("allocation")
}
