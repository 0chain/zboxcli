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

		var addBlobberId, removeBlobberId string
		if flags.Changed("add_blobber") {
			addBlobberId, err = flags.GetString("add_blobber")
			if err != nil {
				log.Fatal("invalid 'add_blobber' flag: ", err)
			}
			if flags.Changed("remove_blobber") {
				removeBlobberId, err = flags.GetString("remove_blobber")
				if err != nil {
					log.Fatal("invalid 'remove_blobber' flag: ", err)
				}
			}
		} else {
			if flags.Changed("remove_blobber") {
				log.Fatal("Error: cannot remove blobber without adding one")
			}
		}

		if flags.Changed("free_storage") {
			lock, freeStorageMarker := processFreeStorageFlags(flags)
			if lock < 0 {
				log.Fatal("Only positive values are allowed for --lock")
			}

			txnHash, n, err := sdk.CreateFreeUpdateAllocation(freeStorageMarker, allocID, lock)
			if err != nil {
				log.Fatal("Error free update allocation: ", err)
			}
			log.Println("Allocation updated with txId : " + txnHash)
			n = n //log.Println("nonce:", n)
			return
		}

		var updateTerms = false
		if flags.Changed("update_terms") {
			updateTerms, err = flags.GetBool("update_terms")
			if err != nil {
				log.Fatal("invalid update terms entry: ", err)
			}
		}

		var lockf float64
		var lock int64
		if lockf, err = flags.GetFloat64("lock"); err != nil {
			log.Fatal("error: invalid 'lock' value:", err)
		}
		if lock < 0 {
			log.Fatal("Only positive values are allowed for --lock")
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

		setImmutable, _ := cmd.Flags().GetBool("set_immutable")

		var allocationName string
		if flags.Changed("name") {
			allocationName, err = flags.GetString("name")
			if err != nil {
				log.Fatal("invalid allocation name: ", err)
			}
		}

		txnHash, n, err := sdk.UpdateAllocation(
			allocationName,
			size,
			int64(expiry/time.Second),
			allocID,
			lock,
			setImmutable,
			updateTerms,
			addBlobberId,
			removeBlobberId,
		)
		if err != nil {
			log.Fatal("Error updating allocation:", err)
		}
		log.Println("Allocation updated with txId : " + txnHash)
		n = n //log.Println("nonce:", n)
	},
}

func init() {
	rootCmd.AddCommand(updateAllocationCmd)
	updateAllocationCmd.PersistentFlags().String("allocation", "",
		"Allocation ID")
	updateAllocationCmd.PersistentFlags().String("add_blobber", "",
		"ID of blobber to add to the allocation")
	updateAllocationCmd.PersistentFlags().String("remove_blobber", "",
		"ID of blobber to remove from the allocation")
	updateAllocationCmd.PersistentFlags().Float64("lock", 0.0,
		"lock write pool with given number of tokens, required")
	updateAllocationCmd.PersistentFlags().Int64("size", 0,
		"adjust allocation size, bytes")
	updateAllocationCmd.PersistentFlags().Duration("expiry", 0,
		"adjust storage expiration time, duration")
	updateAllocationCmd.Flags().Bool("set_immutable", false, "set the allocation's data to be immutable")
	updateAllocationCmd.Flags().String("free_storage", "",
		"json file containing marker for free storage")
	updateAllocationCmd.Flags().Bool("update_terms", false,
		"update blobber terms")

	updateAllocationCmd.MarkFlagRequired("allocation")

	updateAllocationCmd.Flags().String("name", "", "allocation name")

}
