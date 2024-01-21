package cmd

import (
	"log"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var shutDownValidatorCmd = &cobra.Command{
	Use:   "shutdown-validator",
	Short: "deactivate a validator",
	Long:  "deactivate a validator, it will not be used for any new challenge validations",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()
		if fflags.Changed("id") == false {
			PrintError("Error: validator ID should be specified")
			os.Exit(1)
		}
		validatorID := cmd.Flag("id").Value.String()
		_, _, err := sdk.ShutdownProvider(sdk.ProviderValidator, validatorID)
		if err != nil {
			log.Fatal("failed to shut down validator", err)
		}
		log.Println("shut down validator.")

	},
}

func init() {
	rootCmd.AddCommand(shutDownValidatorCmd)
	shutDownValidatorCmd.PersistentFlags().String("id", "", "the validator id which you want to shut down")
	shutDownValidatorCmd.Flags().Float64("fee", 0.0, "fee for transaction")
}
