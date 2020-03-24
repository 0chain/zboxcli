package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func printChallengePoolInfo(info *sdk.ChallengePoolInfo) {
	var header = []string{
		"START", "DUR.", "TIME LEFT", "LOCKED", "BALANCE",
	}
	var data = [][]string{{
		time.Unix(info.StartTime, 0).String(),
		info.Duration.String(),
		info.TimeLeft.String(),
		fmt.Sprint(info.Locked),
		fmt.Sprint(zcncore.ConvertToToken(info.Balance)),
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
		printChallengePoolInfo(info)
	},
}

func init() {
	rootCmd.AddCommand(cpInfo)

	cpInfo.PersistentFlags().String("allocation", "",
		"allocation identifier, required")
}
