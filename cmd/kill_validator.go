package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
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
			log.Fatal("validator id is missing")
		}
		validatorId, err := flags.GetString("id")
		if err != nil {
			log.Fatal("invalid 'validator id flag: ", err)
		}

		_, _, err = sdk.KillProvider(validatorId, sdk.ProviderValidator)
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
