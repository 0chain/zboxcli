package cmd

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

var (
	datashards, parityshards *int
	size                     *int64
	allocationFileName       *string
)

func getPriceRange(val string) (pr sdk.PriceRange, err error) {
	var ss = strings.Split(val, "-")
	if len(ss) != 2 {
		err = fmt.Errorf("invalid price range format: %q", val)
		return
	}
	var minf, maxf float64
	if minf, err = strconv.ParseFloat(ss[0], 64); err != nil {
		return
	}
	if maxf, err = strconv.ParseFloat(ss[1], 64); err != nil {
		return
	}
	pr.Min = zcncore.ConvertToValue(minf)
	pr.Max = zcncore.ConvertToValue(maxf)
	return
}

// newallocationCmd represents the new allocation command
var newallocationCmd = &cobra.Command{
	Use:   "newallocation",
	Short: "Creates a new allocation",
	Long:  `Creates a new allocation`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if datashards == nil || parityshards == nil || size == nil {
			log.Fatal("Invalid allocation parameters.")
		}

		var (
			flags = cmd.Flags() //
			lock  int64         // lock with given number of tokens
			err   error         //
		)
		costOnly, _ := cmd.Flags().GetBool("cost")
		if !costOnly {
			if !flags.Changed("lock") {
				log.Fatal("missing required 'lock' argument")
			}
		}

		convertFromUSD, _ := cmd.Flags().GetBool("usd")

		var lockf float64
		if lockf, err = flags.GetFloat64("lock"); err != nil {
			log.Fatal("error: invalid 'lock' value:", err)
		}

		if convertFromUSD {
			lockf, err = zcncore.ConvertUSDToToken(lockf)
			if err != nil {
				log.Fatal("error: failed to convert to USD : ", err)
			}
		}
		lock = zcncore.ConvertToValue(lockf)

		var (
			readPrice  = sdk.PriceRange{Min: 0, Max: math.MaxInt64}
			writePrice = sdk.PriceRange{Min: 0, Max: math.MaxInt64}

			mcct time.Duration = 1 * time.Hour
		)

		if flags.Changed("read_price") {
			rps, err := flags.GetString("read_price")
			if err != nil {
				log.Fatal("invalid read_price value: ", err)
			}
			pr, err := getPriceRange(rps)
			if err != nil {
				log.Fatal("invalid read_price value: ", err)
			}
			readPrice = pr
		}

		if flags.Changed("write_price") {
			wps, err := flags.GetString("write_price")
			if err != nil {
				log.Fatal("invalid write_price value: ", err)
			}
			pr, err := getPriceRange(wps)
			if err != nil {
				log.Fatal("invalid write_price value: ", err)
			}
			writePrice = pr
		}

		if flags.Changed("mcct") {
			if mcct, err = flags.GetDuration("mcct"); err != nil {
				log.Fatal("invalid mcct value: ", err)
			}
			if mcct <= 1*time.Second {
				log.Fatal("invalid mcct value < 1s")
			}
		}

		var expire time.Duration
		if expire, err = flags.GetDuration("expire"); err != nil {
			log.Fatal("invalid 'expire' flag: ", err)
		}

		var expireAt = time.Now().Add(expire).Unix()

		if costOnly {
			minCost, err := sdk.GetAllocationMinLock(*datashards, *parityshards, *size, expireAt, readPrice, writePrice, mcct)
			if err != nil {
				log.Fatal("Error fetching cost: ", err)
			}
			log.Print("Cost for the given allocation: ", zcncore.ConvertToToken(minCost))
		} else {
			allocationID, err := sdk.CreateAllocation(*datashards, *parityshards,
				*size, expireAt, readPrice, writePrice, mcct, lock)
			if err != nil {
				log.Fatal("Error creating allocation: ", err)
			}
			log.Print("Allocation created: ", allocationID)
			storeAllocation(allocationID)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(newallocationCmd)
	datashards = newallocationCmd.PersistentFlags().Int("data", 2, "--data 2")
	parityshards = newallocationCmd.PersistentFlags().Int("parity", 2, "--parity 2")
	size = newallocationCmd.PersistentFlags().Int64("size", 2147483648, "--size 10000")
	allocationFileName = newallocationCmd.PersistentFlags().String("allocationFileName", "allocation.txt", "--allocationFileName allocation.txt")
	newallocationCmd.PersistentFlags().
		Float64("lock", 0.0,
			"lock write pool with given number of tokens, required")
	newallocationCmd.PersistentFlags().
		String("read_price", "",
			"select blobbers by provided read price range, use form 0.5-1.5, default is [0; inf)")
	newallocationCmd.PersistentFlags().
		String("write_price", "",
			"select blobbers by provided write price range, use form 1.5-2.5, default is [0; inf)")
	newallocationCmd.PersistentFlags().
		Duration("expire", 720*time.Hour, "duration to allocation expiration")
	newallocationCmd.PersistentFlags().
		Duration("mcct", 1*time.Hour,
			"max challenge completion time, optional, default 1h")
	newallocationCmd.Flags().
		Bool("usd", false,
			"pass this option to give token value in USD")
	newallocationCmd.Flags().
		Bool("cost", false,
			"pass this option to only get the min lock demand")

	newallocationCmd.MarkFlagRequired("data")
	newallocationCmd.MarkFlagRequired("parity")
	newallocationCmd.MarkFlagRequired("size")
	newallocationCmd.MarkFlagRequired("allocationFileName")
}

func storeAllocation(allocationID string) {

	allocFilePath := util.GetConfigDir() + string(os.PathSeparator) + *allocationFileName

	file, err := os.Create(allocFilePath)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	//Only one allocation ID per file.
	fmt.Fprintf(file, allocationID)

}
