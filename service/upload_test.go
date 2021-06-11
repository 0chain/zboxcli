package service

import (
	"encoding/json"
	"fmt"
	"github.com/0chain/gosdk/core/zcncrypto"
	"github.com/0chain/gosdk/zboxcore/blockchain"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
	"github.com/0chain/zboxcli/model"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

type TestConfig struct {

}

func TestSingleUploadUsingReaderToDStorage(t *testing.T) {
	initConfig()
	migrationConfig := model.MigrationConfig{
		AppConfig:    model.AppConfig{
			AllocationID: "1077eca7599ba5550d267dd7f7c7dec9b4f2e11dc57921c98576b61b94acd732",
			Commit:       false,
			Encrypt:      false,
			WhoPays:      "",
		},
		//Region:       "us-east-2",
		//Buckets:      ,
		//Prefix:       "",
		//Concurrency:  0,
		DeleteSource: false,
	}

	filePath := "/test/file_1.txt"
	fileContent := "content of file_1.txt"
	fileConfig:= model.SourceFileConfig{
		SourceFileReader: strings.NewReader(fileContent),
		SourceFileType:   "plain/text",
		SourceFileSize:   int64(len([]byte(fileContent))),
		RemoteFilePath:   filePath,
		Incomplete:       false,
	}

	uploadService := NewUploadService(&migrationConfig, &fileConfig)

	err := uploadService.UploadStreamToDStorage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("upload complete, cleaning up test data:")

	err = cleanupUploadedFiles(filePath)
	if err != nil {
		t.Fatal(err)
	}
}

func cleanupUploadedFiles(filePath string) error {
	allocationID := "1077eca7599ba5550d267dd7f7c7dec9b4f2e11dc57921c98576b61b94acd732"
	allocationObj, err := sdk.GetAllocation(allocationID)
	if err != nil {
		return err
	}

	err = allocationObj.DeleteFile(filePath)
	if err != nil {
		log.Println("Delete failed:", err.Error())
		return err
	}

	return nil
}


//================xx setup sdk   xx=========================//
var cfgFile string
var networkFile string
var walletFile string
var walletClientID string
var walletClientKey string
var cDir string
var bVerbose bool

var preferredBlobbers []string
var clientConfig string
var minSubmit int
var minCfm int
var CfmChainLength int

var clientWallet *zcncrypto.Wallet

const (
	ZCNStatusSuccess int = 0
	ZCNStatusError   int = 1
)

type ZCNStatus struct {
	walletString string
	wg           *sync.WaitGroup
	success      bool
	errMsg       string
}

func (zcn *ZCNStatus) OnWalletCreateComplete(status int, wallet string, err string) {
	defer zcn.wg.Done()
	if status == ZCNStatusError {
		zcn.success = false
		zcn.errMsg = err
		zcn.walletString = ""
		return
	}
	zcn.success = true
	zcn.errMsg = ""
	zcn.walletString = wallet
	return
}

func getConfigDir() string {
	if cDir != "" {
		return cDir
	}
	var configDir string
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configDir = home + string(os.PathSeparator) + ".zcn"
	return configDir
}

func initConfig() {
	nodeConfig := viper.New()
	networkConfig := viper.New()
	var configDir string
	if cDir != "" {
		configDir = cDir
	} else {
		configDir = getConfigDir()
	}
	// Search config in home directory with name ".cobra" (without extension).
	nodeConfig.AddConfigPath(configDir)
	if &cfgFile != nil && len(cfgFile) > 0 {
		nodeConfig.SetConfigFile(configDir + "/" + cfgFile)
	} else {
		nodeConfig.SetConfigFile(configDir + "/" + "config.yaml")
	}

	networkConfig.AddConfigPath(configDir)
	if &networkFile != nil && len(networkFile) > 0 {
		networkConfig.SetConfigFile(configDir + "/" + networkFile)
	} else {
		networkConfig.SetConfigFile(configDir + "/" + "network.yaml")
	}

	if err := nodeConfig.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	blockWorker := nodeConfig.GetString("block_worker")
	preferredBlobbers = nodeConfig.GetStringSlice("preferred_blobbers")
	signScheme := nodeConfig.GetString("signature_scheme")
	chainID := nodeConfig.GetString("chain_id")
	minSubmit = nodeConfig.GetInt("min_submit")
	minCfm = nodeConfig.GetInt("min_confirmation")
	CfmChainLength = nodeConfig.GetInt("confirmation_chain_length")
	// additional settings depending network latency
	maxTxnQuery := nodeConfig.GetInt("max_txn_query")
	querySleepTime := nodeConfig.GetInt("query_sleep_time")

	//TODO: move the private key storage to the keychain or secure storage

	//set the log file
	zcncore.SetLogFile("cmdlog.log", bVerbose)
	sdk.SetLogFile("cmdlog.log", bVerbose)

	err := zcncore.InitZCNSDK(blockWorker, signScheme,
		zcncore.WithChainID(chainID),
		zcncore.WithMinSubmit(minSubmit),
		zcncore.WithMinConfirmation(minCfm),
		zcncore.WithConfirmationChainLength(CfmChainLength))
	if err != nil {
		fmt.Println("Error initializing core SDK.", err)
		os.Exit(1)
	}

	if err := networkConfig.ReadInConfig(); err == nil {
		miners := networkConfig.GetStringSlice("miners")
		sharders := networkConfig.GetStringSlice("sharders")
		if len(miners) > 0 && len(sharders) > 0 {
			zcncore.SetNetwork(miners, sharders)
		}
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
	err = sdk.InitStorageSDK(clientConfig, blockWorker, chainID, signScheme, preferredBlobbers)
	if err != nil {
		fmt.Println("Error in sdk init", err)
		os.Exit(1)
	}

	// additional settings depending network latency
	if maxTxnQuery > 0 {
		blockchain.SetMaxTxnQuery(maxTxnQuery)
	}
	if querySleepTime > 0 {
		blockchain.SetQuerySleepTime(querySleepTime)
	}

	if err := networkConfig.ReadInConfig(); err == nil {
		miners := networkConfig.GetStringSlice("miners")
		sharders := networkConfig.GetStringSlice("sharders")
		if len(miners) > 0 && len(sharders) > 0 {
			sdk.SetNetwork(miners, sharders)
		}
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
