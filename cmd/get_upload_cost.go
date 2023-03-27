package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

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
			duration = time.Until(expiry)
		}

		uploadCost(alloc, fi.Size(), localPath, duration)
	},
}

// The uploadCost for a size and duration (if given). If the duration is zero
// of less, then it returns upload cost until allocation ends.
func uploadCost(alloc *sdk.Allocation, size int64, path string,
	duration time.Duration) {

	var cost common.Balance // total price for size / duration

	for _, d := range alloc.BlobberDetails {
		cost += uploadCostForBlobber(float64(d.Terms.WritePrice), size,
			alloc.DataShards)
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

func uploadCostForBlobber(price float64, size int64, data int) (
	cost common.Balance) {

	var ps = (size + int64(data) - 1) / int64(data)

	return common.Balance(price * sizeInGB(ps))
}

func uploadCostFor1GB(alloc *sdk.Allocation) (cost common.Balance) {
	for _, d := range alloc.BlobberDetails {
		cost += uploadCostForBlobber(float64(d.Terms.WritePrice), 1*GB,
			alloc.DataShards)
	}
	return
}

func init() {
	rootCmd.AddCommand(getUploadCostCmd)
	ucpf := getUploadCostCmd.PersistentFlags()
	ucpf.String("allocation", "", "allocation ID, required")
	ucpf.String("localpath", "", "local file path, required")
	ucpf.Duration("duration", 0, "expected duration keep uploaded file")
	ucpf.Bool("end", false, "use the duration until allocation ends")
	getUploadCostCmd.MarkFlagRequired("allocation")
	getUploadCostCmd.MarkFlagRequired("localpath")

}
