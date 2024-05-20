package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// NOTE: This is for testing purpose only.

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback file to previous version",
	Long:  `rollback file to previous version`,
	Args:  cobra.MinimumNArgs(0),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()
		if !fflags.Changed("allocation") {
			log.Fatal("Error: allocation flag is missing")
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			log.Fatal("Error fetching the allocation.", err)
		}

		_, err = allocationObj.GetCurrentVersion()
		if err != nil {
			log.Fatal("Error rolling back", err)
		}
		log.Println("Rollback successful")
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
	rollbackCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	_ = rollbackCmd.MarkFlagRequired("allocation")
}
