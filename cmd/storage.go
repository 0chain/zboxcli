package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

func printBlobbers(nodes []*sdk.Blobber) {
	if len(nodes) == 0 {
		fmt.Println("no blobbers registered yet")
		return
	}
	for _, val := range nodes {
		fmt.Println("- id:                   ", val.ID)
		fmt.Println("  url:                  ", val.BaseURL)
		fmt.Println("  used / total capacity:", val.Used.String(), "/",
			val.Capacity.String())
		fmt.Println("  last_health_check:	 ", val.LastHealthCheck)
		fmt.Println("  terms:")
		fmt.Println("    read_price:         ", val.Terms.ReadPrice.String(), "tok / GB")
		fmt.Println("    write_price:        ", val.Terms.WritePrice.String(), "tok / GB / time_unit")
		fmt.Println("    min_lock_demand:    ", val.Terms.MinLockDemand)
		fmt.Println("    cct:                ", val.Terms.ChallengeCompletionTime.String())
		fmt.Println("    max_offer_duration: ", val.Terms.MaxOfferDuration.String())
	}
}

func filterActiveBlobbers(blobbers []*sdk.Blobber) (activeBlobbers []*sdk.Blobber) {
	for i := range blobbers {
		if blobbers[i].LastHealthCheck.Within(60 * 60) {
			activeBlobbers = append(activeBlobbers, blobbers[i])
		}
	}
	return
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

		var list, err = sdk.GetBlobbers()
		var defaultList = filterActiveBlobbers(list)

		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
		}

		if doAll {
			defaultList = list
		}

		if doJSON {
			util.PrintJSON(defaultList)
		} else {
			printBlobbers(defaultList)
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
		fmt.Println("last_health_check:", blob.LastHealthCheck.ToTime())
		fmt.Println("capacity_used:    ", blob.Used)
		fmt.Println("terms:")
		fmt.Println("  read_price:        ", blob.Terms.ReadPrice, "tok / GB")
		fmt.Println("  write_price:       ", blob.Terms.WritePrice, "tok / GB")
		fmt.Println("  min_lock_demand:   ", blob.Terms.MinLockDemand*100.0, "%")
		fmt.Println("  max_offer_duration:", blob.Terms.MaxOfferDuration)
		fmt.Println("  cct:               ", blob.Terms.ChallengeCompletionTime)
		fmt.Println("settings:")
		fmt.Println("  delegate_wallet:", blob.StakePoolSettings.DelegateWallet)
		fmt.Println("  min_stake:      ", blob.StakePoolSettings.MinStake, "tok")
		fmt.Println("  max_stake:      ", blob.StakePoolSettings.MaxStake, "tok")
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

		var blob *sdk.Blobber
		if blob, err = sdk.GetBlobber(blobberID); err != nil {
			log.Fatal(err)
		}

		if flags.Changed("capacity") {
			var capacity int64
			if capacity, err = flags.GetInt64("capacity"); err != nil {
				log.Fatal(err)
			}
			blob.Capacity = common.Size(capacity)
		}

		if flags.Changed("read_price") {
			var rp float64
			if rp, err = flags.GetFloat64("read_price"); err != nil {
				log.Fatal(err)
			}
			blob.Terms.ReadPrice = common.ToBalance(rp)
		}

		if flags.Changed("write_price") {
			var wp float64
			if wp, err = flags.GetFloat64("write_price"); err != nil {
				log.Fatal(err)
			}
			blob.Terms.WritePrice = common.ToBalance(wp)
		}

		if flags.Changed("min_lock_demand") {
			var mld float64
			if mld, err = flags.GetFloat64("min_lock_demand"); err != nil {
				log.Fatal(err)
			}
			if mld < 0 || mld > 1 {
				log.Fatal("invalid min_lock_demand: out of [0; 1) range")
			}
			blob.Terms.MinLockDemand = mld
		}

		if flags.Changed("max_offer_duration") {
			var mod time.Duration
			if mod, err = flags.GetDuration("max_offer_duration"); err != nil {
				log.Fatal(err)
			}
			blob.Terms.MaxOfferDuration = mod
		}

		if flags.Changed("cct") {
			var cct time.Duration
			if cct, err = flags.GetDuration("cct"); err != nil {
				log.Fatal(err)
			}
			blob.Terms.ChallengeCompletionTime = cct
		}

		if flags.Changed("min_stake") {
			var minStake float64
			if minStake, err = flags.GetFloat64("min_stake"); err != nil {
				log.Fatal(err)
			}
			blob.StakePoolSettings.MinStake = common.ToBalance(minStake)
		}

		if flags.Changed("max_stake") {
			var maxStake float64
			if maxStake, err = flags.GetFloat64("max_stake"); err != nil {
				log.Fatal(err)
			}
			blob.StakePoolSettings.MaxStake = common.ToBalance(maxStake)
		}

		if flags.Changed("num_delegates") {
			var nd int
			if nd, err = flags.GetInt("num_delegates"); err != nil {
				log.Fatal(err)
			}
			blob.StakePoolSettings.NumDelegates = nd
		}

		if flags.Changed("service_charge") {
			var sc float64
			if sc, err = flags.GetFloat64("service_charge"); err != nil {
				log.Fatal(err)
			}
			blob.StakePoolSettings.ServiceCharge = sc
		}

		if _, err = sdk.UpdateBlobberSettings(blob); err != nil {
			log.Fatal(err)
		}

		fmt.Println("blobber settings updated successfully")

	},
}

func init() {
	rootCmd.AddCommand(lsBlobers)
	rootCmd.AddCommand(blobberInfoCmd)
	rootCmd.AddCommand(blobberUpdateCmd)

	lsBlobers.Flags().Bool("json", false, "pass this option to print response as json data")
	lsBlobers.Flags().Bool("all", false, "shows active and non active list of blobbers on ls-blobbers")

	blobberInfoCmd.Flags().String("blobber_id", "", "blobber ID, required")
	blobberInfoCmd.Flags().Bool("json", false,
		"pass this option to print response as json data")
	blobberInfoCmd.MarkFlagRequired("blobber_id")

	buf := blobberUpdateCmd.Flags()
	buf.String("blobber_id", "", "blobber ID, required")
	buf.Int64("capacity", 0, "update blobber capacity bid, optional")
	buf.Float64("read_price", 0.0, "update read_price, optional")
	buf.Float64("write_price", 0.0, "update write_price, optional")
	buf.Float64("min_lock_demand", 0.0, "update min_lock_demand, optional")
	buf.Duration("max_offer_duration", 0*time.Second, "update max_offer_duration, optional")
	buf.Duration("cct", 0*time.Second, "update challenge completion time (cct), optional")
	buf.Float64("min_stake", 0.0, "update min_stake, optional")
	buf.Float64("max_stake", 0.0, "update max_stake, optional")
	buf.Int("num_delegates", 0, "update num_delegates, optional")
	buf.Float64("service_charge", 0.0, "update service_charge, optional")
	blobberUpdateCmd.MarkFlagRequired("blobber_id")
}
