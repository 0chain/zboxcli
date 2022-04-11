package cmd

import (
	"log"
	"sync"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"

	"github.com/spf13/cobra"
)

// The updateAttributesCmd used to update file attributes.
var updateAttributesCmd = &cobra.Command{
	Use:   "update-attributes",
	Short: "update object attributes on blobbers",
	Long:  `update object attributes on blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			fflags = cmd.Flags()

			allocID    string
			remotePath string

			alloc           *sdk.Allocation
			changed, commit bool

			err error
		)

		if !fflags.Changed("allocation") {
			log.Fatal("missing 'allocation' flag")
		}

		if !fflags.Changed("remotepath") {
			log.Fatal("missing 'remotepath' flag")
		}

		if allocID, err = fflags.GetString("allocation"); err != nil {
			log.Fatalf("parsing 'allocation' flag: %v", err)
		}

		if remotePath, err = fflags.GetString("remotepath"); err != nil {
			log.Fatalf("paring 'remotepath' flag: %v", err)
		}

		if fflags.Changed("commit") {
			if commit, err = fflags.GetBool("commit"); err != nil {
				log.Fatalf("parsing 'commit' flag: %v'", err)
			}
		}

		if alloc, err = storageSdk.GetAllocation(allocID); err != nil {
			log.Fatal("fetching the allocation: ", err)
		}

		meta, err := alloc.GetFileMeta(remotePath)
		if err != nil {
			log.Fatal("fetching the metadata: ", err)
		}

		var attrs = meta.Attributes

		// modify attributes by the flags
		if fflags.Changed("who-pays-for-reads") {
			var (
				wp  common.WhoPays
				wps string
			)
			if wps, err = fflags.GetString("who-pays-for-reads"); err != nil {
				log.Fatalf("getting 'who-pays-for-reads' flag: %v", err)
			}
			if err = wp.Parse(wps); err != nil {
				log.Fatal(err)
			}
			if wp != attrs.WhoPaysForReads {
				attrs.WhoPaysForReads, changed = wp, true // change
			}
		}

		if !changed {
			log.Print("no changes")
			return
		}

		if err = alloc.UpdateObjectAttributes(remotePath, attrs); err != nil {
			log.Fatal("updating file attributes: ", err)
		}

		log.Print("attributes updated")

		if !commit {
			return
		}

		log.Print("committing changes...")
		var (
			wg        = &sync.WaitGroup{}
			statusBar = &StatusBar{wg: wg}
		)
		wg.Add(1)
		commitMetaTxn(remotePath, "Update attributes", "", "", alloc, meta,
			statusBar)
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(updateAttributesCmd)
	updateAttributesCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	updateAttributesCmd.PersistentFlags().String("remotepath", "", "Remote path of object to rename")
	updateAttributesCmd.PersistentFlags().String("who-pays-for-reads", "owner",
		"Who pays for reads: owner or 3rd_party")
	updateAttributesCmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	updateAttributesCmd.MarkFlagRequired("allocation")
	updateAttributesCmd.MarkFlagRequired("remotepath")
}
