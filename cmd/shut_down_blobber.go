package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

var shutDownBlobberCmd = &cobra.Command{
	Use:   "shut-down-blobber",
	Short: "deactivate a blobber",
	Long:  "deactivate a blobber, it will not be used for any new allocations",
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

		_, err = sdk.ShutDownBlobber(zcncore.ConvertToValue(fee))
		if err != nil {
			log.Fatal("failed to shut down blobber", err)
		}
		log.Println("shut down blobber")

	},
}

func init() {
	rootCmd.AddCommand(shutDownBlobberCmd)
}
