package cmd

import (
	"fmt"
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get blobber or validator status",
	Long:  `Get blobber or validator status`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags = cmd.Flags()
			json  bool
			err   error
		)

		if flags.Changed("json") {
			if json, err = flags.GetBool("json"); err != nil {
				log.Fatal("invalid 'json' flag: ", err)
			}
		}

		if !flags.Changed("id") {
			log.Fatal("missing required 'id' flag")
		}

		var providerId string
		if providerId, err = flags.GetString("id"); err != nil {
			log.Fatal("error in 'id' flag: ", err)
		}

		var providerType string
		if flags.Changed("type") {
			if providerType, err = flags.GetString("type"); err != nil {
				log.Fatal("error in 'type' flag: ", err)
			}
		}

		var status *sdk.ProviderStatus
		if status, err = sdk.StorageGetProviderStatus(providerId, providerType); err != nil {
			log.Fatal(err)
		}

		if json {
			util.PrintJSON(status)
			return
		}

		fmt.Println("id:               ", providerId)
		fmt.Println("type:             ", providerType)
		fmt.Println("status:           ", status.Status.String())
		if len(status.Reason) > 0 {
			fmt.Println("reason:           ", status.Reason)
		}
	},
}

func init() {
	rootCmd.AddCommand(StatusCmd)
	StatusCmd.Flags().Bool("json", false, "pass this option to print response as json data")
	StatusCmd.Flags().String("id", "", "provider ID, required")
	StatusCmd.Flags().String("type", "blobber", "provider type")

}
