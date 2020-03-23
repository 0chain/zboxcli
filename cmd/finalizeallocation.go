package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// finiAllocationCmd used to change allocation size and expiration
var finiAllocationCmd = &cobra.Command{
	Use:   "finialloc",
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
		log.Print("Allocation updated with txId : " + txnHash)
	},
}

func init() {
	rootCmd.AddCommand(finiAllocationCmd)
	finiAllocationCmd.PersistentFlags().String("allocation", "",
		"Allocation ID")
	finiAllocationCmd.MarkFlagRequired("allocation")
}
