package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/0chain/gosdk/core/conf"
	"github.com/0chain/gosdk/core/logger"
	"github.com/0chain/gosdk/core/transaction"

	"github.com/spf13/cobra"

	"github.com/0chain/gosdk/zboxcore/blockchain"
	"github.com/0chain/zboxcli/util"

	"github.com/0chain/gosdk/core/zcncrypto"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
)

var cfgFile string
var networkFile string
var walletFile string
var walletClientID string
var walletClientKey string
var cDir string
var bSilent bool
var allocUnderRepair bool

var clientConfig string

var rootCmd = &cobra.Command{
	Use:   "zbox",
	Short: "zbox is a decentralized storage application written on the 0Chain platform",
	Long: `zbox is a decentralized storage application written on the 0Chain platform.
			Complete documentation is available at https://docs.0chain.net/0chain/`,
}

var clientWallet *zcncrypto.Wallet

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	rootCmd.PersistentFlags().StringVar(&networkFile, "network", "", "network file to overwrite the network details (if required, default is network.yaml)")
	rootCmd.PersistentFlags().StringVar(&walletFile, "wallet", "", "wallet file (default is wallet.json)")
	rootCmd.PersistentFlags().StringVar(&walletClientID, "wallet_client_id", "", "wallet client_id")
	rootCmd.PersistentFlags().StringVar(&walletClientKey, "wallet_client_key", "", "wallet client_key")
	rootCmd.PersistentFlags().StringVar(&cDir, "configDir", "", "configuration directory (default is $HOME/.zcn)")
	rootCmd.PersistentFlags().BoolVar(&bSilent, "silent", false, "Do not show interactive sdk logs (shown by default)")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {

	var configDir string
	if cDir != "" {
		configDir = cDir
	} else {
		configDir = util.GetConfigDir()
	}

	if cfgFile == "" {
		cfgFile = "config.yaml"
	}
	cfg, err := conf.LoadConfigFile(filepath.Join(configDir, cfgFile))
	if err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	transaction.SetConfig(&cfg)

	if networkFile == "" {
		networkFile = "network.yaml"
	}
	network, _ := conf.LoadNetworkFile(filepath.Join(configDir, networkFile))

	// syncing loggers
	logger.SyncLoggers([]*logger.Logger{zcncore.GetLogger(), sdk.GetLogger()})

	// set the log file
	zcncore.SetLogFile("cmdlog.log", !bSilent)
	sdk.SetLogFile("cmdlog.log", !bSilent)

	err = zcncore.InitZCNSDK(cfg.BlockWorker, cfg.SignatureScheme,
		zcncore.WithChainID(cfg.ChainID),
		zcncore.WithMinSubmit(cfg.MinSubmit),
		zcncore.WithMinConfirmation(cfg.MinConfirmation),
		zcncore.WithConfirmationChainLength(cfg.ConfirmationChainLength))
	if err != nil {
		fmt.Println("Error initializing core SDK.", err)
		os.Exit(1)
	}

	if network.IsValid() {
		zcncore.SetNetwork(network.Miners, network.Sharders)
	}

	// is freshly created wallet?
	var fresh bool

	wallet := &zcncrypto.Wallet{}
	if (&walletClientID != nil) && (len(walletClientID) > 0) && (&walletClientKey != nil) && (len(walletClientKey) > 0) {
		wallet.ClientID = walletClientID
		wallet.ClientKey = walletClientKey
		var clientBytes []byte

		clientBytes, err = json.Marshal(wallet)
		clientConfig = string(clientBytes)
		if err != nil {
			fmt.Println("Invalid wallet data passed:" + walletClientID + " " + walletClientKey)
			os.Exit(1)
		}
		clientWallet = wallet
		fresh = false
	} else {
		var walletFilePath string
		if &walletFile != nil && len(walletFile) > 0 {
			if filepath.IsAbs(walletFile) {
				walletFilePath = walletFile
			} else {
				walletFilePath = configDir + string(os.PathSeparator) + walletFile
			}
		} else {
			walletFilePath = configDir + string(os.PathSeparator) + "wallet.json"
		}

		if _, err = os.Stat(walletFilePath); os.IsNotExist(err) {
			wg := &sync.WaitGroup{}
			statusBar := &ZCNStatus{wg: wg}
			wg.Add(1)
			err = zcncore.CreateWallet(statusBar)
			if err == nil {
				wg.Wait()
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if len(statusBar.walletString) == 0 || !statusBar.success {
				fmt.Println("Error creating the wallet." + statusBar.errMsg)
				os.Exit(1)
			}
			fmt.Println("ZCN wallet created")
			clientConfig = string(statusBar.walletString)
			file, err := os.Create(walletFilePath)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			defer file.Close()
			fmt.Fprintf(file, clientConfig)

			fresh = true
		} else {
			f, err := os.Open(walletFilePath)
			if err != nil {
				fmt.Println("Error opening the wallet", err)
				os.Exit(1)
			}
			clientBytes, err := ioutil.ReadAll(f)
			if err != nil {
				fmt.Println("Error reading the wallet", err)
				os.Exit(1)
			}
			clientConfig = string(clientBytes)
		}
		//minerjson, _ := json.Marshal(miners)
		//sharderjson, _ := json.Marshal(sharders)
		err = json.Unmarshal([]byte(clientConfig), wallet)
		clientWallet = wallet
		if err != nil {
			fmt.Println("Invalid wallet at path:" + walletFilePath)
			os.Exit(1)
		}
	}

	//init the storage sdk with the known miners, sharders and client wallet info
	err = sdk.InitStorageSDK(clientConfig, cfg.BlockWorker, cfg.ChainID, cfg.SignatureScheme, cfg.PreferredBlobbers)
	if err != nil {
		fmt.Println("Error in sdk init", err)
		os.Exit(1)
	}

	// additional settings depending network latency
	blockchain.SetMaxTxnQuery(cfg.MaxTxnQuery)
	blockchain.SetQuerySleepTime(cfg.QuerySleepTime)

	if network.IsValid() {
		sdk.SetNetwork(network.Miners, network.Sharders)
	}

	sdk.SetNumBlockDownloads(10)

	if fresh {
		fmt.Println("Creating related read pool for storage smart-contract...")
		if err = sdk.CreateReadPool(); err != nil {
			fmt.Printf("Failed to create read pool: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Read pool created successfully")
	}
}
