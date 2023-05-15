package cmd

import (
	"fmt"
	"log"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

func printValidators(nodes []*sdk.Validator) {
	if len(nodes) == 0 {
		fmt.Println("no validators registered yet")
		return
	}
	for _, validator := range nodes {
		fmt.Println("id:               ", validator.ID)
		fmt.Println("url:              ", validator.BaseURL)
		fmt.Println("last_health_check: ", validator.LastHealthCheck.ToTime())
		fmt.Println("is killed:        ", validator.IsKilled)
		fmt.Println("is shut down:     ", validator.IsShutdown)
		fmt.Println("settings:")
		fmt.Println("  delegate_wallet:", validator.DelegateWallet)
		fmt.Println("  min_stake:      ", validator.MinStake)
		fmt.Println("  max_stake:      ", validator.MaxStake)
		fmt.Println("  total_stake:    ", validator.StakeTotal)
		fmt.Println("  total_unstake:  ", validator.UnstakeTotal)
		fmt.Println("  num_delegates:  ", validator.NumDelegates)
		fmt.Println("  service_charge: ", validator.ServiceCharge*100, "%")
	}
}

// lsBlobers shows active blobbers
var lsValidators = &cobra.Command{
	Use:   "ls-validators",
	Short: "Show active Validators.",
	Long:  `Show active Validators in the network.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")

		var list, err = sdk.GetValidators()

		if err != nil {
			log.Fatalf("Failed to get storage SC configurations: %v", err)
		}

		if doJSON {
			util.PrintJSON(list)
		} else {
			printValidators(list)
		}
	},
}

var validatorInfoCmd = &cobra.Command{
	Use:   "validator-info",
	Short: "Get validator info",
	Long:  `Get validator info`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags = cmd.Flags()

			json        bool
			validatorID string
			err         error
		)

		if flags.Changed("json") {
			if json, err = flags.GetBool("json"); err != nil {
				log.Fatal("invalid 'json' flag: ", err)
			}
		}

		if !flags.Changed("validator_id") {
			log.Fatal("missing required 'validator_id' flag")
		}

		if validatorID, err = flags.GetString("validator_id"); err != nil {
			log.Fatal("error in 'validator_id' flag: ", err)
		}

		var validator *sdk.Validator
		if validator, err = sdk.GetValidator(validatorID); err != nil {
			log.Fatal(err)
		}

		if json {
			util.PrintJSON(validator)
		} else {
			printValidators([]*sdk.Validator{validator})
		}

	},
}

var validatorUpdateCmd = &cobra.Command{
	Use:   "validator-update",
	Short: "Update validator settings by its delegate_wallet owner",
	Long:  `Update validator settings by its delegate_wallet owner`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			flags = cmd.Flags()

			validatorID string
			err         error
		)

		if !flags.Changed("validator_id") {
			log.Fatal("missing required 'validator_id' flag")
		}

		if validatorID, err = flags.GetString("validator_id"); err != nil {
			log.Fatal("error in 'validator_id' flag: ", err)
		}

		var validator *sdk.Validator
		if validator, err = sdk.GetValidator(validatorID); err != nil {
			log.Fatal(err)
		}

		if flags.Changed("min_stake") {
			var minStake float64
			if minStake, err = flags.GetFloat64("min_stake"); err != nil {
				log.Fatal(err)
			}
			stake, err := common.ToBalance(minStake)
			if err != nil {
				log.Fatal(err)
			}
			validator.MinStake = stake
		}

		if flags.Changed("max_stake") {
			var maxStake float64
			if maxStake, err = flags.GetFloat64("max_stake"); err != nil {
				log.Fatal(err)
			}
			stake, err := common.ToBalance(maxStake)
			if err != nil {
				log.Fatal(err)
			}
			validator.MaxStake = stake
		}

		if flags.Changed("num_delegates") {
			var nd int
			if nd, err = flags.GetInt("num_delegates"); err != nil {
				log.Fatal(err)
			}
			validator.NumDelegates = nd
		}

		if flags.Changed("service_charge") {
			var sc float64
			if sc, err = flags.GetFloat64("service_charge"); err != nil {
				log.Fatal(err)
			}
			validator.ServiceCharge = sc
		}

		if _, _, err = sdk.UpdateValidatorSettings(validator); err != nil {
			log.Fatal(err)
		}

		fmt.Println("validator settings updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(validatorUpdateCmd)
	rootCmd.AddCommand(validatorInfoCmd)
	rootCmd.AddCommand(lsValidators)

	validatorInfoCmd.Flags().String("validator_id", "", "validator ID, required")
	validatorInfoCmd.Flags().Bool("json", false,
		"(default false) pass this option to print response as json data")
	validatorInfoCmd.MarkFlagRequired("validator_id")

	buf := validatorUpdateCmd.Flags()
	buf.String("validator_id", "", "validator ID, required")
	buf.Float64("min_stake", 0.0, "update min_stake, optional")
	buf.Float64("max_stake", 0.0, "update max_stake, optional")
	buf.Int("num_delegates", 0, "update num_delegates, optional")
	buf.Float64("service_charge", 0.0, "update service_charge, optional")
	validatorUpdateCmd.MarkFlagRequired("validator_id")

	lsValidators.Flags().Bool("json", false, "(default false) pass this flag to get response as json object")
}
