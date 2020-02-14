package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var updateSize, updateExpiry *int64

// newallocationCmd represents the new allocation command
var updateallocationCmd = &cobra.Command{
	Use:   "updateallocation",
	Short: "Updates allocation's expiry and size",
	Long:  `Updates allocation's expiry and size`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and return
		}
		if updateSize == nil || updateExpiry == nil {
			PrintError("Invalid updateallocation parameters.")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		txnHash, err := sdk.UpdateAllocation(*updateSize, *updateExpiry, allocationID)
		if err != nil {
			PrintError("Error creating allocation." + err.Error())
			os.Exit(1)
		}
		fmt.Println("Allocation updated with txId : " + txnHash)
		return
	},
}

func init() {
	rootCmd.AddCommand(updateallocationCmd)
	updateallocationCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	updateSize = updateallocationCmd.PersistentFlags().Int64("size", 2147483648, "--size 10000")
	updateExpiry = updateallocationCmd.PersistentFlags().Int64("expiry", 2592000, "--size 10000")
	updateallocationCmd.MarkFlagRequired("allocation")
	updateallocationCmd.MarkFlagRequired("size")
	updateallocationCmd.MarkFlagRequired("expiry")
}
