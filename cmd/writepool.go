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

// wpInfo information
var wpInfo = &cobra.Command{
	Use:   "wp-info",
	Short: "Write pool information.",
	Long:  `Write pool information.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags   = cmd.Flags()
			allocID string
			err     error
		)

		if flags.Changed("allocation") {
			if allocID, err = flags.GetString("allocation"); err != nil {
				log.Fatalf("can't get 'allocation' flag: %v", err)
			}
		}
		doJSON, _ := cmd.Flags().GetBool("json")

		var info *sdk.AllocationPoolStats
		if info, err = sdk.GetWritePoolInfo(""); err != nil {
			log.Fatalf("Failed to get write pool info: %v", err)
		}
		if len(info.Pools) == 0 {
			fmt.Println("no tokens locked")
			return
		}

		info.AllocFilter(allocID)
		if doJSON {
			util.PrintJSON(info.Pools)
			return
		}
		printReadPoolStat(info.Pools)
	},
}

// wpLock locks tokens in write pool
var wpLock = &cobra.Command{
	Use:   "wp-lock",
	Short: "Lock some tokens in write pool.",
	Long:  `Lock some tokens in write pool.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags     = cmd.Flags()
			duration  time.Duration
			allocID   string // required
			blobberID string // optional
			tokens    float64
			fee       float64
			err       error
		)

		if !flags.Changed("duration") {
			log.Fatal("missing required 'duration' flag")
		}

		if !flags.Changed("allocation") {
			log.Fatal("missing required 'allocation' flag")
		}

		if !flags.Changed("tokens") {
			log.Fatal("missing required 'tokens' flag")
		}

		if duration, err = flags.GetDuration("duration"); err != nil {
			log.Fatal("invalid 'duration' flag: ", err)
		}

		if allocID, err = flags.GetString("allocation"); err != nil {
			log.Fatal("invalid 'allocation' flag: ", err)
		}

		if flags.Changed("blobber") {
			if blobberID, err = flags.GetString("blobber"); err != nil {
				log.Fatal("invalid 'blobber' flag: ", err)
			}
		}

		if tokens, err = flags.GetFloat64("tokens"); err != nil {
			log.Fatal("invalid 'tokens' flag: ", err)
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		err = sdk.WritePoolLock(duration, allocID, blobberID,
			zcncore.ConvertToValue(tokens), zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to lock tokens in write pool: %v", err)
		}
		fmt.Println("locked")
	},
}

// wpUnlock unlocks tokens in a write pool
var wpUnlock = &cobra.Command{
	Use:   "wp-unlock",
	Short: "Unlock some expired tokens in a write pool.",
	Long:  `Unlock some expired tokens in a write pool.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags  = cmd.Flags()
			poolID string
			fee    float64
			err    error
		)

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

		err = sdk.WritePoolUnlock(poolID, zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to unlock tokens in write pool: %v", err)
		}
		fmt.Println("unlocked")
	},
}

func init() {
	rootCmd.AddCommand(wpInfo)
	rootCmd.AddCommand(wpLock)
	rootCmd.AddCommand(wpUnlock)

	wpInfo.PersistentFlags().String("allocation", "",
		"allocation, optional")
	wpInfo.Flags().Bool("json", false, "pass this option to print response as json data")

	wpLock.PersistentFlags().Duration("duration", 0,
		"lock duration, required")
	wpLock.PersistentFlags().String("allocation", "",
		"allocation id to lock for, required")
	wpLock.PersistentFlags().String("blobber", "",
		"blobber id to lock for, optional")
	wpLock.PersistentFlags().Float64("tokens", 0.0,
		"lock tokens number, required")
	wpLock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")

	wpLock.MarkFlagRequired("duration")
	wpLock.MarkFlagRequired("allocation")
	wpLock.MarkFlagRequired("tokens")

	wpUnlock.PersistentFlags().String("pool_id", "",
		"expired write pool identifier, required")
	wpUnlock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")

	wpUnlock.MarkFlagRequired("pool_id")
}
