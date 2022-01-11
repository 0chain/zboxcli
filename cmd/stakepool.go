package cmd

import (
	"fmt"
	"log"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

func printStakePoolInfo(info *sdk.StakePoolInfo) {
	fmt.Println("blobber_id:            ", info.ID)
	fmt.Println("total stake:           ", info.Balance)
	fmt.Println("going to unlock, total:", info.Unstake)

	fmt.Println("capacity:")
	fmt.Println("  free:       ", info.Free, "(for current write price)")
	fmt.Println("  capacity:   ", info.Capacity, "(blobber bid)")
	fmt.Println("  write_price:", info.WritePrice, "(blobber write price)")

	fmt.Println("  offers_total:", info.OffersTotal, "(total stake committed to offers)")
	fmt.Println("  unstake_total:", info.UnstakeTotal, "(total amount marked to unstake)")

	if len(info.Delegate) == 0 {
		fmt.Println("delegate_pools: no delegate pools")
	} else {
		fmt.Println("delegate_pools:")
		for _, dp := range info.Delegate {
			fmt.Println("- id:               ", dp.ID)
			fmt.Println("  balance:          ", dp.Balance)
			fmt.Println("  delegate_id:      ", dp.DelegateID)
			fmt.Println("  rewards:          ", dp.Rewards)
			fmt.Println("  penalty:          ", dp.Penalty)
			fmt.Println("  interests:        ", dp.Interests, "(payed)")
			fmt.Println("  pending_interests:", dp.PendingInterests, "(not payed yet, can be given by 'sp-pay-interests' command)")
			var gtu string
			if dp.Unstake {
				gtu = "<unstaking>"
			} else {
				gtu = "<not going>"
			}

			fmt.Println("  going to unlock at:", gtu)
		}
	}
	fmt.Println("penalty:  ", info.Penalty, "(total blobber penalty for all time)")

	fmt.Println("rewards:")
	fmt.Println("  charge:  ", info.Rewards.Charge, "(for all time)")
	fmt.Println("  blobber:  ", info.Rewards.Blobber, "(for all time)")
	fmt.Println("  validator:", info.Rewards.Validator, "(for all time)")

	// settings
	fmt.Println("settings:")
	fmt.Println("  delegate_wallet:", info.Settings.DelegateWallet)
	fmt.Println("  min_stake:      ", info.Settings.MinStake.String())
	fmt.Println("  max_stake:      ", info.Settings.MaxStake.String())
	fmt.Println("  num_delegates:  ", info.Settings.NumDelegates)
}

func printStakePoolUserInfo(info *sdk.StakePoolUserInfo) {
	if len(info.Pools) == 0 {
		fmt.Print("no delegate pools")
		return
	}
	for blobberID, dps := range info.Pools {
		fmt.Println("- blobber_id: ", blobberID)
		for _, dp := range dps {
			fmt.Println("  - id:               ", dp.ID)
			fmt.Println("    balance:          ", dp.Balance)
			fmt.Println("    delegate_id:      ", dp.DelegateID)
			fmt.Println("    rewards:          ", dp.Rewards)
			fmt.Println("    penalty:          ", dp.Penalty)
			fmt.Println("    interests:        ", dp.Interests, "(payed)")
			fmt.Println("    pending_interests:", dp.PendingInterests,
				"(not payed yet, can be given by 'sp-pay-interests' command)")
		}
	}
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

		doJSON, _ := cmd.Flags().GetBool("json")

		if flags.Changed("blobber_id") {
			if blobberID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("can't get 'blobber_id' flag: %v", err)
			}
		}

		var info *sdk.StakePoolInfo
		if info, err = sdk.GetStakePoolInfo(blobberID); err != nil {
			log.Fatalf("Failed to get stake pool info: %v", err)
		}
		if doJSON {
			util.PrintJSON(info)
		} else {
			printStakePoolInfo(info)
		}
	},
}

// spUserInfo information per user
var spUserInfo = &cobra.Command{
	Use:   "sp-user-info",
	Short: "Stake pool information for a user.",
	Long:  `Stake pool information for a user.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags    = cmd.Flags()
			clientID string
			err      error
		)

		doJSON, _ := cmd.Flags().GetBool("json")

		if flags.Changed("client_id") {
			if clientID, err = flags.GetString("client_id"); err != nil {
				log.Fatalf("can't get 'client_id' flag: %v", err)
			}
		}

		var info *sdk.StakePoolUserInfo
		if info, err = sdk.GetStakePoolUserInfo(clientID); err != nil {
			log.Fatalf("Failed to get stake pool info: %v", err)
		}
		if doJSON {
			util.PrintJSON(info)
		} else {
			printStakePoolUserInfo(info)
		}
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
			flags     = cmd.Flags()
			blobberID string
			tokens    float64
			fee       float64
			err       error
		)

		if flags.Changed("blobber_id") {
			if blobberID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("invalid 'blobber_id' flag: %v", err)
			}
		}

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

		var poolID string
		poolID, err = sdk.StakePoolLock(blobberID,
			zcncore.ConvertToValue(tokens), zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to lock tokens in stake pool: %v", err)
		}
		fmt.Println("tokens locked, pool id:", poolID)
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
			flags             = cmd.Flags()
			blobberID, poolID string
			fee               float64
			err               error
		)

		if flags.Changed("blobber_id") {
			if blobberID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("invalid 'blobber_id' flag: %v", err)
			}
		}

		if !flags.Changed("pool_id") {
			log.Fatal("missing required 'pool_id' flag")
		}

		if poolID, err = flags.GetString("pool_id"); err != nil {
			log.Fatal("invalid 'pool_id' flag: ", err)
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		var unstake common.Timestamp
		unstake, err = sdk.StakePoolUnlock(blobberID, poolID,
			zcncore.ConvertToValue(fee))

		// an error
		if err != nil {
			log.Fatalf("Failed to unlock tokens in stake pool: %v", err)
		}

		// can't unlock for now
		if unstake > 0 {
			fmt.Println("tokens can't be unlocked due to opened offers")
			fmt.Printf("the pool marked as releasing, wait %s and retry to succeed", unstake.ToTime())
			fmt.Println()
			return
		}

		// success
		fmt.Println("tokens has unlocked, pool deleted")
	},
}

// spPayInterests pays interests not payed yet. A stake pool changes
// pays all interests can be payed. But if stake pool is not changed,
// then user can manually pay the interests.
var spPayInterests = &cobra.Command{
	Use:   "sp-pay-interests",
	Short: "Pay interests not payed yet.",
	Long:  `Pay interests not payed.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags     = cmd.Flags()
			blobberID string
			err       error
		)

		if flags.Changed("blobber_id") {
			if blobberID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("invalid 'blobber_id' flag: %v", err)
			}
		}

		if err = sdk.StakePoolPayInterests(blobberID); err != nil {
			log.Fatalf("Failed to pay interests: %v", err)
		}
		fmt.Println("interests has payed")
	},
}

func init() {
	rootCmd.AddCommand(spInfo)
	rootCmd.AddCommand(spUserInfo)
	rootCmd.AddCommand(spLock)
	rootCmd.AddCommand(spUnlock)
	rootCmd.AddCommand(spPayInterests)

	spInfo.PersistentFlags().String("blobber_id", "",
		"for given blobber, default is current client")
	spInfo.PersistentFlags().Bool("json", false, "pass this option to print response as json data")

	spUserInfo.PersistentFlags().Bool("json", false, "pass this option to print response as json data")

	spLock.PersistentFlags().String("blobber_id", "",
		"for given blobber, default is current client")
	spLock.PersistentFlags().Float64("tokens", 0.0,
		"tokens to lock, required")
	spLock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
	spLock.MarkFlagRequired("tokens")
	spLock.MarkFlagRequired("blobber_id")

	spUnlock.PersistentFlags().String("blobber_id", "",
		"for given blobber, default is current client")
	spUnlock.PersistentFlags().String("pool_id", "",
		"pool id to unlock")
	spUnlock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
	spUnlock.MarkFlagRequired("tokens")
	spUnlock.MarkFlagRequired("pool_id")

	spPayInterests.PersistentFlags().String("blobber_id", "",
		"for given blobber, default is current client")

}
