package cmd

import (
	"github.com/spf13/cobra"
	"log"
)
import "github.com/0chain/gosdk/zboxcore/sdk"

var addCmd = &cobra.Command{
	Use:   "newallocation",
	Short: "Creates a new allocation",
	Long:  `Creates a new allocation`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var flags = cmd.Flags()

		name, err := flags.GetString("name")
		if err != nil {
			log.Fatal("invalid 'name' flag: ", err)
		}
		key, err := flags.GetString("key")
		if err != nil {
			log.Fatal("invalid 'key' flag: ", err)
		}
		limit, err := flags.GetFloat64("limit")
		if err != nil {
			log.Fatal("invalid 'limit' flag: ", err)
		}
		max, err := flags.GetString("max")
		if err != nil {
			log.Fatal("invalid 'max' flag: ", err)
		}

		err = sdk.AddFreeStorageAssigner(name, key, limit, max)
		if err != nil {
			log.Fatal("Error adding free storage assigner:", err)
		}
		log.Print(name + " added as free storage assigner")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	newallocationCmd.PersistentFlags().
		Float64("limit", 1.0,
			"maximum number of tokens to use in a free allocation")
	newallocationCmd.PersistentFlags().
		Float64("max", 1.0,
			"the total number of tokens that can be given in free allocations")
	newallocationCmd.Flags().
		String("giver_id", "",
			"the account number that will be creating free storage markers")
	newallocationCmd.Flags().
		String("key", "",
			"the public key used for singing markers")

	updateAllocationCmd.MarkFlagRequired("name")
	updateAllocationCmd.MarkFlagRequired("limit")
	updateAllocationCmd.MarkFlagRequired("max")
	updateAllocationCmd.MarkFlagRequired("key")
}
