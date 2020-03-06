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
			fill  int64         // fill with given number of tokens
			auto  bool          // fill automatically
		)
		if flags.Changed("fill") {
			var fills, err = flags.GetString("fill")
			if err != nil {
				log.Fatal("error: invalid 'fill' value:", err)
			}
			if fills == "auto" {
				auto = true
			} else {
				var fillf float64
				if fillf, err = strconv.ParseFloat(fills, 64); err != nil {
					log.Fatal("error: invalid 'fill' value: ", err)
				}
				fill = zcncore.ConvertToValue(fillf)
			}
		}

		var (
			readPrice  = sdk.PriceRange{Min: 0, Max: math.MaxInt64}
			writePrice = sdk.PriceRange{Min: 0, Max: math.MaxInt64}
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
			readPrice = pr
		}

		var expire, err = flags.GetDuration("expire")
		if err != nil {
			log.Fatal("invalid 'expire' flag: ", err)
		}

		var expireAt = time.Now().Add(expire).Unix()

		allocationID, err := sdk.CreateAllocation(*datashards, *parityshards,
			*size, expireAt, readPrice, writePrice, fill)
		if err != nil {
			log.Fatal("Error creating allocation: ", err)
		}
		log.Print("Allocation created: ", allocationID)
		storeAllocation(allocationID)

		if auto {
			allocationObj, err := sdk.GetAllocation(allocationID)
			if err != nil {
				log.Fatal("Error filling allocation write pool: ", err)
			}
			resp, err := sdk.WritePoolLock(allocationID, allocationObj.MinLockDemand)
			if err != nil {
				log.Fatal("Error filling allocation write pool: ", err)
			}
			log.Println(resp)
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
	newallocationCmd.PersistentFlags().String("fill", "0", "fill write pool with given number of tokens, or use 'auto'")
	newallocationCmd.PersistentFlags().String("read_price", "", "select blobbers by provided read price range, use form 0.5-1.5, default is [0; inf)")
	newallocationCmd.PersistentFlags().String("write_price", "", "select blobbers by provided write price range, use form 1.5-2.5, default is [0; inf)")
	newallocationCmd.PersistentFlags().Duration("expire", 48*time.Hour, "duration to allocation expiration")

	newallocationCmd.MarkFlagRequired("data")
	newallocationCmd.MarkFlagRequired("parity")
	newallocationCmd.MarkFlagRequired("size")
	newallocationCmd.MarkFlagRequired("allocationFileName")
}

func storeAllocation(allocationID string) {

	allocFilePath := getConfigDir() + "/" + *allocationFileName

	file, err := os.Create(allocFilePath)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	//Only one allocation ID per file.
	fmt.Fprintf(file, allocationID)

}
