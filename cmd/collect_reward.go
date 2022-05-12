package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var collectRewards = &cobra.Command{
	Use:   "collect-reward",
	Short: "Collect accrued rewards for a stake pool.",
	Long:  "Collect accrued rewards for a stake pool.",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		if !flags.Changed("pool_id") && !flags.Changed("provider_id") {
			log.Fatal("must have pool id or provider id")
		}

		var poolId, providerId string
		var err error

		if flags.Changed("pool_id") {
			poolId, err = flags.GetString("pool_id")
			if err != nil {
				log.Fatal(err)
			}
		}

		if flags.Changed("provider_id") {
			providerId, err = flags.GetString("provider_id")
			if err != nil {
				log.Fatal(err)
			}
		}

		if !flags.Changed("provider_type") {
			log.Fatal("missing tokens flag")
		}

		providerName, err := flags.GetString("provider_type")
		if err != nil {
			log.Fatal(err)
		}

		switch providerName {
		case "blobber":
			_, _, err = sdk.CollectRewards(providerId, poolId, sdk.ProviderBlobber)
		case "validator":
			_, _, err = sdk.CollectRewards(providerId, poolId, sdk.ProviderValidator)
		default:
			log.Fatal("provider type must be blobber or validator")
		}
		if err != nil {
			log.Fatal("Error paying reward:", err)
		}
		log.Println("transferred reward tokens")
	},
}

func init() {
	rootCmd.AddCommand(collectRewards)
	collectRewards.PersistentFlags().
		String("pool_id", "",
			"stake pool id")
	collectRewards.PersistentFlags().
		String("provider_type", "blobber",
			"provider type")
	collectRewards.PersistentFlags().String("provider_id", "", "blobber or validator id")
	collectRewards.MarkFlagRequired("provider_id")
	collectRewards.MarkFlagRequired("pool_id")
	collectRewards.MarkFlagRequired("provider_type")

}
