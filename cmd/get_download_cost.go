package cmd

import (
	"fmt"
	"log"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

func calculateDownloadCost(alloc *sdk.Allocation, fileSize int64, numBlocks int64) common.Balance {

	var cost float64

	for _, d := range alloc.BlobberDetails {
		readPrice, err := d.Terms.ReadPrice.ToToken()
		if err != nil {
			log.Fatalf("failed to convert %v to token, %v", d.Terms.ReadPrice, err)
		}

		cost += sizeInGB(numBlocks*fileref.CHUNK_SIZE) * float64(readPrice)

	}

	balance, err := common.ToBalance(cost)
	if err != nil {
		log.Fatalf("failed to convert %v to balance, %v", cost, err)
	}
	return balance
}

func downloadCost(alloc *sdk.Allocation, meta *sdk.ConsolidatedFileMeta, blocksPerMarker int) {
	if meta.Type != fileref.FILE {
		log.Fatal("not a file")
	}

	shardSize := (meta.ActualFileSize + int64(alloc.DataShards) - 1) / int64(alloc.DataShards)

	numBlocks := (shardSize + fileref.CHUNK_SIZE - 1) / fileref.CHUNK_SIZE

	requiredBalance := calculateDownloadCost(alloc, meta.ActualFileSize, numBlocks)

	fmt.Printf("%s tokens for %d 64KB blocks (%s) of %s\n", requiredBalance, numBlocks*int64(alloc.DataShards+alloc.ParityShards),
		common.Size(meta.Size), meta.Path)
}

// The getDownloadCostCmd returns value in tokens to download a file.
var getDownloadCostCmd = &cobra.Command{
	Use:   "get-download-cost",
	Short: "Get downloading cost",
	Long:  `Get downloading cost`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			fflags  = cmd.Flags()
			allocID string
			err     error
		)

		if !fflags.Changed("allocation") {
			log.Fatal("missing required 'allocation' flag")
		}

		allocID = cmd.Flag("allocation").Value.String()
		blocksPerMarker, err := cmd.Flags().GetInt("blocks-per-marker")

		if err != nil {
			log.Fatal("invalid blocks-per-marker value: ", err)
		}

		if blocksPerMarker <= 0 {
			log.Fatal("blocks-per-marker value cannot be <= 0")
		}

		var (
			remotePath string
			authTicket string
			lookupHash string
		)

		if fflags.Changed("remotepath") {
			if remotePath, err = fflags.GetString("remotepath"); err != nil {
				log.Fatal("invalid 'remotepath' flag: ", err)
			}
		}

		if fflags.Changed("authticket") {
			if authTicket, err = fflags.GetString("authticket"); err != nil {
				log.Fatal("invalid 'authticket' flag: ", err)
			}
		}

		if fflags.Changed("lookuphash") {
			if lookupHash, err = fflags.GetString("lookuphash"); err != nil {
				log.Fatal("invalid 'lookuphash' flag: ", err)
			}
		}

		var (
			alloc *sdk.Allocation
			meta  *sdk.ConsolidatedFileMeta
		)

		if remotePath != "" && allocID != ""{

			// by remote path

			if alloc, err = sdk.GetAllocation(allocID); err != nil {
				log.Fatal("fetching the allocation: ", err)
			}

			if meta, err = alloc.GetFileMeta(remotePath); err != nil {
				log.Fatal("can't get file meta: ", err)
			}

			downloadCost(alloc, meta, blocksPerMarker)
			return
		}

		// by authentication ticket

		//if authTicket is provided no need of anything else to get the cost
		if authTicket == "" {
			log.Fatal("'authTicket' flag  OR 'remotepath' & 'allocation' flag required")
		}

		alloc, err = sdk.GetAllocationFromAuthTicket(authTicket)
		if err != nil {
			log.Fatal("can't get allocation object: ", err)
		}
		var at = sdk.InitAuthTicket(authTicket)

		if lookupHash == "" {
			if lookupHash, err = at.GetLookupHash(); err != nil {
				log.Fatal("can't get lookup hash from auth ticket: ", err)
			}
		}

		meta, err = alloc.GetFileMetaFromAuthTicket(authTicket, lookupHash)
		if err != nil {
			log.Fatal("can't get file meta: ", err)
		}

		downloadCost(alloc, meta, blocksPerMarker)
	},
}

func init() {
	rootCmd.AddCommand(getDownloadCostCmd)
	dcpf := getDownloadCostCmd.PersistentFlags()
	dcpf.String("allocation", "", "allocation ID, required")
	dcpf.String("remotepath", "", "remote path of file")
	dcpf.Int("blocks-per-marker", 10, "blocks signed per Read Marker")
	dcpf.String("authticket", "", "authticket")
	dcpf.String("lookuphash", "", "lookuphash, for the remote file")
	getDownloadCostCmd.MarkFlagRequired("allocation")
}
