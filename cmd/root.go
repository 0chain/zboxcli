package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/0chain/gosdk/core/conf"
	"github.com/0chain/gosdk/core/logger"
	"github.com/0chain/gosdk/core/sys"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/0chain/gosdk/zboxcore/blockchain"
	"github.com/0chain/gosdk/zboxcore/client"
	"github.com/0chain/zboxcli/util"

	"github.com/0chain/gosdk/core/zcncrypto"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
)

var cfgFile string
var networkFile string
var walletFile string
var walletFilePath string
var walletClientID string
var walletClientKey string
var cDir string
var nonce int64
var gTxnFee float64
var bSilent bool
var allocUnderRepair bool
var gConfig *viper.Viper

var walletJSON string

var rootCmd = &cobra.Command{
	Use:   "zbox",
	Short: "zbox is a decentralized storage application written on the 0Chain platform",
	Long: `zbox is a decentralized storage application written on the 0Chain platform.
			Complete documentation is available at https://docs.zus.network/guides/zbox-cli`,
}

var clientWallet *zcncrypto.Wallet

func init() {

	InstallDLLs()

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	rootCmd.PersistentFlags().StringVar(&networkFile, "network", "", "network file to overwrite the network details (if required, default is network.yaml)")
	rootCmd.PersistentFlags().StringVar(&walletFile, "wallet", "", "wallet file (default is wallet.json)")
	rootCmd.PersistentFlags().StringVar(&walletClientID, "wallet_client_id", "", "wallet client_id")
	rootCmd.PersistentFlags().StringVar(&walletClientKey, "wallet_client_key", "", "wallet client_key")
	rootCmd.PersistentFlags().Int64Var(&nonce, "withNonce", 0, "nonce that will be used in transaction (default is 0)")
	rootCmd.PersistentFlags().StringVar(&cDir, "configDir", "", "configuration directory (default is $HOME/.zcn)")
	rootCmd.PersistentFlags().BoolVar(&bSilent, "silent", false, "(default false) Do not show interactive sdk logs (shown by default)")
	rootCmd.PersistentFlags().Float64Var(&gTxnFee, "fee", 0, "transaction fee for the given transaction (if unset, it will be set to blockchain min fee)")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = "config.yaml"
	}

	var (
		configDir = util.SetConfigDir(cDir)
		cfg       conf.Config
		err       error
	)

	cfg, err = conf.LoadConfigFile(filepath.Join(configDir, cfgFile))
	if err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	gConfig = cfg.V

	if networkFile == "" {
		networkFile = "network.yaml"
	}
	network, _ := conf.LoadNetworkFile(filepath.Join(configDir, networkFile))

	// syncing loggers
	logger.SyncLoggers([]*logger.Logger{zcncore.GetLogger(), sdk.GetLogger()})

	// set the log file
	zcncore.SetLogFile("cmdlog.log", !bSilent)
	sdk.SetLogFile("cmdlog.log", !bSilent)
	sdk.SetMinSubmit(cfg.MinSubmit)

	if network.IsValid() {
		zcncore.SetNetwork(network.Miners, network.Sharders)
		conf.InitChainNetwork(&conf.Network{
			Miners:   network.Miners,
			Sharders: network.Sharders,
		})
	}

	clientWallet = loadWallet(configDir)
	if cfg.ZauthServer != "" {
		sys.SetAuthorize(zcncore.ZauthSignTxn(cfg.ZauthServer))
		client.SetClient(clientWallet, cfg.SignatureScheme, getTxnFee())
		if err := zcncore.SetWallet(*clientWallet, clientWallet.IsSplit); err != nil {
			fmt.Println("Error setting wallet", err)
			os.Exit(1)
		}
	}

	err = zcncore.InitZCNSDK(cfg.BlockWorker, cfg.SignatureScheme,
		zcncore.WithChainID(cfg.ChainID),
		zcncore.WithMinSubmit(cfg.MinSubmit),
		zcncore.WithMinConfirmation(cfg.MinConfirmation),
		zcncore.WithConfirmationChainLength(cfg.ConfirmationChainLength),
		zcncore.WithIsSplitWallet(clientWallet.IsSplit))
	if err != nil {
		fmt.Println("Error initializing core SDK.", err)
		os.Exit(1)
	}

	// Initialize sys functions for auth signing
	sys.AuthCommon = zcncore.ZauthAuthCommon(cfg.ZauthServer)
	sys.SignWithAuth = zcncore.ZauthSignMsg(cfg.ZauthServer)

	//init the storage sdk with the known miners, sharders and client wallet info
	if err = sdk.InitStorageSDK(
		walletJSON,
		cfg.BlockWorker,
		cfg.ChainID,
		cfg.SignatureScheme,
		cfg.PreferredBlobbers,
		nonce,
		zcncore.ConvertToValue(gTxnFee),
	); err != nil {
		fmt.Println("Error in sdk init", err)
		os.Exit(1)
	}

	// additional settings depending network latency
	blockchain.SetMaxTxnQuery(cfg.MaxTxnQuery)
	blockchain.SetQuerySleepTime(cfg.QuerySleepTime)

	conf.InitClientConfig(&cfg)

	if network.IsValid() {
		sdk.SetNetwork(network.Miners, network.Sharders)
	}

	sdk.SetNumBlockDownloads(10)
}

func loadWallet(configDir string) *zcncrypto.Wallet {
	// is freshly created wallet?
	//var fresh bool
	wallet := &zcncrypto.Wallet{}
	if (&walletClientID != nil) && (len(walletClientID) > 0) && (&walletClientKey != nil) && (len(walletClientKey) > 0) {
		wallet.ClientID = walletClientID
		wallet.ClientKey = walletClientKey
		var clientBytes []byte

		clientBytes, err := json.Marshal(wallet)
		walletJSON = string(clientBytes)
		if err != nil {
			fmt.Println("Invalid wallet data passed:" + walletClientID + " " + walletClientKey)
			os.Exit(1)
		}
		return wallet
	}

	if &walletFile != nil && len(walletFile) > 0 {
		if filepath.IsAbs(walletFile) {
			walletFilePath = walletFile
		} else {
			walletFilePath = configDir + string(os.PathSeparator) + walletFile
		}
	} else {
		walletFilePath = configDir + string(os.PathSeparator) + "wallet.json"
	}

	if _, err := os.Stat(walletFilePath); os.IsNotExist(err) {
		wallet, err := zcncore.CreateWalletOffline()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("ZCN wallet created")
		walletJSON = wallet
		file, err := os.Create(walletFilePath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()
		fmt.Fprint(file, walletJSON)
	} else {
		f, err := os.Open(walletFilePath)
		if err != nil {
			fmt.Println("Error opening the wallet", err)
			os.Exit(1)
		}
		clientBytes, err := io.ReadAll(f)
		if err != nil {
			fmt.Println("Error reading the wallet", err)
			os.Exit(1)
		}
		walletJSON = string(clientBytes)
	}
	//minerjson, _ := json.Marshal(miners)
	//sharderjson, _ := json.Marshal(sharders)
	err := json.Unmarshal([]byte(walletJSON), wallet)
	if err != nil {
		fmt.Println("Invalid wallet at path:" + walletFilePath)
		os.Exit(1)
	}
	return wallet
}

func getTxnFee() uint64 {
	return zcncore.ConvertToValue(gTxnFee)
}
