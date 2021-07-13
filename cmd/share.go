package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// shareCmd represents share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "share files from blobbers",
	Long:  `share files from blobbers`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		if fflags.Changed("remotepath") == false {
			PrintError("Error: remotepath flag is missing")
			os.Exit(1)
		}
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			PrintError("Error fetching the allocation", err)
			os.Exit(1)
		}
		remotepath := cmd.Flag("remotepath").Value.String()
		refType := fileref.FILE
		statsMap, err := allocationObj.GetFileStats(remotepath)
		if err != nil {
			PrintError("Error in getting information about the object." + err.Error())
			os.Exit(1)
		}
		isFile := false
		for _, v := range statsMap {
			if v != nil {
				isFile = true
				break
			}
		}
		if !isFile {
			refType = fileref.DIRECTORY
		}

		var fileName string
		_, fileName = filepath.Split(remotepath)
		refereeClientID := cmd.Flag("clientid").Value.String()
		encryptionpublickey := cmd.Flag("encryptionpublickey").Value.String()
		ref, err := allocationObj.GetAuthTicket(remotepath, fileName, refType, refereeClientID, encryptionpublickey)
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}
		fmt.Println("auth ticket :" + ref)
		return
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	shareCmd.PersistentFlags().String("remotepath", "", "Remote path to share")
	shareCmd.PersistentFlags().String("clientid", "", "ClientID of the user to share with. Leave blank for public share")
	shareCmd.PersistentFlags().String("encryptionpublickey", "", "Encryption public key of the client you want to share with. Can be retrieved by the getwallet command")
	shareCmd.MarkFlagRequired("allocation")
	shareCmd.MarkFlagRequired("remotepath")
}
