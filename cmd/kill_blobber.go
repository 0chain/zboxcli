package cmd

import (
	"log"

	"github.com/0chain/gosdk/zcncore"

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

		var fee float64
		if flags.Changed("fee") {
			if fee, err = flags.GetFloat64("fee"); err != nil {
				log.Fatal("invalid 'fee' flag: ", err)
			}
		}

		_, err = sdk.KillBlobber(blobberId, zcncore.ConvertToValue(fee))
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
