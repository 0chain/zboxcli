package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/blockchain"
	. "github.com/0chain/gosdk/zboxcore/logger"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	// "github.com/0chain/gosdk/zcncore"
	"github.com/0chain/gosdk/zboxcore/fileref"

	"github.com/spf13/cobra"
)

func sizePerBlobber(size int64, data, parity int) int64 {
	return ((size + int64(data) - 1) / int64(data)) * (int64(data) + int64(parity))
}

// getallocationCmd represents the get allocation info command
var getallocationCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the allocation info",
	Long:  `Gets the allocation info`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		doJSON, _ := cmd.Flags().GetBool("json")
		alloc, err := sdk.GetAllocation(allocationID)
		if err != nil {
			Logger.Error("Error fetching the allocation", err)
			log.Fatal("Error fetching/verifying the allocation")
		}
		if doJSON {
			util.PrintJSON(alloc)
			return
		}
		var getBaseURL = func(bid string, bs []*blockchain.StorageNode) string {
			for _, b := range bs {
				if b.ID == bid {
					return b.Baseurl
				}
			}
			return "(not found)"
		}

		var priceRangeString = func(pr sdk.PriceRange) string {
			return fmt.Sprintf("%s-%s", common.Balance(pr.Min), common.Balance(pr.Max))
		}

		fmt.Println("allocation:")
		fmt.Println("  id:             ", alloc.ID)
		fmt.Println("  tx:             ", alloc.Tx, "(latest create/update allocation transaction hash)")
		fmt.Println("  data_shards:    ", alloc.DataShards)
		fmt.Println("  parity_shards:  ", alloc.ParityShards)
		fmt.Println("  size:           ", common.Size(alloc.Size))
		fmt.Println("  expiration_date:", common.Timestamp(alloc.Expiration).ToTime())
		fmt.Println("  blobbers:")

		var (
			gbsize   = sizePerBlobber(1*GB, alloc.DataShards, alloc.ParityShards)
			gbblocks = sizeInGB(maxInt64(gbsize/(64*KB), 2) * 64 * KB)

			totalRead, totalWrite common.Balance
		)

		for _, d := range alloc.BlobberDetails {

			totalRead += common.Balance(float64(d.Terms.ReadPrice) * gbblocks)
			totalWrite += common.Balance(float64(d.Terms.WritePrice) * sizeInGB(gbsize))

			fmt.Println("    - blobber_id:      ", d.BlobberID)
			fmt.Println("      base URL:        ", getBaseURL(d.BlobberID, alloc.Blobbers))
			fmt.Println("      size:            ", common.Size(d.Size))
			fmt.Println("      min_lock_deman:  ", common.Balance(d.MinLockDemand))
			fmt.Println("      spent:           ", common.Balance(d.Spent), "(moved to challenge pool or to the blobber)")
			fmt.Println("      penalty:         ", common.Balance(d.Penalty), "(blobber stake slash)")
			fmt.Println("      read_reward:     ", common.Balance(d.ReadReward))
			fmt.Println("      returned:        ", common.Balance(d.Returned), "(on challenge failed)")
			fmt.Println("      challenge_reward:", common.Balance(d.ChallengeReward), "(on challenge passed)")
			fmt.Println("      final_reward:    ", common.Balance(d.FinalReward), "(if finalized)")
			fmt.Println("      terms: (allocation related terms)")
			fmt.Println("        read_price:               ", d.Terms.ReadPrice, "tok / GB (by 64KB chunks)")
			fmt.Println("        write_price:              ", d.Terms.WritePrice, "tok / GB")
			fmt.Println("        min_lock_demand:          ", d.Terms.MinLockDemand*100, "%")
			fmt.Println("        max_offer_duration:       ", d.Terms.MaxOfferDuration)
			fmt.Println("        challenge_completion_time:", d.Terms.ChallengeCompletionTime)
		}

		fmt.Println("  read_price_range:         ", priceRangeString(alloc.ReadPriceRange), "(requested)")
		fmt.Println("  write_price_range:        ", priceRangeString(alloc.WritePriceRange), "(requested)")
		fmt.Println("  challenge_completion_time:", alloc.ChallengeCompletionTime, "(max)")
		fmt.Println("  start_time:               ", common.Timestamp(alloc.StartTime).ToTime())
		fmt.Println("  finalized:                ", alloc.Finalized)
		fmt.Println("  canceled:                 ", alloc.Canceled)
		fmt.Println("  moved_to_challenge:       ", common.Balance(alloc.MovedToChallenge))
		fmt.Println("  moved_back:               ", common.Balance(alloc.MovedBack))
		fmt.Println("  moved_to_validators:      ", common.Balance(alloc.MovedToValidators))

		fmt.Println("  stats:")
		fmt.Println("    total size:             ", common.Size(alloc.Size))
		fmt.Println("    used size:              ", common.Size(alloc.Stats.UsedSize))
		fmt.Println("    number of writes:       ", alloc.Stats.NumWrites)
		fmt.Println("    total challenges:       ", alloc.Stats.TotalChallenges)
		fmt.Println("    passed challenges:      ", alloc.Stats.SuccessChallenges)
		fmt.Println("    failed challenges:      ", alloc.Stats.FailedChallenges)
		fmt.Println("    open challenges:        ", alloc.Stats.OpenChallenges)
		fmt.Println("    last challenge redeemed:", alloc.Stats.LastestClosedChallengeTxn)

		fmt.Println("  price:")
		fmt.Println("    read_price: ", totalRead, "tok / GB (by 64KB)")
		fmt.Println("    write_price:", totalWrite, "tok / GB")
		return
	},
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

func sizeInGB(size int64) float64 {
	return float64(size) / GB
}

func downloadCost(alloc *sdk.Allocation, meta *sdk.ConsolidatedFileMeta) {

	if meta.Type != fileref.FILE {
		log.Fatal("not a file")
	}

	var (
		bsize  = sizePerBlobber(meta.Size, alloc.DataShards, alloc.ParityShards)
		blocks = maxInt64(bsize/fileref.CHUNK_SIZE, 2)
		size   = sizeInGB(blocks * fileref.CHUNK_SIZE)
		cost   common.Balance
	)

	for _, d := range alloc.BlobberDetails {
		cost += common.Balance(float64(d.Terms.ReadPrice) * size)
	}

	fmt.Printf("%s tokens for %d 64KB blocks (%s) of %s", cost, blocks,
		common.Size(meta.Size), meta.Path)
	fmt.Println()
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

		if remotePath == "" && authTicket == "" {
			log.Fatal("'remotepath' or 'authticket' flag required")
		}

		var (
			alloc *sdk.Allocation
			meta  *sdk.ConsolidatedFileMeta
		)

		if remotePath != "" {

			// by remote path

			if alloc, err = sdk.GetAllocation(allocID); err != nil {
				log.Fatal("fetching the allocation: ", err)
			}

			if meta, err = alloc.GetFileMeta(remotePath); err != nil {
				log.Fatal("can't get file meta: ", err)
			}

			downloadCost(alloc, meta)
			return
		}

		// by authentication ticket

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

		downloadCost(alloc, meta)
	},
}

func uploadCost(alloc *sdk.Allocation, size int64, path string) {

	var (
		gb   = sizeInGB(sizePerBlobber(size, alloc.DataShards, alloc.ParityShards))
		cost common.Balance
	)

	for _, d := range alloc.BlobberDetails {
		cost += common.Balance(float64(d.Terms.WritePrice) * gb)
	}

	fmt.Printf("%s tokens for %s of %s", cost, common.Size(size), path)
	fmt.Println()
}

// The getUploadCostCmd returns value in tokens to upload a file.
var getUploadCostCmd = &cobra.Command{
	Use:   "get-upload-cost",
	Short: "Get uploading cost",
	Long:  `Get uploading cost`,
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

		var localPath string

		if !fflags.Changed("localpath") {
			log.Fatal("missing requried 'localpath' flag")
		}

		if localPath, err = fflags.GetString("localpath"); err != nil {
			log.Fatal("invalid 'localpath' flag: ", err)
		}

		if localPath == "" {
			log.Fatal("empty local path")
		}

		var alloc *sdk.Allocation
		if alloc, err = sdk.GetAllocation(allocID); err != nil {
			log.Fatal("fetching the allocation: ", err)
		}

		var fi os.FileInfo
		if fi, err = os.Stat(localPath); err != nil {
			log.Fatal(err)
		}

		if !fi.Mode().IsRegular() {
			log.Fatal("not a regular file")
		}

		uploadCost(alloc, fi.Size(), localPath)
	},
}

func init() {
	rootCmd.AddCommand(getallocationCmd)
	rootCmd.AddCommand(getDownloadCostCmd)
	rootCmd.AddCommand(getUploadCostCmd)

	getallocationCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	getallocationCmd.MarkFlagRequired("allocation")
	getallocationCmd.Flags().Bool("json", false, "pass this option to print response as json data")

	dcpf := getDownloadCostCmd.PersistentFlags()
	dcpf.String("allocation", "", "allocation ID, required")
	dcpf.String("remotepath", "", "remote path of file")
	dcpf.String("authticket", "", "authticket")
	dcpf.String("lookuphash", "", "lookuphash, for the remote file")
	getDownloadCostCmd.MarkFlagRequired("allocation")

	ucpf := getUploadCostCmd.PersistentFlags()
	ucpf.String("allocation", "", "allocation ID, required")
	ucpf.String("localpath", "", "local file path, required")
	getUploadCostCmd.MarkFlagRequired("allocation")
	getUploadCostCmd.MarkFlagRequired("localpath")
}
