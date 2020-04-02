package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// finiAllocationCmd used to change allocation size and expiration
var finiAllocationCmd = &cobra.Command{
	Use:   "alloc-fini",
	Short: "Finalize an expired allocation",
	Long: `Finalize an expired allocation by allocation owner or one of
blobbers of the allocation. It moves all tokens have to be moved between pools
and empties write pool moving left tokens to client.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var flags = cmd.Flags()

		if flags.Changed("allocation") == false {
			log.Fatal("Error: allocation flag is missing")
		}

		allocID, err := flags.GetString("allocation")
		if err != nil {
			log.Fatal("invalid 'allocation' flag: ", err)
		}

		txnHash, err := sdk.FinalizeAllocation(allocID)
		if err != nil {
			log.Fatal("Error creating allocation:", err)
		}
		log.Print("Allocation finalized with txId : " + txnHash)
	},
}

// cancelAllocationCmd used to cancel allocation where blobbers
// doesn't provides their service in reality
var cancelAllocationCmd = &cobra.Command{
	Use:   "alloc-cancel",
	Short: "Cancel an allocation",
	Long: `Cancel allocation used to terminate an allocation where, because
of blobbers, it can't be used. Thus, the blobbers will not receive their
min_lock_demand. Other aspects of the cancellation follows the finalize
allocation flow.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var flags = cmd.Flags()

		if flags.Changed("allocation") == false {
			log.Fatal("Error: allocation flag is missing")
		}

		allocID, err := flags.GetString("allocation")
		if err != nil {
			log.Fatal("invalid 'allocation' flag: ", err)
		}

		txnHash, err := sdk.CancelAlloctioan(allocID)
		if err != nil {
			log.Fatal("Error creating allocation:", err)
		}
		log.Print("Allocation canceled with txId : " + txnHash)
	},
}

func init() {
	rootCmd.AddCommand(finiAllocationCmd)
	rootCmd.AddCommand(cancelAllocationCmd)

	finiAllocationCmd.PersistentFlags().String("allocation", "",
		"Allocation ID")
	finiAllocationCmd.MarkFlagRequired("allocation")

	cancelAllocationCmd.PersistentFlags().String("allocation", "",
		"Allocation ID")
	cancelAllocationCmd.MarkFlagRequired("allocation")
}
