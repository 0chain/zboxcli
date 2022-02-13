package cmd

import (
	"log"

	"github.com/0chain/zboxcli/util"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds free storage assigner",
	Long:  "Adds free storage assigner",
	Args:  cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var flags = cmd.Flags()

		name, err := flags.GetString("name")
		if err != nil {
			return util.LogFatalErrf("invalid 'name' flag: ", err)
		}
		key, err := flags.GetString("key")
		if err != nil {
			return util.LogFatalErrf("invalid 'name' flag: ", err)
		}
		limit, err := flags.GetFloat64("limit")
		if err != nil {
			return util.LogFatalErrf("invalid 'limit' flag: ", err)
		}
		max, err := flags.GetFloat64("max")
		if err != nil {
			return util.LogFatalErrf("invalid 'max' flag: ", err)
		}

		err = sdk.AddFreeStorageAssigner(name, key, limit, max)
		if err != nil {
			return util.LogFatalErrf("Error adding free storage assigner:", err)
		}
		log.Print(name + " added as free storage assigner")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().
		Float64("limit", 1.0,
			"maximum number of tokens to use in a free allocation")
	addCmd.PersistentFlags().
		Float64("max", 1.0,
			"the total number of tokens that can be given in free allocations")
	addCmd.Flags().
		String("name", "",
			"the account number that will be creating free storage markers")
	addCmd.Flags().
		String("key", "",
			"the public key used for singing markers")

	updateAllocationCmd.MarkFlagRequired("name")
	updateAllocationCmd.MarkFlagRequired("limit")
	updateAllocationCmd.MarkFlagRequired("max")
	updateAllocationCmd.MarkFlagRequired("key")
}
