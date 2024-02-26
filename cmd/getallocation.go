package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/blockchain"
	"github.com/0chain/gosdk/zboxcore/logger"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

// getallocationCmd represents the get allocation info command
var getallocationCmd = &cobra.Command{
	Use:   "getallocation",
	Short: "Gets the allocation info",
	Long:  `Gets the allocation info`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		doJSON, _ := cmd.Flags().GetBool("json")
		alloc, err := sdk.GetAllocation(allocationID)
		if err != nil {
			logger.Logger.Error("Error fetching the allocation", err)
			log.Fatal("Error fetching/verifying the allocation")
		}
		if doJSON {
			util.PrintJSON(alloc)
			return
		}
		var getBaseURL = func(bid string, bs []*blockchain.StorageNode) string {
			for _, b := range bs {
				if b.ID == bid {
					return b.Baseurl
				}
			}
			return "(not found)"
		}

		var priceRangeString = func(pr sdk.PriceRange) string {
			return fmt.Sprintf("%s-%s", common.Balance(pr.Min), common.Balance(pr.Max))
		}

		blocksPerMarker, err := cmd.Flags().GetInt("blocks-per-marker")
		if err != nil {
			log.Fatal("invalid blocks-per-marker. Error: ", err)
		}

		if blocksPerMarker <= 0 {
			log.Fatal("invalid blocks-per-marker. Should be greater than 0")
		}

		fmt.Println("allocation:")
		fmt.Println("  id:             ", alloc.ID)
		fmt.Println("  tx:             ", alloc.Tx, "(latest create/update allocation transaction hash)")
		fmt.Println("  data_shards:    ", alloc.DataShards)
		fmt.Println("  parity_shards:  ", alloc.ParityShards)
		fmt.Println("  size:           ", common.Size(alloc.Size))
		fmt.Println("  expiration_date:", common.Timestamp(alloc.Expiration).ToTime())
		fmt.Println("  third_party_extendable:      ", alloc.ThirdPartyExtendable)
		fmt.Printf("  file_options:      %08b\n", alloc.FileOptions)
		fmt.Println("  write pool      ", alloc.WritePool)
		fmt.Println("  blobbers:")
		fmt.Println("  min_lock_demand:", alloc.MinLockDemand*100, "%")
		for _, d := range alloc.BlobberDetails {
			fmt.Println("    - blobber_id:      ", d.BlobberID)
			fmt.Println("      base URL:        ", getBaseURL(d.BlobberID, alloc.Blobbers))
			fmt.Println("      size:            ", common.Size(d.Size))
			fmt.Println("      min_lock_demand: ", common.Balance(d.MinLockDemand))
			fmt.Println("      spent:           ", common.Balance(d.Spent), "(moved to challenge pool or to the blobber)")
			fmt.Println("      penalty:         ", common.Balance(d.Penalty), "(blobber stake slash)")
			fmt.Println("      read_reward:     ", common.Balance(d.ReadReward))
			fmt.Println("      returned:        ", common.Balance(d.Returned), "(on challenge failed)")
			fmt.Println("      challenge_reward:", common.Balance(d.ChallengeReward), "(on challenge passed)")
			fmt.Println("      final_reward:    ", common.Balance(d.FinalReward), "(if finalized)")
			fmt.Println("      terms: (allocation related terms)")
			fmt.Println("        read_price:               ", d.Terms.ReadPrice, "/ GB (by 64KB chunks)")
			fmt.Println("        write_price:              ", d.Terms.WritePrice, "/ GB")
			fmt.Println("        min_lock_demand:          ", d.MinLockDemand*100, "%")
			fmt.Println("        max_offer_duration:       ", d.Terms.MaxOfferDuration)
		}

		fmt.Println("  read_price_range:         ", priceRangeString(alloc.ReadPriceRange), "(requested)")
		fmt.Println("  write_price_range:        ", priceRangeString(alloc.WritePriceRange), "(requested)")
		fmt.Println("  challenge_completion_time:", alloc.ChallengeCompletionTime, "(max)")
		fmt.Println("  start_time:               ", common.Timestamp(alloc.StartTime).ToTime())
		fmt.Println("  finalized:                ", alloc.Finalized)
		fmt.Println("  canceled:                 ", alloc.Canceled)
		fmt.Println("  moved_to_challenge:       ", common.Balance(alloc.MovedToChallenge))
		fmt.Println("  moved_back:               ", common.Balance(alloc.MovedBack))
		fmt.Println("  moved_to_validators:      ", common.Balance(alloc.MovedToValidators))

		fmt.Println("  stats:")
		fmt.Println("    total size:             ", common.Size(alloc.Size))
		fmt.Println("    used size:              ", common.Size(alloc.Stats.UsedSize))
		fmt.Println("    number of writes:       ", alloc.Stats.NumWrites)
		fmt.Println("    total challenges:       ", alloc.Stats.TotalChallenges)
		fmt.Println("    passed challenges:      ", alloc.Stats.SuccessChallenges)
		fmt.Println("    failed challenges:      ", alloc.Stats.FailedChallenges)
		fmt.Println("    open challenges:        ", alloc.Stats.OpenChallenges)
		fmt.Println("    last challenge redeemed:", alloc.Stats.LastestClosedChallengeTxn)

		fmt.Println("  price:")
		fmt.Println("    time_unit:  ", alloc.TimeUnit)
		fmt.Println("    write_price:", uploadCostFor1GB(alloc), fmt.Sprintf("/ GB / %s", alloc.TimeUnit))
	},
}

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

func sizeInGB(size int64) float64 {
	return float64(size) / GB
}

func init() {
	rootCmd.AddCommand(getallocationCmd)

	getallocationCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	getallocationCmd.PersistentFlags().Int("blocks-per-marker", 10, "blocks signed per Read Marker")
	getallocationCmd.MarkFlagRequired("allocation")
	getallocationCmd.Flags().Bool("json", false, "(default false) pass this option to print response as json data")

}
