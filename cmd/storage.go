package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func printStorageSCConfig(conf *sdk.StorageSCConfig) {
	fmt.Println("min_alloc_size:               ", conf.MinAllocSize)
	fmt.Println("min_alloc_duration:           ", conf.MinAllocDuration)
	fmt.Println("max_challenge_completion_time:", conf.MaxChallengeCompletionTime)
	fmt.Println("min_offer_duration:           ", conf.MinOfferDuration)
	fmt.Println("min_blobber_capacity:         ", conf.MinBlobberCapacity)
	fmt.Println("readpool:")
	fmt.Println("  min_lock:", conf.ReadPool.MinLock, "tok")
	fmt.Println("  min_lock_period:", conf.ReadPool.MinLockPeriod)
	fmt.Println("  max_lock_period:", conf.ReadPool.MaxLockPeriod)
	fmt.Println("writepool:")
	fmt.Println("  min_lock:", conf.WritePool.MinLock, "tok")
	fmt.Println("  min_lock_period:", conf.WritePool.MinLockPeriod)
	fmt.Println("  max_lock_period:", conf.WritePool.MaxLockPeriod)
	fmt.Println("stakepool:")
	fmt.Println("  min_lock:", conf.StakePool.MinLock, "tok")
	fmt.Println("  interest_rate:", conf.StakePool.InterestRate)
	fmt.Println("  interest_interval:", conf.StakePool.InterestInterval)
	fmt.Println("validator_reward:                    ", conf.ValidatorReward)
	fmt.Println("blobber_slash:                       ", conf.BlobberSlash)
	fmt.Println("max_read_price:                      ", conf.MaxReadPrice, "tok / GB")
	fmt.Println("max_write_price:                     ", conf.MaxWritePrice, "tok / GB")
	fmt.Println("failed_challenges_to_cancel:         ", conf.FailedChallengesToCancel)
	fmt.Println("failed_challenges_to_revoke_min_lock:", conf.FailedChallengesToRevokeMinLock)
	fmt.Println("challenge_enabled:                   ", conf.ChallengeEnabled)
	fmt.Println("max_challenges_per_generation:       ", conf.MaxChallengesPerGeneration)
	fmt.Println("challenge_rate_per_mb_min:           ", conf.ChallengeGenerationRate)
}

// scConfig shows SC configurations
var scConfig = &cobra.Command{
	Use:   "sc-config",
	Short: "Show storage SC configuration.",
	Long:  `Show storage SC configuration.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")
		var conf, err = sdk.GetStorageSCConfig()
		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
		}
		if doJSON {
			util.PrintJSON(conf)
			return
		}
		printStorageSCConfig(conf)
	},
}

func printBlobbers(nodes []*sdk.Blobber) {
	fmt.Println("Blobbers:")
	header := []string{
		"URL", "ID", "CAP", "R / W PRICE", "DEMAND",
	}
	data := make([][]string, len(nodes))
	for i, val := range nodes {
		data[i] = []string{
			val.BaseURL,
			string(val.ID),
			fmt.Sprintf("%s / %s",
				val.Used.String(), val.Capacity.String()),
			fmt.Sprintf("%s / %s",
				val.Terms.ReadPrice.String(), val.Terms.WritePrice.String()),
			fmt.Sprint(val.Terms.MinLockDemand),
		}
	}
	util.WriteTable(os.Stdout, header, []string{}, data)
	fmt.Println("")
}

// lsBlobers shows active blobbers
var lsBlobers = &cobra.Command{
	Use:   "ls-blobbers",
	Short: "Show active blobbers in storage SC.",
	Long:  `Show active blobbers in storage SC.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")
		var list, err = sdk.GetBlobbers()
		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
		}
		if doJSON {
			util.PrintJSON(list)
			return
		}
		printBlobbers(list)
	},
}

func init() {
	rootCmd.AddCommand(scConfig)
	rootCmd.AddCommand(lsBlobers)
	scConfig.Flags().Bool("json", false, "pass this option to print response as json data")
	lsBlobers.Flags().Bool("json", false, "pass this option to print response as json data")

	scConfig.PersistentFlags().String("allocation", "",
		"allocation identifier, required")
}
