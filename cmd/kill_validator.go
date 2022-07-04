package cmd

import (
	"log"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

var killValidatorCmd = &cobra.Command{
	Use:   "kill-validator",
	Short: "punitively deactivate a validator",
	Long:  "punitively deactivate a validator",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		flags := cmd.Flags()
		if flags.Changed("id") == false {
			PrintError("validator id is missing")
			os.Exit(1)
		}
		validatorId, err := flags.GetString("id")
		if err != nil {
			log.Fatal("invalid 'allocation' flag: ", err)
		}

		var fee float64
		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, err = sdk.KillBlobber(validatorId, zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatal("failed to kill validator, id: "+validatorId, err)
		}
		log.Println("killed validator, id: " + validatorId)

	},
}

func init() {
	rootCmd.AddCommand(killValidatorCmd)
	killValidatorCmd.PersistentFlags().String("id", "", "validator's id")
	_ = killValidatorCmd.MarkFlagRequired("id")
}
