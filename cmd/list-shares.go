package cmd

import (
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

// shareCmd represents share command
var listsharesCmd = &cobra.Command{
	Use:   "list-shares",
	Short: "list the files shared by the current user",
	Long:  `list the files shared by the current user`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}

		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}

		shares, err := allocationObj.ListShare()
		if err != nil {
			PrintError("failed to list shares", err)
			os.Exit(1)
		}
		util.PrintJSON(shares)

	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.PersistentFlags().String("allocation", "", "Allocation ID")
}
