package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var transferAllocationCmd = &cobra.Command{
	Use:   "transferallocation",
	Short: "Transfer an allocation between owners",
	Long:  "Transfer an allocation between owners, only a curator can transfer an allocation",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var flags = cmd.Flags()

		if flags.Changed("allocation") == false {
			log.Fatal("Error: curator flag is missing")
		}
		allocationId, err := flags.GetString("allocation")
		if err != nil {
			log.Fatal("invalid 'allocation_id' flag: ", err)
		}

		if flags.Changed("new_owner") == false {
			log.Fatal("Error: curator flag is missing")
		}
		newOwnerId, err := flags.GetString("new_owner")
		if err != nil {
			log.Fatal("invalid 'new_owner_id' flag: ", err)
		}

		if flags.Changed("new_owner_key") == false {
			log.Fatal("Error: curator flag is missing")
		}
		newOwnerPublicKey, err := flags.GetString("new_owner_key")
		if err != nil {
			log.Fatal("invalid 'new_owner_key' flag: ", err)
		}

		_, n, err := sdk.CuratorTransferAllocation(allocationId, newOwnerId, newOwnerPublicKey)
		if err != nil {
			log.Fatal("Error adding curator:", err)
		}
		log.Println("transferred ownership of allocation " + allocationId + " to " + newOwnerId)
		log.Println("Nonce:", n)
	},
}

func init() {
	rootCmd.AddCommand(transferAllocationCmd)
	transferAllocationCmd.PersistentFlags().
		String("allocation", "",
			"allocation which is to have its ownership changed")
	transferAllocationCmd.PersistentFlags().
		String("new_owner", "",
			"id of the new owner")
	transferAllocationCmd.PersistentFlags().
		String("new_owner_key", "",
			"the public key of the new owner")

	updateAllocationCmd.MarkFlagRequired("allocation")
	updateAllocationCmd.MarkFlagRequired("new_owner")
	updateAllocationCmd.MarkFlagRequired("new_owner_key")
}
