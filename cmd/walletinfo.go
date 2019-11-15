package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/zboxcore/client"
	"github.com/0chain/zboxcli/util"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

// walletinfo used for getting the wallet info
var walletinfoCmd = &cobra.Command{
	Use:   "getwallet",
	Short: "Get wallet information",
	Long:  `Get wallet information`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		header := []string{"Public Key", "ClientID", "Encryption Public Key"}
		data := make([][]string, 1)
		encPubKey, err := sdk.GetClientEncryptedPublicKey()
		if err != nil {
			fmt.Println("Error getting the public key for encryption. ", err.Error())
			return
		}
		data[0] = []string{client.GetClientPublicKey(), client.GetClientID(), encPubKey}
		util.WriteTable(os.Stdout, header, []string{}, data)
		return
	},
}

func init() {
	rootCmd.AddCommand(walletinfoCmd)
}
