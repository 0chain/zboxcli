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

			txnHash, _, err := sdk.CreateFreeUpdateAllocation(freeStorageMarker, allocID, lock)
			if err != nil {
				log.Fatal("Error free update allocation: ", err)
			}
			log.Println("Allocation updated with txId : " + txnHash)
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
		var lock uint64
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

		setThirdPartyExtendable, _ := cmd.Flags().GetBool("set_third_party_extendable")

		// Read the file options flags
		var fileOptionParams sdk.FileOptionsParameters
		if flags.Changed("forbid_upload") {
			forbidUpload, err := flags.GetBool("forbid_upload")
			if err != nil {
				log.Fatal("invalid forbid_upload: ", err)
			}
			fileOptionParams.ForbidUpload.Changed = true
			fileOptionParams.ForbidUpload.Value = forbidUpload
		}
		if flags.Changed("forbid_delete") {
			forbidDelete, err := flags.GetBool("forbid_delete")
			if err != nil {
				log.Fatal("invalid forbid_upload: ", err)
			}
			fileOptionParams.ForbidDelete.Changed = true
			fileOptionParams.ForbidDelete.Value = forbidDelete
		}
		if flags.Changed("forbid_update") {
			forbidUpdate, err := flags.GetBool("forbid_update")
			if err != nil {
				log.Fatal("invalid forbid_upload: ", err)
			}
			fileOptionParams.ForbidUpdate.Changed = true
			fileOptionParams.ForbidUpdate.Value = forbidUpdate
		}
		if flags.Changed("forbid_move") {
			forbidMove, err := flags.GetBool("forbid_move")
			if err != nil {
				log.Fatal("invalid forbid_upload: ", err)
			}
			fileOptionParams.ForbidMove.Changed = true
			fileOptionParams.ForbidMove.Value = forbidMove
		}
		if flags.Changed("forbid_copy") {
			forbidCopy, err := flags.GetBool("forbid_copy")
			if err != nil {
				log.Fatal("invalid forbid_upload: ", err)
			}
			fileOptionParams.ForbidCopy.Changed = true
			fileOptionParams.ForbidCopy.Value = forbidCopy
		}
		if flags.Changed("forbid_rename") {
			forbidRename, err := flags.GetBool("forbid_rename")
			if err != nil {
				log.Fatal("invalid forbid_upload: ", err)
			}
			fileOptionParams.ForbidRename.Changed = true
			fileOptionParams.ForbidRename.Value = forbidRename
		}

		txnHash, _, err := sdk.UpdateAllocation(
			size,
			int64(expiry/time.Second),
			allocID,
			lock,
			updateTerms,
			addBlobberId,
			removeBlobberId,
			setThirdPartyExtendable,
			&fileOptionParams,
		)
		if err != nil {
			log.Fatal("Error updating allocation:", err)
		}
		log.Println("Allocation updated with txId : " + txnHash)
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
	updateAllocationCmd.Flags().String("free_storage", "",
		"json file containing marker for free storage")
	updateAllocationCmd.Flags().Bool("update_terms (boolean)", false,
		"update blobber terms")

	updateAllocationCmd.MarkFlagRequired("allocation")

	updateAllocationCmd.Flags().String("name", "", "allocation name")

	updateAllocationCmd.Flags().Bool("set_third_party_extendable (boolean)", false, "specify if the allocation can be extended by users other than the owner")
	updateAllocationCmd.Flags().Bool("forbid_upload (boolean)", false, "specify if users cannot upload to this allocation")
	updateAllocationCmd.Flags().Bool("forbid_delete (boolean)", false, "specify if the users cannot delete objects from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_update (boolean)", false, "specify if the users cannot update objects in this allocation")
	updateAllocationCmd.Flags().Bool("forbid_move (boolean)", false, "specify if the users cannot move objects from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_copy (boolean)", false, "specify if the users cannot copy object from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_rename (boolean)", false, "specify if the users cannot rename objects in this allocation")

}
