package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var shutDownBlobberCmd = &cobra.Command{
	Use:   "shutdown-blobber",
	Short: "deactivate a blobber",
	Long:  "deactivate a blobber, it will not be used for any new allocations",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		_, _, err = sdk.ShutdownProvider(sdk.ProviderBlobber)
		if err != nil {
			log.Fatal("failed to shut down blobber", err)
		}
		log.Println("shut down blobber.")

	},
}

func init() {
	rootCmd.AddCommand(shutDownBlobberCmd)
	shutDownBlobberCmd.PersistentFlags().String("id", "", "the blobber id which you want to shut down")
	shutDownBlobberCmd.Flags().Float64("fee", 0.0, "fee for transaction")
}
