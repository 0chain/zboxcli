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

		setImmutable		, _ := cmd.Flags().GetBool("set_immutable")
		thirdPartyExtendable, _ := cmd.Flags().GetBool("third_party_extendable")
		forbidUpload        , _ := flags.GetBool("forbid_upload");
		forbidDelete        , _ := flags.GetBool("forbid_delete");
		forbidUpdate        , _ := flags.GetBool("forbid_update");
		forbidMove          , _ := flags.GetBool("forbid_move");
		forbidCopy	        , _ := flags.GetBool("forbid_copy");
		forbidRename        , _ := flags.GetBool("forbid_rename");
		allowUpload        	, _ := flags.GetBool("allow_upload");
		allowDelete        	, _ := flags.GetBool("allow_delete");
		allowUpdate        	, _ := flags.GetBool("allow_update");
		allowMove          	, _ := flags.GetBool("allow_move");
		allowCopy	        , _ := flags.GetBool("allow_copy");
		allowRename        	, _ := flags.GetBool("allow_rename");

		var allocationName string
		if flags.Changed("name") {
			allocationName, err = flags.GetString("name")
			if err != nil {
				log.Fatal("invalid allocation name: ", err)
			}
		}

		txnHash, _, err := sdk.UpdateAllocation(
			allocationName,
			size,
			int64(expiry/time.Second),
			allocID,
			lock,
			setImmutable,
			updateTerms,
			addBlobberId,
			removeBlobberId,
			thirdPartyExtendable,
			&sdk.FileOptionsParameters{
				ForbidUpload: forbidUpload,
				ForbidDelete: forbidDelete,
				ForbidUpdate: forbidUpdate,
				ForbidMove: forbidMove,
				ForbidCopy: forbidCopy,
				ForbidRename: forbidRename,
				AllowUpload: allowUpload,
				AllowDelete: allowDelete,
				AllowUpdate: allowUpdate,
				AllowMove: allowMove,
				AllowCopy: allowCopy,
				AllowRename: allowRename,
			},
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
	updateAllocationCmd.Flags().Bool("set_immutable", false, "set the allocation's data to be immutable")
	updateAllocationCmd.Flags().String("free_storage", "",
		"json file containing marker for free storage")
	updateAllocationCmd.Flags().Bool("update_terms", false,
		"update blobber terms")

	updateAllocationCmd.MarkFlagRequired("allocation")

	updateAllocationCmd.Flags().String("name", "", "allocation name")

	updateAllocationCmd.Flags().Bool("third_party_extendable", false, "specify if the allocation can be extended by users other than the owner")
	updateAllocationCmd.Flags().Bool("forbid_upload", false, "specify if users cannot upload to this allocation")
	updateAllocationCmd.Flags().Bool("forbid_delete", false, "specify if the users cannot delete objects from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_update", false, "specify if the users cannot update objects in this allocation")
	updateAllocationCmd.Flags().Bool("forbid_move", false, "specify if the users cannot move objects from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_copy", false, "specify if the users cannot copy object from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_rename", false, "specify if the users cannot rename objects in this allocation")
	updateAllocationCmd.Flags().Bool("allow_upload", false, "specify if users can upload to this allocation")
	updateAllocationCmd.Flags().Bool("allow_delete", false, "specify if the users can delete objects from this allocation")
	updateAllocationCmd.Flags().Bool("allow_update", false, "specify if the users can update objects in this allocation")
	updateAllocationCmd.Flags().Bool("allow_move", false, "specify if the users can move objects from this allocation")
	updateAllocationCmd.Flags().Bool("allow_copy", false, "specify if the users can copy object from this allocation")
	updateAllocationCmd.Flags().Bool("allow_rename", false, "specify if the users can rename objects in this allocation")

}
