package cmd

import (
	"log"

	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

var zauthCmd = &cobra.Command{
	Use:   "zauth",
	Short: "Enable zauth",
	Long:  `Enable zauth to sign transactions and messages, setup split keys and configure the zauth service.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Add your code here
		serverAddr, err := cmd.Flags().GetString("server")
		if err != nil {
			log.Fatalf("Could not find zauth server address")
		}

		token, err := cmd.Flags().GetString("token")
		if err != nil {
			log.Fatalf("Could not find zauth access token")
		}

		// update or setup the zauth server address
		gConfig.Set("zauth.server", serverAddr)
		if err := gConfig.WriteConfig(); err != nil {
			log.Fatalf("Could not save config: %v", err)
		}

		if clientWallet == nil {
			log.Fatalf("Wallet is initialized yet")
		}

		if clientWallet.IsSplit {
			log.Fatalln("Wallet is already split")
		}

		sw, err := zcncore.SplitKeysWallet(clientWallet.Keys[0].PrivateKey, 2)
		if err != nil {
			log.Fatalf("Failed to split keys: %v", err)
		}

		if err := zcncore.CallZauthSetup(serverAddr, token, zcncore.SplitWallet{
			ClientID:      sw.ClientID,
			ClientKey:     sw.ClientKey,
			PublicKey:     sw.Keys[1].PublicKey,
			PrivateKey:    sw.Keys[1].PrivateKey,
			PeerPublicKey: sw.Keys[0].PublicKey,
		}); err != nil {
			log.Fatalf("Failed to setup zauth server: %v", err)
		}

		// remove the keys[1]
		sw.PeerPublicKey = sw.Keys[1].PublicKey
		sw.Keys = sw.Keys[:1]
		clientWallet.SetSplitKeys(sw)
		if err := clientWallet.SaveTo(walletFilePath); err != nil {
			log.Fatalf("Failed to save wallet: %v", err)
		}

		log.Printf("Setup zauth server successfully")
	},
}

func init() {
	rootCmd.AddCommand(zauthCmd)
	zauthCmd.PersistentFlags().String("server", "s", "The zauth server address")
	zauthCmd.PersistentFlags().String("token", "t", "The /setup access token")
	zauthCmd.MarkFlagRequired("server")
}
