package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/core/encryption"
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
		doJSON, _ := cmd.Flags().GetBool("json")

		header := []string{"Public FileKey", "ClientID", "Encryption Public FileKey"}
		data := make([][]string, 1)
		encPubKey, err := sdk.GetClientEncryptedPublicKey()
		if err != nil {
			fmt.Println("Error getting the public key for encryption. ", err.Error())
			return
		}
		data[0] = []string{client.GetClientPublicKey(), client.GetClientID(), encPubKey}
		if doJSON {
			j := make(map[string]string)
			j["client_public_key"] = client.GetClientPublicKey()
			j["client_id"] = client.GetClientID()
			j["encryption_public_key"] = encPubKey
			util.PrintJSON(j)
			return
		}
		util.WriteTable(os.Stdout, header, []string{}, data)
		return
	},
}

var signCmd = &cobra.Command{
	Use:   "sign-data",
	Short: "Sign given data",
	Long:  `Sign given data`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := cmd.Flags().GetString("data")
		if data == "" {
			data = client.GetClientID()
		} else {
			data = encryption.Hash(data)
		}
		sign, err := client.Sign(data)
		if err != nil {
			fmt.Println("Error generating the signature. ", err.Error())
			return
		}
		fmt.Println("Signature : " + sign)
		return
	},
}

func init() {
	rootCmd.AddCommand(walletinfoCmd)
	walletinfoCmd.Flags().Bool("json", false, "pass this option to print response as json data")

	rootCmd.AddCommand(signCmd)
	signCmd.Flags().String("data", "", "give data for signing, Default will be clientID")
}
