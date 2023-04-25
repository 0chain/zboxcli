package cmd

import (
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// NOTE: This is for testing purpose only.

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback file to previous version",
	Long:  `rollback file to previous version`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()
		if !fflags.Changed("allocation") {
			PrintError("Error: allocation flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation.", err)
			os.Exit(1)
		}

		_, err = allocationObj.GetCurrentVersion()
		if err != nil {
			PrintError("Error rolling back", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
	uploadCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	uploadCmd.MarkFlagRequired("allocation")
}
