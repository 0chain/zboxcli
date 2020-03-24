package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func printStakePoolInfo(info *sdk.StakePoolInfo) {
	var header = []string{
		"LOCKED", "UNLOCKED", "OFFERS TOTAL", "REQUIRED STAKE",
	}
	var data = [][]string{{
		fmt.Sprint(zcncore.ConvertToToken(info.Locked)),
		fmt.Sprint(zcncore.ConvertToToken(info.Unlocked)),
		fmt.Sprint(zcncore.ConvertToToken(info.OffersTotal)),
		fmt.Sprint(zcncore.ConvertToToken(info.RequiredStake)),
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
			fmt.Sprint(zcncore.ConvertToToken(val.Lock)),
			time.Unix(int64(val.Expire), 0).String(),
			val.AllocationID,
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
	rootCmd.AddCommand(spUnlock)

	spInfo.PersistentFlags().String("blobber_id", "",
		"for given blobber, default is current client")

	spUnlock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
}
