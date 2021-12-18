package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// allocRenameCmd represents alloc-rename command
var allocRenameCmd = &cobra.Command{
	Use:   "alloc-rename",
	Short: "rename allocation in blobber",
	Long:  `rename allocation in blobber`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags() // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false {
			PrintError("Error: allocation flag is missing")
			os.Exit(1)
		}

		if fflags.Changed("name") == false {
			PrintError("Error: name flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		name := cmd.Flag("name").Value.String()

		err = allocationObj.AllocationRename(name)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}
		fmt.Println(name + " renamed")

	},
}

func init() {
	rootCmd.AddCommand(allocRenameCmd)
	allocRenameCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	allocRenameCmd.PersistentFlags().String("name", "", "New Name for the allocation")
	allocRenameCmd.MarkFlagRequired("allocation")
	allocRenameCmd.MarkFlagRequired("name")
}
