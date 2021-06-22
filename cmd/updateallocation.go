package cmd

import (
	"log"
	"time"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

// updateAllocationCmd used to change allocation size and expiration
var updateAllocationCmd = &cobra.Command{
	Use:   "updateallocation",
	Short: "Updates allocation's expiry and size",
	Long:  `Updates allocation's expiry and size`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var flags = cmd.Flags()
		if flags.Changed("allocation") == false {
			log.Fatal("Error: allocation flag is missing")
		}

		allocID, err := flags.GetString("allocation")
		if err != nil {
			log.Fatal("invalid 'allocation_id' flag: ", err)
		}

		if flags.Changed("free_storage") {
			lock, freeStorageMarker := processFreeStorageFlags(flags)

			txnHash, err := sdk.CreateFreeUpdateAllocation(freeStorageMarker, allocID, lock)
			if err != nil {
				log.Fatal("Error free update allocation: ", err)
			}
			log.Print("Allocation updated with txId : " + txnHash)
			return
		}

		var lockf float64
		var lock int64
		if lockf, err = flags.GetFloat64("lock"); err != nil {
			log.Fatal("error: invalid 'lock' value:", err)
		}
		lock = zcncore.ConvertToValue(lockf)

		size, err := flags.GetInt64("size")
		if err != nil {
			log.Fatal("invalid 'size' flag: ", err)
		}

		expiry, err := flags.GetDuration("expiry")
		if err != nil {
			log.Fatal("invalid 'expiry' flag: ", err)
		}

		txnHash, err := sdk.UpdateAllocation(size,
			int64(expiry/time.Second), allocID, lock)
		if err != nil {
			log.Fatal("Error creating allocation:", err)
		}
		log.Print("Allocation updated with txId : " + txnHash)
	},
}

func init() {
	rootCmd.AddCommand(updateAllocationCmd)
	updateAllocationCmd.PersistentFlags().String("allocation", "",
		"Allocation ID")
	updateAllocationCmd.PersistentFlags().Float64("lock", 0.0,
		"lock write pool with given number of tokens, required")
	updateAllocationCmd.PersistentFlags().Int64("size", 0,
		"adjust allocation size, bytes")
	updateAllocationCmd.PersistentFlags().Duration("expiry", 0,
		"adjust storage expiration time, duration")
	updateAllocationCmd.Flags().
		String("free_storage", "",
			"json file containing marker for free storage")
	updateAllocationCmd.MarkFlagRequired("allocation")

}
