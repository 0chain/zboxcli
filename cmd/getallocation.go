package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/blockchain"
	"github.com/0chain/gosdk/zboxcore/fileref"
	. "github.com/0chain/gosdk/zboxcore/logger"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"

	"github.com/spf13/cobra"
)

func sizePerBlobber(size int64, data, parity int) (part int64) {
	var dp = data + parity
	part = (size + int64(dp-1)) / int64(dp)
	return
}

func perShard(size int64, data, parity int) (ps int64) {
	ps = (size + int64(data) - 1) / int64(data)
	return
}

func uploadCostForBlobber(price float64, size int64, data, parity int) (
	cost common.Balance) {

	var ps = (size + int64(data) - 1) / int64(data)
	ps = ps * int64(data+parity)

	return common.Balance(price * sizeInGB(ps))
}

func downloadCostFor1GB(alloc *sdk.Allocation) (cost common.Balance) {
	var (
		ps   = perShard(1*GB, alloc.DataShards, alloc.ParityShards)
		cps  = (ps + fileref.CHUNK_SIZE - 1) / fileref.CHUNK_SIZE
		size float64 // GB
	)
	// the Go SDK requests block by 10
	if cps%10 > 0 {
		cps = ((cps/10)*10 + 10)
	}
	size = sizeInGB(cps * fileref.CHUNK_SIZE)
	for _, d := range alloc.BlobberDetails {
		cost += common.Balance(float64(d.Terms.ReadPrice) * size)
	}
	cost = cost / common.Balance(len(alloc.BlobberDetails))
	return
}

func uploadCostFor1GB(alloc *sdk.Allocation) (cost common.Balance) {
	for _, d := range alloc.BlobberDetails {
		cost += uploadCostForBlobber(float64(d.Terms.WritePrice), 1*GB,
			alloc.DataShards, alloc.ParityShards)
	}
	return
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

		for _, d := range alloc.BlobberDetails {
			fmt.Println("    - blobber_id:      ", d.BlobberID)
			fmt.Println("      base URL:        ", getBaseURL(d.BlobberID, alloc.Blobbers))
			fmt.Println("      size:            ", common.Size(d.Size))
			fmt.Println("      min_lock_demand: ", common.Balance(d.MinLockDemand))
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

		if len(alloc.Curators) < 1 {
			fmt.Println("  no curators")
		} else if len(alloc.Curators) == 1 {
			fmt.Println("  curator: " + alloc.Curators[0])
		} else {
			fmt.Println("  curators:")
			for _, curator := range alloc.Curators {
				fmt.Println("  ", curator)
			}
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
		fmt.Println("    time_unit:  ", alloc.TimeUnit)
		fmt.Println("    read_price: ", downloadCostFor1GB(alloc), "tok / GB (by 64KB)")
		fmt.Println("    write_price:", uploadCostFor1GB(alloc),
			fmt.Sprintf("tok / GB / %s", alloc.TimeUnit))
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
		ps  = perShard(meta.Size, alloc.DataShards, alloc.ParityShards)
		cps = (ps + fileref.CHUNK_SIZE - 1) / fileref.CHUNK_SIZE

		size float64        // GB
		cost common.Balance //
	)

	// the Go SDK requests block by 10
	if cps%10 > 0 {
		cps = ((cps/10)*10 + 10)
	}

	size = sizeInGB(cps * fileref.CHUNK_SIZE)

	for _, d := range alloc.BlobberDetails {
		cost += common.Balance(float64(d.Terms.ReadPrice) * size)
	}

	cost = cost / common.Balance(len(alloc.BlobberDetails))

	fmt.Printf("%s tokens for %d 64KB blocks (%s) of %s", cost, cps,
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

// The uploadCost for a size and duration (if given). If the duration is zero
// of less, then it returns upload cost until allocation ends.
func uploadCost(alloc *sdk.Allocation, size int64, path string,
	duration time.Duration) {

	var cost common.Balance // total price for size / duration

	for _, d := range alloc.BlobberDetails {
		cost += uploadCostForBlobber(float64(d.Terms.WritePrice), size,
			alloc.DataShards, alloc.ParityShards)
	}

	switch {
	case duration == 0:
		fmt.Printf("%s tokens / %s for %s of %s",
			cost, alloc.TimeUnit, common.Size(size), path)
	case duration < 0:
		fmt.Println("allocation expired, 'end' flag can't be used")
		return
	default:
		var dtu = float64(duration) / float64(alloc.TimeUnit)
		cost = common.Balance(float64(cost) * dtu)
		fmt.Printf("%s tokens / %s for %s of %s",
			cost, duration, common.Size(size), path)
	}

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
			fflags   = cmd.Flags()
			allocID  string
			duration time.Duration
			end      bool
			err      error
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

		var fi os.FileInfo
		if fi, err = os.Stat(localPath); err != nil {
			log.Fatal(err)
		}

		if !fi.Mode().IsRegular() {
			log.Fatal("not a regular file")
		}

		if duration, err = fflags.GetDuration("duration"); err != nil {
			log.Fatal("invalid 'duration' flag:", err)
		} else if duration < 0 {
			log.Fatal("negative duration not allowed: ", duration)
		}

		if end, err = fflags.GetBool("end"); err != nil {
			log.Fatal("invalid 'end' flag:", err)
		}

		var alloc *sdk.Allocation
		if alloc, err = sdk.GetAllocation(allocID); err != nil {
			log.Fatal("fetching the allocation: ", err)
		}

		// until allocation ends
		if end {
			var expiry = time.Unix(alloc.Expiration, 0)
			duration = expiry.Sub(time.Now())
		}

		uploadCost(alloc, fi.Size(), localPath, duration)
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
	ucpf.Duration("duration", 0, "expected duration keep uploaded file")
	ucpf.Bool("end", false, "use the duration until allocation ends")
	getUploadCostCmd.MarkFlagRequired("allocation")
	getUploadCostCmd.MarkFlagRequired("localpath")
}
