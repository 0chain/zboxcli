package cmd

import (
	"fmt"
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(totalStoredData)
	totalStoredData.PersistentFlags().Bool("json", false, "pass this option to print response as json data")
}

var totalStoredData = &cobra.Command{
	Use:   "total-stored-data",
	Short: "Total data currently stored across all blobbers.",
	Long:  `Total data currently stored across all blobbers.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		doJSON, _ := cmd.Flags().GetBool("json")
		var info = make(map[string]int64)
		if info, err = sdk.GetTotalStoredData(); err != nil {
			log.Fatalf("Failed to get total stored data: %v", err)
		}
		if doJSON {
			util.PrintJSON(info)
		} else {
			fmt.Println("total stored data:", info["total-stored-data"])
		}
	},
}
