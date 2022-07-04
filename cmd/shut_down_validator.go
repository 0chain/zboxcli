package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

var shutDownValidatorCmd = &cobra.Command{
	Use:   "shut-down-validator",
	Short: "deactivate a validator",
	Long:  "deactivate a validator, it will not be used for any new challenge validations",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		flags := cmd.Flags()

		var fee float64
		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, err = sdk.ShutDownValidator(zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatal("failed to shut down validator", err)
		}
		log.Println("shut down validator")

	},
}

func init() {
	rootCmd.AddCommand(shutDownValidatorCmd)
}
