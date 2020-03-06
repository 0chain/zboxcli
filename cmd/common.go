package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"gopkg.in/cheggaaa/pb.v1"
)

const (
	ZCNStatusSuccess int = 0
	ZCNStatusError   int = 1
)

func (s *StatusBar) Started(allocationId, filePath string, op int, totalBytes int) {
	s.b = pb.StartNew(totalBytes)
	s.b.Set(0)
}
func (s *StatusBar) InProgress(allocationId, filePath string, op int, completedBytes int) {
	s.b.Set(completedBytes)
}

func (s *StatusBar) Completed(allocationId, filePath string, filename string, mimetype string, size int, op int) {
	// Not required
	// s.b.PrependElapsed()
	if s.b != nil {
		s.b.Finish()
	}
	s.success = true
	defer s.wg.Done()
	fmt.Println("Status completed callback. Type = " + mimetype + ". Name = " + filename)
}

func (s *StatusBar) Error(allocationID string, filePath string, op int, err error) {
	if s.b != nil {
		s.b.Finish()
	}
	s.success = false
	defer s.wg.Done()
	PrintError("Error in file operation." + err.Error())
}

type StatusBar struct {
	b       *pb.ProgressBar
	wg      *sync.WaitGroup
	success bool
}

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

func PrintError(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}

func PrintInfo(v ...interface{}) {
	fmt.Fprintln(os.Stdin, v...)
}

func commitMetaTxn(path, crudOp, authTicket, lookupHash string, a *sdk.Allocation, fileMeta *sdk.ConsolidatedFileMeta) {
	metaTxnData, err := a.CommitMetaTransaction(path, crudOp, authTicket, lookupHash, fileMeta)
	if err != nil {
		PrintError("Commit failed.", err)
		os.Exit(1)
	}
	metaDataBytes, _ := json.Marshal(metaTxnData.MetaData)
	PrintInfo("TxnID :", metaTxnData.TxnID)
	PrintInfo("MetaData :", string(metaDataBytes))
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}
