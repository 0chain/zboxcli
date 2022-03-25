package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var removeCuratorCmd = &cobra.Command{
	Use:   "removecurator",
	Short: "Removes a curator from an allocation",
	Long:  "Removes a curator from an allocation",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var flags = cmd.Flags()

		if flags.Changed("allocation") == false {
			log.Fatal("Error: allocation flag is missing")
		}
		allocationID, err := flags.GetString("allocation")
		if err != nil {
			log.Fatal("invalid 'allocation_id' flag: ", err)
		}

		if flags.Changed("curator") == false {
			log.Fatal("Error: curator flag is missing")
		}
		curatorID, err := flags.GetString("curator")
		if err != nil {
			log.Fatal("invalid 'curator_id' flag: ", err)
		}

		_, n, err := sdk.RemoveCurator(curatorID, allocationID)
		if err != nil {
			log.Fatal("Error adding curator:", err)
		}
		log.Println(curatorID + " removed " + curatorID + " as a curator to allocation " + allocationID)
		//log.Println("nonce:", n)
	},
}

func init() {
	rootCmd.AddCommand(removeCuratorCmd)
	removeCuratorCmd.PersistentFlags().
		String("curator", "",
			"the curator to remove from allocation")
	removeCuratorCmd.PersistentFlags().
		String("allocation", "",
			"allocation from which the curator is to be removed")

	removeCuratorCmd.MarkFlagRequired("curator")
	removeCuratorCmd.MarkFlagRequired("allocation")
}
