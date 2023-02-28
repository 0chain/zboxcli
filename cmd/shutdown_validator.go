package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var shutDownValidatorCmd = &cobra.Command{
	Use:   "shutdown-validator",
	Short: "deactivate a validator",
	Long:  "deactivate a validator, it will not be used for any new challenge validations",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		_, _, err = sdk.ShutdownProvider(sdk.ProviderValidator)
		if err != nil {
			log.Fatal("failed to shut down validator", err)
		}
		log.Println("shut down validator.")

	},
}

func init() {
	rootCmd.AddCommand(shutDownValidatorCmd)
	shutDownValidatorCmd.PersistentFlags().String("id", "", "the blobber id which you want to shut down")
	shutDownValidatorCmd.Flags().Float64("fee", 0.0, "fee for transaction")
}
