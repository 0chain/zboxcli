package cmd

import (
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

var addCuratorCmd = &cobra.Command{
	Use:   "addcurator",
	Short: "Adds a curator to an allocation",
	Long:  "Adds a curator to an allocation",
	Args:  cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var flags = cmd.Flags()

		if flags.Changed("allocation") == false {
			return util.LogFatalErr("Error: allocation flag is missing")
		}
		allocationID, err := flags.GetString("allocation")
		if err != nil {
			return util.LogFatalErrf("invalid 'allocation_id' flag: %s", err)
		}

		if flags.Changed("curator") == false {
			return util.LogFatalErr("Error: curator flag is missing")
		}
		curatorID, err := flags.GetString("curator")
		if err != nil {
			return util.LogFatalErrf("invalid 'curator_id' flag: %s", err)
		}

		_, err = sdk.AddCurator(curatorID, allocationID)
		if err != nil {
			return util.LogFatalErrf("Error adding curator: %s", err)
		}
		util.LogPrintf("%s added %s as curator to allocation %s", curatorID, curatorID, allocationID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCuratorCmd)
	addCuratorCmd.PersistentFlags().
		String("curator", "",
			"new curator to add to allocation")
	addCuratorCmd.PersistentFlags().
		String("allocation", "",
			"allocation that the curator is to be added")

	addCuratorCmd.MarkFlagRequired("curator")
	addCuratorCmd.MarkFlagRequired("allocation")
}
