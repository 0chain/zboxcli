package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func printChallengePoolInfo(info *sdk.ChallengePoolInfo) {
	var header = []string{
		"BALANCE", "START", "EXPIRE", "FINIALIZED",
	}
	var data = [][]string{{
		info.Balance.String(),
		info.StartTime.ToTime().String(),
		info.Expiration.ToTime().String(),
		fmt.Sprint(info.Finalized),
	}}
	fmt.Println("POOL ID:", info.ID)
	util.WriteTable(os.Stdout, header, []string{}, data)
	fmt.Println()
}

// cpInfo information
var cpInfo = &cobra.Command{
	Use:   "cp-info",
	Short: "Challenge pool information.",
	Long:  `Challenge pool information.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flags   = cmd.Flags()
			allocID string
			err     error
		)

		doJSON, _ := cmd.Flags().GetBool("json")

		if !flags.Changed("allocation") {
			log.Fatal("missing required 'allocation' flag")
		}

		if allocID, err = flags.GetString("allocation"); err != nil {
			log.Fatalf("can't get 'allocation' flag: %v", err)
		}

		var info *sdk.ChallengePoolInfo
		if info, err = sdk.GetChallengePoolInfo(allocID); err != nil {
			log.Fatalf("Failed to get challenge pool info: %v", err)
		}
		if doJSON {
			util.PrintJSON(info)
		} else {
			printChallengePoolInfo(info)
		}
	},
}

func init() {
	rootCmd.AddCommand(cpInfo)

	cpInfo.PersistentFlags().String("allocation", "",
		"allocation identifier, required")
	cpInfo.Flags().Bool("json", false, "(default false) pass this option to print response as json data")
}
