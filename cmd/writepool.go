package cmd

import (
	"fmt"
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

// wpLock locks tokens in write pool
var wpLock = &cobra.Command{
	Use:   "wp-lock",
	Short: "Lock some tokens in write pool.",
	Long:  `Lock some tokens in write pool.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags   = cmd.Flags()
			allocID string // required
			tokens  float64
			fee     float64
			err     error
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

		if allocID, err = flags.GetString("allocation"); err != nil {
			log.Fatal("invalid 'allocation' flag: ", err)
		}

		if tokens, err = flags.GetFloat64("tokens"); err != nil {
			log.Fatal("invalid 'tokens' flag: ", err)
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, _, err = sdk.WritePoolLock(allocID, zcncore.ConvertToValue(tokens), zcncore.ConvertToValue(fee))
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
			flags   = cmd.Flags()
			allocID string
			fee     float64
			err     error
		)

		if !flags.Changed("allocation") {
			log.Fatal("missing required 'pool_id' flag")
		}

		if allocID, err = flags.GetString("pool_id"); err != nil {
			log.Fatal("invalid 'pool_id' flag: ", err)
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, _, err = sdk.WritePoolUnlock(allocID, zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to unlock tokens in write pool: %v", err)
		}
		fmt.Println("unlocked")
	},
}

func init() {
	rootCmd.AddCommand(wpLock)
	rootCmd.AddCommand(wpUnlock)

	wpLock.PersistentFlags().String("allocation", "",
		"allocation id to lock for, required")
	wpLock.PersistentFlags().Float64("tokens", 0.0,
		"lock tokens number, required")
	wpLock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")

	wpLock.MarkFlagRequired("allocation")
	wpLock.MarkFlagRequired("tokens")

	wpUnlock.PersistentFlags().String("allocation", "",
		"allocation id from which to unlock tokens")
	wpUnlock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")

	wpUnlock.MarkFlagRequired("allocation")
}
