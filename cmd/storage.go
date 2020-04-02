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
	fmt.Printf(`
challenge_enabled:                %t
challenge_rate_per_mb_min:        %v
min_alloc_size:                   %s
min_alloc_duration:               %v
max_challenge_completion_time:    %v
min_offer_duration:               %v
min_blobber_capacity:             %v
readpool:
  min_lock:         %g tok
  min_lock_period:  %v
  max_lock_period:  %v
writepool:
  min_lock:         %g tok
validator_reward:   %f
blobber_slash:      %f
`,
		conf.ChallengeEnabled,
		conf.ChallengeRatePerMBMin,
		conf.MinAllocSize.String(),
		conf.MinAllocDuration,
		conf.MaxChallengeCompletionTime,
		conf.MinOfferDuration,
		conf.MinBlobberCapacity.String(),
		conf.ReadPool.MinLock.String(),
		conf.ReadPool.MinLockPeriod,
		conf.ReadPool.MaxLockPeriod,
		conf.WritePool.MinLock.String(),
		conf.ValidatorReward,
		conf.BlobberSlash,
	)
}

// scConfig shows SC configurations
var scConfig = &cobra.Command{
	Use:   "sc-config",
	Short: "Show storage SC configuration.",
	Long:  `Show storage SC configuration.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var conf, err = sdk.GetStorageSCConfig()
		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
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
			val.ID,
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
		var list, err = sdk.GetBlobbers()
		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
		}
		printBlobbers(list)
	},
}

func init() {
	rootCmd.AddCommand(scConfig)
	rootCmd.AddCommand(lsBlobers)

	scConfig.PersistentFlags().String("allocation", "",
		"allocation identifier, required")
}
