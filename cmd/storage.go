package cmd

import (
	"fmt"
	"github.com/0chain/common/core/currency"
	"log"
	"time"

	"github.com/0chain/gosdk/zboxcore/blockchain"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

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
		util.PrintJSON(conf.Fields)
	},
}

func printBlobbers(nodes []*sdk.Blobber, isActive bool) {
	if len(nodes) == 0 {
		if isActive {
			fmt.Println("no active blobbers")
		} else {
			fmt.Println("no blobbers registered yet")
		}
		return
	}
	for _, val := range nodes {
		fmt.Println("- id:                   ", val.ID)
		fmt.Println("  url:                  ", val.BaseURL)
		fmt.Println("  allocated / total capacity:", val.Allocated.String(), "/",
			val.Capacity.String())
		fmt.Println("  last_health_check:	 ", val.LastHealthCheck.ToTime())
		fmt.Println("  terms:")
		fmt.Println("    read_price:         ", val.Terms.ReadPrice.String(), "/ GB")
		fmt.Println("    write_price:        ", val.Terms.WritePrice.String(), "/ GB / time_unit")
		fmt.Println("    max_offer_duration: ", val.Terms.MaxOfferDuration.String())
	}
}

// lsBlobers shows active blobbers
var lsBlobers = &cobra.Command{
	Use:   "ls-blobbers",
	Short: "Show active blobbers in storage SC.",
	Long:  `Show active blobbers in storage SC.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")
		doAll, _ := cmd.Flags().GetBool("all")
		isStakable, err := cmd.Flags().GetBool("stakable")
		if err != nil {
			log.Fatalf("err parsing in stakable flag: %v", err)
		}
		// set is_active=true to get only active blobbers
		isActive := true
		if doAll {
			isActive = false
		}
		list, err := sdk.GetBlobbers(isActive, isStakable)
		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
		}

		if doJSON {
			util.PrintJSON(list)
		} else {
			printBlobbers(list, isActive)
		}
	},
}

var blobberInfoCmd = &cobra.Command{
	Use:   "bl-info",
	Short: "Get blobber info",
	Long:  `Get blobber info`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags = cmd.Flags()

			json      bool
			blobberID string
			err       error
		)

		if flags.Changed("json") {
			if json, err = flags.GetBool("json"); err != nil {
				log.Fatal("invalid 'json' flag: ", err)
			}
		}

		if !flags.Changed("blobber_id") {
			log.Fatal("missing required 'blobber_id' flag")
		}

		if blobberID, err = flags.GetString("blobber_id"); err != nil {
			log.Fatal("error in 'blobber_id' flag: ", err)
		}

		var blob *sdk.Blobber
		if blob, err = sdk.GetBlobber(blobberID); err != nil {
			log.Fatal(err)
		}

		if json {
			util.PrintJSON(blob)
			return
		}

		fmt.Println("id:               ", blob.ID)
		fmt.Println("url:              ", blob.BaseURL)
		fmt.Println("capacity:         ", blob.Capacity)
		fmt.Println("is killed:        ", blob.IsKilled)
		fmt.Println("is shut down:     ", blob.IsShutdown)
		fmt.Println("last_health_check:", blob.LastHealthCheck.ToTime())
		fmt.Println("capacity_used:    ", blob.Allocated)
		fmt.Println("total_stake:      ", blob.TotalStake)
		fmt.Println("not_available:     ", blob.NotAvailable)
		fmt.Println("is_restricted:     ", blob.IsRestricted)
		fmt.Println("terms:")
		fmt.Println("  read_price:        ", blob.Terms.ReadPrice, "/ GB")
		fmt.Println("  write_price:       ", blob.Terms.WritePrice, "/ GB")
		fmt.Println("  max_offer_duration:", blob.Terms.MaxOfferDuration)
		fmt.Println("settings:")
		fmt.Println("  delegate_wallet:", blob.StakePoolSettings.DelegateWallet)
		//fmt.Println("  min_stake:      ", blob.StakePoolSettings.MinStake)
		//fmt.Println("  max_stake:      ", blob.StakePoolSettings.MaxStake)
		fmt.Println("  num_delegates:  ", blob.StakePoolSettings.NumDelegates)
		fmt.Println("  service_charge: ", blob.StakePoolSettings.ServiceCharge*100, "%")
	},
}

var blobberUpdateCmd = &cobra.Command{
	Use:   "bl-update",
	Short: "Update blobber settings by its delegate_wallet owner",
	Long:  `Update blobber settings by its delegate_wallet owner`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			flags = cmd.Flags()

			blobberID string
			err       error
		)

		if !flags.Changed("blobber_id") {
			log.Fatal("missing required 'blobber_id' flag")
		}

		if blobberID, err = flags.GetString("blobber_id"); err != nil {
			log.Fatal("error in 'blobber_id' flag: ", err)
		}

		if _, err = sdk.GetBlobber(blobberID); err != nil {
			log.Fatal(err)
		}

		updateBlobber := new(sdk.UpdateBlobber)
		updateBlobber.ID = common.Key(blobberID)
		if flags.Changed("capacity") {
			var capacity int64
			if capacity, err = flags.GetInt64("capacity"); err != nil {
				log.Fatal(err)
			}

			changedCapacity := common.Size(capacity)
			updateBlobber.Capacity = &changedCapacity
		}

		terms := &sdk.UpdateTerms{}
		var termsChanged bool
		if flags.Changed("read_price") {
			var rp float64
			if rp, err = flags.GetFloat64("read_price"); err != nil {
				log.Fatal(err)
			}
			readPriceBalance, err := common.ToBalance(rp)
			if err != nil {
				log.Fatal(err)
			}
			terms.ReadPrice = &readPriceBalance
			termsChanged = true
		}

		if flags.Changed("write_price") {
			var wp float64
			if wp, err = flags.GetFloat64("write_price"); err != nil {
				log.Fatal(err)
			}
			writePriceBalance, err := common.ToBalance(wp)
			if err != nil {
				log.Fatal(err)
			}
			terms.WritePrice = &writePriceBalance
			termsChanged = true
		}

		if flags.Changed("max_offer_duration") {
			var mod time.Duration
			if mod, err = flags.GetDuration("max_offer_duration"); err != nil {
				log.Fatal(err)
			}
			terms.MaxOfferDuration = &mod
		}

		stakePoolSettings := &blockchain.UpdateStakePoolSettings{}
		var stakePoolSettingChanged bool
		if flags.Changed("num_delegates") {
			var nd int
			if nd, err = flags.GetInt("num_delegates"); err != nil {
				log.Fatal(err)
			}
			stakePoolSettings.NumDelegates = &nd
			stakePoolSettingChanged = true
		}

		if flags.Changed("service_charge") {
			var sc float64
			if sc, err = flags.GetFloat64("service_charge"); err != nil {
				log.Fatal(err)
			}
			stakePoolSettings.ServiceCharge = &sc
			stakePoolSettingChanged = true
		}

		if flags.Changed("url") {
			var url string
			if url, err = flags.GetString("url"); err != nil {
				log.Fatal(err)
			}
			updateBlobber.BaseURL = &url
		}

		if flags.Changed("not_available") {
			var na bool
			if na, err = flags.GetBool("not_available"); err != nil {
				log.Fatal(err)
			}
			if !na {
				na = false
			}
			updateBlobber.NotAvailable = &na
		}

		if flags.Changed("is_restricted") {
			var ia bool
			// Check if the flag is set to true
			if ia, err = flags.GetBool("is_restricted"); err != nil {
				log.Fatal(err)
			}
			if !ia {
				ia = false
			}
			updateBlobber.IsRestricted = &ia
		}

		if termsChanged {
			updateBlobber.Terms = terms
		}

		if stakePoolSettingChanged {
			updateBlobber.StakePoolSettings = stakePoolSettings
		}

		if _, _, err = sdk.UpdateBlobberSettings(updateBlobber); err != nil {
			log.Fatal(err)
		}
		fmt.Println("blobber settings updated successfully")
	},
}

var resetBlobberStatsCmd = &cobra.Command{
	Use:    "reset-blobber-stats",
	Short:  "Reset blobber stats",
	Long:   `Reset blobber stats`,
	Args:   cobra.MinimumNArgs(0),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			flags = cmd.Flags()

			blobberID       string
			prevTotalOffers currency.Coin
			newTotalOffers  currency.Coin
			err             error
		)

		if !flags.Changed("blobber_id") {
			log.Fatal("missing required 'blobber_id' flag")
		}
		if blobberID, err = flags.GetString("blobber_id"); err != nil {
			log.Fatal("error in 'blobber_id' flag: ", err)
		}

		if !flags.Changed("prev_total_offers") {
			log.Fatal("missing required 'prev_total_offers' flag")
		}
		var prevTotalOffersInt int64
		if prevTotalOffersInt, err = flags.GetInt64("prev_total_offers"); err != nil {
			log.Fatal("error in 'prev_total_offers' flag: ", err)
		}
		prevTotalOffers = currency.Coin(prevTotalOffersInt)

		var newTotalOffersInt int64
		if !flags.Changed("new_total_offers") {
			log.Fatal("missing required 'new_total_offers' flag")
		}
		if newTotalOffersInt, err = flags.GetInt64("new_total_offers"); err != nil {
			log.Fatal("error in 'new_total_offers' flag: ", err)
		}
		newTotalOffers = currency.Coin(newTotalOffersInt)

		resetBlobberStatsDto := &sdk.ResetBlobberStatsDto{
			BlobberID:       blobberID,
			PrevTotalOffers: prevTotalOffers,
			NewTotalOffers:  newTotalOffers,
		}
		fmt.Println(*resetBlobberStatsDto)

		_, _, err = sdk.ResetBlobberStats(resetBlobberStatsDto)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("reset blobber stats successfully")
	},
}

func init() {
	rootCmd.AddCommand(scConfig)
	rootCmd.AddCommand(lsBlobers)
	rootCmd.AddCommand(blobberInfoCmd)
	rootCmd.AddCommand(blobberUpdateCmd)
	rootCmd.AddCommand(resetBlobberStatsCmd)

	scConfig.Flags().Bool("json", false, "(default false) pass this option to print response as json data")
	lsBlobers.Flags().Bool("json", false, "(default false) pass this option to print response as json data")
	lsBlobers.Flags().Bool("all", false, "(default false) shows active and non active list of blobbers on ls-blobbers")
	lsBlobers.Flags().Bool("stakable", false, "(default false) gets only stakable list of blobbers if set to true")

	blobberInfoCmd.Flags().String("blobber_id", "", "blobber ID, required")
	blobberInfoCmd.Flags().Bool("json", false,
		"(default false) pass this option to print response as json data")
	blobberInfoCmd.MarkFlagRequired("blobber_id")

	buf := blobberUpdateCmd.Flags()
	buf.String("blobber_id", "", "blobber ID, required")
	buf.Int64("capacity", 0, "update blobber capacity bid, optional")
	buf.Float64("read_price", 0.0, "update read_price, optional")
	buf.Float64("write_price", 0.0, "update write_price, optional")
	buf.Duration("max_offer_duration", 0*time.Second, "update max_offer_duration, optional")
	buf.Float64("min_stake", 0.0, "update min_stake, optional")
	buf.Float64("max_stake", 0.0, "update max_stake, optional")
	buf.Int("num_delegates", 0, "update num_delegates, optional")
	buf.Float64("service_charge", 0.0, "update service_charge, optional")
	buf.Bool("not_available", true, "(default false) set blobber's availability for new allocations")
	buf.Bool("is_restricted", true, "(default false) set is_restricted")
	buf.String("url", "", "update the url of the blobber, optional")
	blobberUpdateCmd.MarkFlagRequired("blobber_id")

	resetBlobberStatsCmd.Flags().String("blobber_id", "", "blobber_id is required")
	resetBlobberStatsCmd.Flags().Int64("prev_allocated", 0, "prev_allocated is required")
	resetBlobberStatsCmd.Flags().Int64("prev_saved_data", 0, "prev_saved_data is required")
	resetBlobberStatsCmd.Flags().Int64("new_allocated", 0, "new_allocated is required")
	resetBlobberStatsCmd.Flags().Int64("new_saved_data", 0, "new_saved_data is required")
	resetBlobberStatsCmd.Flags().Int64("prev_total_offers", 0, "prev_total_offers is required")
	resetBlobberStatsCmd.Flags().Int64("new_total_offers", 0, "new_total_offers is required")
	resetBlobberStatsCmd.MarkFlagRequired("blobber_id")
}
