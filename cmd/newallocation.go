package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

var (
	datashards, parityshards *int
	size                     *int64
	allocationFileName       *string
)

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
				if fillf, err = flags.GetFloat64("fill"); err != nil {
					log.Fatal("error: invalid 'fill' value:", err)
				}
				fill = zcncore.ConvertToValue(fillf)
			}
		}

		allocationID, err := sdk.CreateAllocation(*datashards, *parityshards,
			*size, common.Now()+7776000, fill)
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
