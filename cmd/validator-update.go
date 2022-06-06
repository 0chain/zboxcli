package cmd

import (
	"fmt"
	"log"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"

	"github.com/spf13/cobra"
)

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
			validator.StakePoolSettings.MinStake = common.ToBalance(minStake)
		}

		if flags.Changed("max_stake") {
			var maxStake float64
			if maxStake, err = flags.GetFloat64("max_stake"); err != nil {
				log.Fatal(err)
			}
			validator.StakePoolSettings.MaxStake = common.ToBalance(maxStake)
		}

		if flags.Changed("num_delegates") {
			var nd int
			if nd, err = flags.GetInt("num_delegates"); err != nil {
				log.Fatal(err)
			}
			validator.StakePoolSettings.NumDelegates = nd
		}

		if flags.Changed("service_charge") {
			var sc float64
			if sc, err = flags.GetFloat64("service_charge"); err != nil {
				log.Fatal(err)
			}
			validator.StakePoolSettings.ServiceCharge = sc
		}

		if _, _, err = sdk.UpdateValidatorSettings(validator); err != nil {
			log.Fatal(err)
		}

		fmt.Println("validator settings updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(validatorUpdateCmd)

	buf := validatorUpdateCmd.Flags()
	buf.String("validator_id", "", "validator ID, required")
	buf.Float64("min_stake", 0.0, "update min_stake, optional")
	buf.Float64("max_stake", 0.0, "update max_stake, optional")
	buf.Int("num_delegates", 0, "update num_delegates, optional")
	buf.Float64("service_charge", 0.0, "update service_charge, optional")
	validatorUpdateCmd.MarkFlagRequired("validator_id")
}
