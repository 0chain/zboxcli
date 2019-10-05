package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/0chain/gosdk/zcncore"
	"github.com/spf13/cobra"
)

// registerWalletCmd represents the register wallet command
var registerWalletCmd = &cobra.Command{
	Use:   "register",
	Short: "Registers the wallet with the blockchain",
	Long:  `Registers the wallet with the blockchain`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if clientWallet == nil {
			PrintError("Invalid wallet. Wallet not initialized in sdk")
			os.Exit(1)
		}
		wg := &sync.WaitGroup{}
		statusBar := &ZCNStatus{wg: wg}
		wg.Add(1)
		zcncore.RegisterToMiners(clientWallet, statusBar)
		wg.Wait()
		if statusBar.success {
			fmt.Println("Wallet registered")
		} else {
			PrintError("Wallet registration failed. " + statusBar.errMsg)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerWalletCmd)
}
