package cmd

import (
	// "encoding/json"
	"fmt"
	"os"
	// "time"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/blockchain"
	. "github.com/0chain/gosdk/zboxcore/logger"
	"github.com/0chain/gosdk/zboxcore/sdk"
	// "github.com/0chain/gosdk/zcncore"

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
		alloc, err := sdk.GetAllocation(allocationID)
		if err != nil {
			Logger.Error("Error fetching the allocation", err)
			PrintError("Error fetching/verifying the allocation")
			os.Exit(1)
		}

		var getBaseURL = func(bid string, bs []*blockchain.StorageNode) string {
			for _, b := range bs {
				if b.ID == bid {
					return b.Baseurl
				}
			}
			return "(not found)"
		}

		fmt.Println("allocation:")
		fmt.Println("  id:             ", alloc.ID)
		fmt.Println("  tx:             ", alloc.Tx, "(latest create/update allocation transaction hash)")
		fmt.Println("  data_shards:    ", alloc.DataShards)
		fmt.Println("  parity_shards:  ", alloc.ParityShards)
		fmt.Println("  size:           ", common.Size(alloc.Size))
		fmt.Println("  expiration_date:", common.Timestamp(alloc.Expiration).ToTime())
		fmt.Println("  blobbers:")

		for _, d := range alloc.BlobberDetails {
			fmt.Println("    - blobber_id:      ", d.BlobberID)
			fmt.Println("      base URL:        ", getBaseURL(d.BlobberID, alloc.Blobbers))
			fmt.Println("      size:            ", common.Size(d.Size))
			fmt.Println("      min_lock_deman:  ", common.Balance(d.MinLockDemand))
			fmt.Println("      spent:           ", common.Balance(d.Spent), "(moved to challenge pool or to the blobber)")
			fmt.Println("      penalty:         ", common.Balance(d.Penalty), "(blobber stake slash)")
			fmt.Println("      read_reward:     ", common.Balance(d.ReadReward))
			fmt.Println("      returned:        ", common.Balance(d.Returned), "(on challenge failed)")
			fmt.Println("      challenge_reward:", common.Balance(d.ChallengeReward), "(on challenge passed)")
			fmt.Println("      final_reward:    ", common.Balance(d.FinalReward), "(if finalized)")
			fmt.Println("      terms: (allocation related terms)")
			fmt.Println("        read_price:               ", d.Terms.ReadPrice, "tok / GB (by 64KB chunks)")
			fmt.Println("        write_price:              ", d.Terms.WritePrice, "tok / GB")
			fmt.Println("        min_lock_demand:          ", d.Terms.MinLockDemand*100, "%")
			fmt.Println("        max_offer_duration:       ", d.Terms.MaxOfferDuration)
			fmt.Println("        challenge_completion_time:", d.Terms.ChallengeCompletionTime)
		}

		fmt.Println("  read_price_range:         ", alloc.ReadPriceRange, "(requested)")
		fmt.Println("  write_price_range:        ", alloc.WritePriceRange, "(requested)")
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
		return
	},
}

func init() {
	rootCmd.AddCommand(getallocationCmd)
	getallocationCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	getallocationCmd.MarkFlagRequired("allocation")
}
