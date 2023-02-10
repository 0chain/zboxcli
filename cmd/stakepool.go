package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

func printStakePoolInfo(info *sdk.StakePoolInfo) {
	fmt.Println("pool id:           ", info.ID)
	fmt.Println("balance:           ", info.Balance)
	fmt.Println("total stake:       ", info.StakeTotal)
	fmt.Println("total unstake:     ", info.UnstakeTotal, "(total stake not available for further commitments)")
	fmt.Println("unclaimed rewards: ", info.Rewards)
	if len(info.Delegate) == 0 {
		fmt.Println("delegate_pools: no delegate pools")
	} else {
		fmt.Println("delegate_pools:")
		for _, dp := range info.Delegate {
			fmt.Println("- id:               ", dp.ID)
			fmt.Println("  balance:          ", dp.Balance)
			fmt.Println("  delegate_id:      ", dp.DelegateID)
			fmt.Println("  unclaimed reward: ", dp.Rewards)
			fmt.Println("  total_reward:     ", dp.TotalReward)
			fmt.Println("  total_penalty:    ", dp.TotalPenalty)
			fmt.Println("  status:           ", dp.Status)
			fmt.Println("  round_created:    ", dp.RoundCreated)
			fmt.Println("  unstake:          ", dp.UnStake)
			fmt.Println("  staked_at:        ", time.Unix(0, dp.StakedAt*int64(time.Second)).String())
		}
	}
	// settings
	fmt.Println("settings:")
	fmt.Println("  delegate_wallet:  ", info.Settings.DelegateWallet)
	fmt.Println("  min_stake:        ", info.Settings.MinStake.String())
	fmt.Println("  max_stake:        ", info.Settings.MaxStake.String())
	fmt.Println("  num_delegates:    ", info.Settings.NumDelegates)
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
			fmt.Println("    unclaimed reward:       ", dp.Rewards)
			fmt.Println("    total rewards:          ", dp.TotalReward)
			fmt.Println("    total penalty:          ", dp.TotalPenalty)
			fmt.Println("    status:          ", dp.Status)
			fmt.Println("    round_created:   ", dp.RoundCreated)
			fmt.Println("    unstake:         ", dp.UnStake)
			fmt.Println("    staked_at:       ", dp.StakedAt.String())
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
			flags        = cmd.Flags()
			err          error
			providerID   string
			providerType sdk.ProviderType
		)

		doJSON, _ := cmd.Flags().GetBool("json")

		if flags.Changed("blobber_id") {
			if providerID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("Error: cannot get the value of blobber_id")
			} else {
				providerType = sdk.ProviderBlobber
			}
		} else if flags.Changed("validator_id") {
			if providerID, err = flags.GetString("validator_id"); err != nil {
				log.Fatalf("Error: cannot get the value of validator_id")
			} else {
				providerType = sdk.ProviderValidator
			}
		}

		if providerType == 0 || providerID == "" {
			log.Fatal("Error: missing flag: one of 'blobber_id' or 'validator_id' is required")
		}

		var info *sdk.StakePoolInfo
		if info, err = sdk.GetStakePoolInfo(providerType, providerID); err != nil {
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
			flags        = cmd.Flags()
			providerID   string
			providerType sdk.ProviderType
			tokens       float64
			fee          float64
			err          error
		)

		if flags.Changed("blobber_id") {
			if providerID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("invalid 'blobber_id' flag: %v", err)
			} else {
				providerType = sdk.ProviderBlobber
			}
		} else if flags.Changed("validator_id") {
			if providerID, err = flags.GetString("validator_id"); err != nil {
				log.Fatalf("invalid 'validator_id' flag: %v", err)
			} else {
				providerType = sdk.ProviderValidator
			}
		}

		if providerType == 0 || providerID == "" {
			log.Fatal("missing flag: one of 'blobber_id' or 'validator_id' is required")
		}

		if !flags.Changed("tokens") {
			log.Fatal("missing required 'tokens' flag")
		}

		if tokens, err = flags.GetFloat64("tokens"); err != nil {
			log.Fatal("invalid 'tokens' flag: ", err)
		}

		if tokens < 0 {
			log.Fatal("invalid token amount: negative")
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		var hash string
		hash, _, err = sdk.StakePoolLock(providerType, providerID,
			zcncore.ConvertToValue(tokens), zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to lock tokens in stake pool: %v", err)
		}
		fmt.Println("tokens locked, txn hash:", hash)
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
			flags        = cmd.Flags()
			providerID   string
			providerType sdk.ProviderType
			fee          float64
			err          error
		)

		if flags.Changed("blobber_id") {
			if providerID, err = flags.GetString("blobber_id"); err != nil {
				log.Fatalf("invalid 'blobber_id' flag: %v", err)
			} else {
				providerType = sdk.ProviderBlobber
			}
		} else if flags.Changed("validator_id") {
			if providerID, err = flags.GetString("validator_id"); err != nil {
				log.Fatalf("invalid 'validator_id' flag: %v", err)
			} else {
				providerType = sdk.ProviderValidator
			}
		}

		if providerType == 0 || providerID == "" {
			log.Fatal("missing flag: one of 'blobber_id' or 'validator_id' is required")
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, _, err = sdk.StakePoolUnlock(providerType, providerID, zcncore.ConvertToValue(fee))
		// an error
		if err != nil {
			log.Fatalf("Failed to unlock tokens in stake pool: %v", err)
		}

		// success
		fmt.Println("tokens unlocked, pool deleted")
	},
}

func init() {
	rootCmd.AddCommand(spInfo)
	rootCmd.AddCommand(spUserInfo)
	rootCmd.AddCommand(spLock)
	rootCmd.AddCommand(spUnlock)

	spInfo.PersistentFlags().String("blobber_id", "",
		"for given blobber")
	spInfo.PersistentFlags().String("validator_id", "",
		"for given validator")
	spInfo.PersistentFlags().Bool("json", false, "pass this option to print response as json data")

	spUserInfo.PersistentFlags().Bool("json", false, "pass this option to print response as json data")

	spLock.PersistentFlags().String("blobber_id", "", "for given blobber")
	spLock.PersistentFlags().String("validator_id", "", "for given validator")
	spLock.PersistentFlags().Float64("tokens", 0.0, "tokens to lock, required")
	spLock.PersistentFlags().Float64("fee", 0.0, "transaction fee, default 0")

	spLock.MarkFlagRequired("tokens")

	spUnlock.PersistentFlags().String("blobber_id", "", "for given blobber")
	spUnlock.PersistentFlags().String("validator_id", "", "for given validator")
	spUnlock.PersistentFlags().Float64("fee", 0.0, "transaction fee, default 0")
	spUnlock.MarkFlagRequired("tokens")
}
