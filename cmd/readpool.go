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

// rpCreate creates read pool
var rpCreate = &cobra.Command{
	Use:   "rp-create",
	Short: "Create read pool if missing",
	Long:  `Create read pool in storage SC if the pool is missing.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if err = sdk.CreateReadPool(); err != nil {
			log.Fatalf("Failed to create read pool: %v\n", err)
		}
		fmt.Println("Read pool created successfully")
	},
}

func printReadPoolStat(stat []*sdk.ReadPoolStat) {
	var header = []string{
		"ID", "START TIME", "DUR", "TIME LEFT", "LOCKED", "BALANCE",
	}
	var data = make([][]string, len(stat))
	for i, val := range stat {
		data[i] = []string{
			string(val.ID),
			val.StartTime.ToTime().String(),
			val.Duration.String(),
			val.TimeLeft.String(),
			fmt.Sprint(val.Locked),
			val.Balance.String(),
		}
	}
	util.WriteTable(os.Stdout, header, []string{}, data)
	fmt.Println()
}

// rpInfo information
var rpInfo = &cobra.Command{
	Use:   "rp-info",
	Short: "Read pool information.",
	Long:  `Read pool information.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags    = cmd.Flags()
			clientID string
			err      error
		)
		if flags.Changed("client_id") {
			if clientID, err = flags.GetString("client_id"); err != nil {
				log.Fatalf("can't get 'client_id' flag: %v", err)
			}
		}

		var info *sdk.ReadPoolInfo
		if info, err = sdk.GetReadPoolInfo(clientID); err != nil {
			log.Fatalf("Failed to get read pool info: %v", err)
		}
		if len(info.Stats) == 0 {
			fmt.Println("no locked tokens in the read pool")
			return
		}
		printReadPoolStat(info.Stats)
	},
}

// rpLock locks tokens in read pool
var rpLock = &cobra.Command{
	Use:   "rp-lock",
	Short: "Lock some tokens in read pool.",
	Long:  `Lock some tokens in read pool.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags    = cmd.Flags()
			duration time.Duration
			tokens   float64
			fee      float64
			err      error
		)

		if !flags.Changed("duration") {
			log.Fatal("missing required 'duration' flag")
		}

		if !flags.Changed("tokens") {
			log.Fatal("missing required 'tokens' flag")
		}

		if duration, err = flags.GetDuration("duration"); err != nil {
			log.Fatal("invalid 'duration' flag: ", err)
		}

		if tokens, err = flags.GetFloat64("tokens"); err != nil {
			log.Fatal("invalid 'tokens' flag: ", err)
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		err = sdk.ReadPoolLock(duration, zcncore.ConvertToValue(tokens),
			zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to lock tokens in read pool: %v", err)
		}
		fmt.Println("locked")
	},
}

// rpUnlock unlocks tokens in a read pool
var rpUnlock = &cobra.Command{
	Use:   "rp-unlock",
	Short: "Unlock some expired tokens in a read pool.",
	Long:  `Unlock some expired tokens in a read pool.`,
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

		err = sdk.ReadPoolUnlock(poolID, zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to unlock tokens in read pool: %v", err)
		}
		fmt.Println("unlocked")
	},
}

func init() {
	rootCmd.AddCommand(rpCreate)
	rootCmd.AddCommand(rpInfo)
	rootCmd.AddCommand(rpLock)
	rootCmd.AddCommand(rpUnlock)

	rpInfo.PersistentFlags().String("client_id", "",
		"for given client, default is current client")

	rpLock.PersistentFlags().Duration("duration", 0,
		"lock duration, required")
	rpLock.PersistentFlags().Float64("tokens", 0.0,
		"lock tokens number, required")
	rpLock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")

	rpLock.MarkFlagRequired("duration")
	rpLock.MarkFlagRequired("tokens")

	rpUnlock.PersistentFlags().String("pool_id", "",
		"expired read pool identifier, required")
	rpUnlock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
}
