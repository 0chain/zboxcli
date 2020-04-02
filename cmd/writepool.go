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

func printWritePoolInfo(info *sdk.WritePoolInfo) {
	var header = []string{
		"BALANCE", "START", "EXPIRE", "FINIALIZED",
	}
	var data = [][]string{{
		info.Balance.String(),
		info.StartTime.ToTime().String(),
		info.Expiration.ToTime().String(),
		fmt.Sprint(info.Finalized),
	}}
	fmt.Println("POOL ID:", info.ID)
	util.WriteTable(os.Stdout, header, []string{}, data)
	fmt.Println()
}

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

		if !flags.Changed("allocation") {
			log.Fatal("missing required 'allocation' flag")
		}

		if allocID, err = flags.GetString("allocation"); err != nil {
			log.Fatalf("can't get 'allocation' flag: %v", err)
		}

		var info *sdk.WritePoolInfo
		if info, err = sdk.GetWritePoolInfo(allocID); err != nil {
			log.Fatalf("Failed to get stake pool info: %v", err)
		}
		printWritePoolInfo(info)
	},
}

// wpLock locks additional tokens to a write pool of an allocation
var wpLock = &cobra.Command{
	Use:   "wp-lock",
	Short: "Lock tokens to write pool.",
	Long:  `Lock additional tokens to a write pool of an allocation.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags   = cmd.Flags()
			allocID string
			tokens  float64
			fee     float64
			err     error
		)

		if !flags.Changed("allocation") {
			log.Fatal("missing required 'allocation' flag")
		}

		if !flags.Changed("tokens") {
			log.Fatal("missing required 'tokens' flag")
		}

		allocID, err = flags.GetString("allocation")
		if err != nil {
			log.Fatal("invalid 'allocation' flag: ", err)
		}

		tokens, err = flags.GetFloat64("tokens")
		if err != nil {
			log.Fatal("invalid 'tokens' flag: ", err)
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		err = sdk.WritePoolLock(allocID, zcncore.ConvertToValue(tokens),
			zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatalf("Failed to lock tokens to write pool: %v", err)
		}
		fmt.Println("locked")
	},
}

func init() {
	rootCmd.AddCommand(wpInfo)
	rootCmd.AddCommand(wpLock)

	wpInfo.PersistentFlags().String("allocation", "",
		"allocation identifier, required")

	wpLock.PersistentFlags().String("allocation", "",
		"allocation identifier, required")
	wpLock.PersistentFlags().Float64("tokens", 0.0,
		"tokens to lock, required")
	wpLock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
}
