package cmd

import (
	"log"
	"os"
	"sync"

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

		extend, err := flags.GetBool("extend")
		if err != nil {
			log.Fatal("invalid 'extend' flag: ", err)
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

		if addBlobberId != "" {
			allocationObj, err := sdk.GetAllocation(allocID)
			if err != nil {
				log.Fatal("Error updating allocation:couldnt_find_allocation: Couldn't find the allocation required for update")
			}

			wg := &sync.WaitGroup{}
			statusBar := &StatusBar{wg: wg}
			wg.Add(1)
			allocUnderRepair = true
			if txnHash, err := allocationObj.UpdateWithRepair(
				size,
				extend,
				lock,
				addBlobberId,
				removeBlobberId,
				setThirdPartyExtendable,
				&fileOptionParams,
				statusBar,
			); err != nil {
				allocUnderRepair = false
				log.Fatal("Error updating allocation:", err)
			} else {
				log.Println("Allocation updated with txId : " + txnHash)
			}
			wg.Wait()
			if !statusBar.success {
				os.Exit(1)
			}
		} else {
			txnHash, _, err := sdk.UpdateAllocation(
				size,
				extend,
				allocID,
				lock,
				addBlobberId,
				removeBlobberId,
				setThirdPartyExtendable,
				&fileOptionParams,
			)
			if err != nil {
				log.Fatal("Error updating allocation:", err)
			}
			log.Println("Allocation updated with txId : " + txnHash)
		}
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
		"Increase (positive value) or decrease (negative value) the allocation size by a given value, in bytes")
	updateAllocationCmd.PersistentFlags().Bool("extend", false,
		"(default false) adjust storage expiration time, duration")

	updateAllocationCmd.MarkFlagRequired("allocation")

	updateAllocationCmd.Flags().String("name", "", "allocation name")

	updateAllocationCmd.Flags().Bool("set_third_party_extendable", false, "(default false) specify if the allocation can be extended by users other than the owner")
	updateAllocationCmd.Flags().Bool("forbid_upload", false, "(default false) specify if users cannot upload to this allocation")
	updateAllocationCmd.Flags().Bool("forbid_delete", false, "(default false) specify if the users cannot delete objects from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_update", false, "(default false) specify if the users cannot update objects in this allocation")
	updateAllocationCmd.Flags().Bool("forbid_move", false, "(default false) specify if the users cannot move objects from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_copy", false, "(default false) specify if the users cannot copy object from this allocation")
	updateAllocationCmd.Flags().Bool("forbid_rename", false, "(default false) specify if the users cannot rename objects in this allocation")

}
