package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

type Provider int

const (
	ProviderMiner Provider = iota
	ProviderSharder
	ProviderBlobber
	ProviderValidator
	ProviderAuthorizer
)

var collectRewards = &cobra.Command{
	Use:   "collect_reward",
	Short: "Collect accrued rewards for a stake pool.",
	Long:  "Collect accrued rewards for a stake pool.",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		if !flags.Changed("pool_id") {
			log.Fatal("missing pool id flag")
		}

		poolId, err := flags.GetString("pool_id")
		if err != nil {
			log.Fatal(err)
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
			_, err = sdk.CollectRewards(poolId, int(ProviderBlobber))
		case "validator":
			_, err = sdk.CollectRewards(poolId, int(ProviderValidator))
		default:
			log.Fatal("provider type must be blobber or validator")
		}
		if err != nil {
			log.Fatal("Error paying reward:", err)
		}
		log.Print("transferred reward tokens")
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

	collectRewards.MarkFlagRequired("pool_id")
	collectRewards.MarkFlagRequired("provider_type")

}
