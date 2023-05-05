package cmd

import (
	"fmt"
	"log"

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
		if _, _, err = sdk.CreateReadPool(); err != nil {
			log.Fatalf("Failed to create read pool: %v\n", err)
		}
		fmt.Println("Read pool created successfully")
	},
}

// rpInfo information
var rpInfo = &cobra.Command{
	Use:   "rp-info",
	Short: "Read pool information.",
	Long:  `Read pool information.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")

		info, err := sdk.GetReadPoolInfo("")
		if err != nil {
			log.Fatalf("Failed to get read pool info: %v", err)
		}

		token, err := info.Balance.ToToken()
		if err != nil {
			log.Fatal(err)
		}
		usd, err := zcncore.ConvertTokenToUSD(token)
		var bt = float64(info.Balance) / 1e10
		if err != nil {
			log.Fatalf("Failed to convert token to usd: %v", err)
		}

		if info.Balance == 0 {
			fmt.Println("no tokens locked")
			return
		}

		if doJSON {
			jsonCurrencies := map[string]interface{}{
				"usd": usd,
				"zcn": bt,
				"fmt": info.Balance,
			}

			util.PrintJSON(jsonCurrencies)
			return
		}
		fmt.Printf("\nRead pool Balance: %v (%.2f USD)\n", info.Balance, usd)
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

		if tokens < 0 {
			log.Fatal("invalid token amount: negative")
		}

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, _, err = sdk.ReadPoolLock(zcncore.ConvertToValue(tokens), zcncore.ConvertToValue(fee))
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
			flags = cmd.Flags()
			fee   float64
			err   error
		)

		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, _, err = sdk.ReadPoolUnlock(zcncore.ConvertToValue(fee))
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

	rpInfo.Flags().Bool("json", false, "<type:bool> pass this option to print response as json data")

	rpLock.PersistentFlags().Float64("tokens", 0.0,
		"lock tokens number, required")
	rpLock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
	rpLock.MarkFlagRequired("tokens")

	rpUnlock.PersistentFlags().Float64("fee", 0.0,
		"transaction fee, default 0")
}
