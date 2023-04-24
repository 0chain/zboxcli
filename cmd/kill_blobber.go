package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"

	"github.com/spf13/cobra"
)

var killBlobberCmd = &cobra.Command{
	Use:   "kill-blobber",
	Short: "punitively deactivate a blobber",
	Long:  "punitively deactivate a blobber",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		flags := cmd.Flags()
		if flags.Changed("id") == false {
			log.Fatal("id is missing")
		}
		blobberId, err := flags.GetString("id")
		if err != nil {
			log.Fatal("invalid 'blobber id flag: ", err)
		}
		_, _, err = sdk.KillProvider(blobberId, sdk.ProviderBlobber)
		if err != nil {
			log.Fatal("failed to kill blobber "+blobberId, err)
		}
		log.Println("killed blobber " + blobberId)
	},
}

func init() {
	rootCmd.AddCommand(killBlobberCmd)
	killBlobberCmd.PersistentFlags().String("id", "", "blobber's id")
	_ = killBlobberCmd.MarkFlagRequired("id")
}
