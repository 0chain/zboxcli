package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func printStakePoolInfo(info *sdk.StakePoolInfo) {
	var header = []string{
		"LOCKED", "OFFERS TOTAL", "CAP. STAKE", "LACK", "OVERFILL", "REWARD",
	}
	var data = [][]string{{
		info.Locked.String(),
		info.OffersTotal.String(),
		info.CapacityStake.String(),
		info.Lack.String(),
		info.Overfill.String(),
		info.Reward.String(),
	}}
	fmt.Println("POOL ID:", info.ID)
	util.WriteTable(os.Stdout, header, []string{}, data)
	fmt.Println()
}

func printStakePoolOffers(offers []*sdk.StakePoolOfferStat) {
	if len(offers) == 0 {
		fmt.Println("NO OFFERS")
		return
	}
	fmt.Println("OFFERS:")
	var header = []string{
		"LOCK", "EXPIRE", "ALLOC.", "EXPIRED",
	}
	var data = make([][]string, len(offers))
	for i, val := range offers {
		data[i] = []string{
			val.Lock.String(),
			val.Expire.ToTime().String(),
			string(val.AllocationID),
			fmt.Sprint(val.IsExpired),
		}
	}
	util.WriteTable(os.Stdout, header, []string{}, data)
	fmt.Println()
}

// spInfo information
var spInfo = &cobra.Command{
	Use:   "sp-info",
	Short: "Stake pool information.",
	Long:  `Stake pool information.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags     = cmd.Flags()
			blobberID string
			err       error
		)

		if flags.Changed("blobber_id") {
			if blobberID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("can't get 'blobber_id' flag: %v", err)
			}
		}

		var info *sdk.StakePoolInfo
		if info, err = sdk.GetStakePoolInfo(blobberID); err != nil {
			log.Fatalf("Failed to get stake pool info: %v", err)
		}
		printStakePoolInfo(info)
		printStakePoolOffers(info.Offers)
	},
}

// spLock locks tokens a stake pool lack
var spLock = &cobra.Command{
	Use:   "sp-lock",
	Short: "Lock tokens lacking in stake pool.",
	Long:  `Lock tokens lacking in stake pool.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags  = cmd.Flags()
			tokens float64
			fee    float64
			err    error
		)

		if !flags.Changed("tokens") {
			log.Fatal("missing required 'tokens' flag")
		}

		if tokens, err = flags.GetFloat64("tokens"); err != nil {
			log.Fatal("invalid 'tokens' flag: ", err)
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		err = sdk.StakePoolLock(zcncore.ConvertToValue(tokens),
			zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to unlock tokens in stake pool: %v", err)
		}
		fmt.Println("unlocked")
	},
}

// spUnlock unlocks tokens in stake pool
var spUnlock = &cobra.Command{
	Use:   "sp-unlock",
	Short: "Unlock tokens in stake pool.",
	Long:  `Unlock tokens in stake pool.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags = cmd.Flags()
			fee   float64
			err   error
		)

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		err = sdk.StakePoolUnlock(zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to unlock tokens in stake pool: %v", err)
		}
		fmt.Println("unlocked")
	},
}

func init() {
	rootCmd.AddCommand(spInfo)
	rootCmd.AddCommand(spLock)
	rootCmd.AddCommand(spUnlock)

	spInfo.PersistentFlags().String("blobber_id", "",
		"for given blobber, default is current client")

	spLock.PersistentFlags().Float64("tokens", 0.0,
		"tokens to lock, required")
	spLock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")

	spUnlock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
}
