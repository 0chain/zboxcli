package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/0chain/gosdk/core/common"
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
		fflags := cmd.Flags()              // fflags is a *flag.FlagSet
		if !fflags.Changed("allocation") { // check if the flag "path" is set
			PrintError("Error: allocation flag is missing") // If not, we'll let the user know
			os.Exit(1)                                      // and os.Exit(1)
		}
		if !fflags.Changed("remotepath") {
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

		var isFile bool
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

		revoke, _ := cmd.Flags().GetBool("revoke")
		if revoke {
			err := allocationObj.RevokeShare(remotepath, refereeClientID)
			if err != nil {
				PrintError(err.Error())
				os.Exit(1)
			}
			fmt.Println("Share revoked for client " + refereeClientID)
		} else {
			expiration, err := cmd.Flags().GetInt64("expiration-seconds")
			if err != nil {
				PrintError(err.Error())
				os.Exit(1)
			}

			availableAfter := time.Now()
			availableAfterInput, err := cmd.Flags().GetString("available-after")
			if err != nil {
				PrintError(err.Error())
				os.Exit(1)
			}
			if len(availableAfterInput) > 0 {
				aa, err := common.ParseTime(availableAfter, availableAfterInput)
				if err != nil {
					PrintError(err.Error())
					os.Exit(1)
				}
				availableAfter = *aa
			}
			whoPays, err := cmd.Flags().GetInt("who-pays")
			if err != nil {
				PrintError(err.Error())
				os.Exit(1)
			}
			encryptionpublickey := cmd.Flag("encryptionpublickey").Value.String()
			ref, err := allocationObj.GetAuthTicket(
				remotepath, fileName, refType, refereeClientID, encryptionpublickey,
				whoPays, expiration, &availableAfter)

			if err != nil {
				PrintError(err.Error())
				os.Exit(1)
			}
			fmt.Println("Auth token :" + ref)
		}
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	shareCmd.PersistentFlags().String("remotepath", "", "Remote path to share")
	shareCmd.PersistentFlags().String("clientid", "", "ClientID of the user to share with. Leave blank for public share")
	shareCmd.PersistentFlags().String("encryptionpublickey", "", "Encryption public key of the client you want to share with. Can be retrieved by the getwallet command")
	shareCmd.PersistentFlags().Int64("expiration-seconds", 0, "Authticket will expire when the seconds specified have elapsed after the instant of its creation")
	shareCmd.PersistentFlags().String("available-after", "", "Timelock for private file that makes the file available for download at certain time. 4 input formats are supported: +1h30m, +30, 1647858200 and 2022-03-21 10:21:38. Default value is current local time.")
	shareCmd.PersistentFlags().Bool("revoke", false, "Revoke share for remotepath")
	shareCmd.PersistentFlags().Int("who-pays", 0, "Who Pays; Owner or 3rd party. Default is owner")
	shareCmd.MarkFlagRequired("allocation")
	shareCmd.MarkFlagRequired("remotepath")
}
