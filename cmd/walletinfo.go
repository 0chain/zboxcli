package cmd

import (
	"fmt"
	"os"

	"github.com/0chain/gosdk/core/client"
	"github.com/0chain/gosdk/core/encryption"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/util"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var walletDecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt text with passphrase",
	Long:  `Decrypt text with passphrase`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		passphrase, _ := cmd.Flags().GetString("passphrase")
		text, _ := cmd.Flags().GetString("text")

		decrypted, err := zcncore.Decrypt(passphrase, text)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(decrypted)
		return
	},
}

// walletinfo used for getting the wallet info
var walletinfoCmd = &cobra.Command{
	Use:   "getwallet",
	Short: "Get wallet information",
	Long:  `Get wallet information`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		doJSON, _ := cmd.Flags().GetBool("json")

		header := []string{"Public Key", "ClientID", "Encryption Public Key"}
		data := make([][]string, 1)
		encPubKey, err := sdk.GetClientEncryptedPublicKey()
		if err != nil {
			fmt.Println("Error getting the public key for encryption. ", err.Error())
			return
		}
		data[0] = []string{client.PublicKey(), client.ClientID(), encPubKey}
		if doJSON {
			j := make(map[string]string)
			j["client_public_key"] = client.PublicKey()
			j["client_id"] = client.ClientID()
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
			data = client.ClientID()
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
	walletinfoCmd.Flags().Bool("json", false, "(default false) pass this option to print response as json data")

	rootCmd.AddCommand(signCmd)
	signCmd.Flags().String("data", "", "give data for signing, Default will be clientID")

	rootCmd.AddCommand(walletDecryptCmd)
	walletDecryptCmd.Flags().String("passphrase", "", "Passphrase to decrypt text")
	walletDecryptCmd.Flags().String("text", "", "Encrypted text")
}
