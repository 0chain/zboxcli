package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var datashards, parityshards *int
var size *int64
var allocationFileName *string

// newallocationCmd represents the new allocation command
var newallocationCmd = &cobra.Command{
	Use:   "newallocation",
	Short: "Creates a new allocation",
	Long:  `Creates a new allocation`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if datashards == nil || parityshards == nil || size == nil {
			PrintError("Invalid allocation parameters.")
			os.Exit(1)
		}
		allocationID, err := sdk.CreateAllocation(*datashards, *parityshards, *size, common.Now()+7776000)
		if err != nil {
			PrintError("Error creating allocation." + err.Error())
			os.Exit(1)
		}
		fmt.Println("Allocation created : " + allocationID)
		storeAllocation(allocationID)
		return
	},
}

func init() {
	rootCmd.AddCommand(newallocationCmd)
	datashards = newallocationCmd.PersistentFlags().Int("data", 2, "--data 2")
	parityshards = newallocationCmd.PersistentFlags().Int("parity", 2, "--parity 2")
	size = newallocationCmd.PersistentFlags().Int64("size", 2147483648, "--size 10000")
	allocationFileName = newallocationCmd.PersistentFlags().String("allocationFileName", "allocation.txt", "--allocationFileName allocation.txt")

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
