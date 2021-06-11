package service

import (
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"sync"
)

var (
	allocUnderRepair bool
)

func PrintError(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}

func (s *StatusBar) Started(allocationId, filePath string, op int, totalBytes int) {
	s.b = pb.StartNew(totalBytes)
	s.b.Set(0)
}
func (s *StatusBar) InProgress(allocationId, filePath string, op int, completedBytes int, data []byte) {
	s.b.Set(completedBytes)
}

func (s *StatusBar) Completed(allocationId, filePath string, filename string, mimetype string, size int, op int) {
	if s.b != nil {
		s.b.Finish()
	}
	s.success = true
	if !allocUnderRepair {
		defer s.wg.Done()
	}
	fmt.Println("Status completed callback. Type = " + mimetype + ". Name = " + filename)
}

func (s *StatusBar) Error(allocationID string, filePath string, op int, err error) {
	if s.b != nil {
		s.b.Finish()
	}
	s.success = false
	if !allocUnderRepair {
		defer s.wg.Done()
	}

	var errDetail interface{} = "Unknown Error"
	if err != nil {
		errDetail = err.Error()
	}

	PrintError("Error in file operation:", errDetail)
}

func (s *StatusBar) CommitMetaCompleted(request, response string, err error) {
	defer s.wg.Done()
	if err != nil {
		s.success = false
		PrintError("Error in commitMetaTransaction." + err.Error())
	} else {
		s.success = true
		fmt.Println("Commit Metadata successful, Response :", response)
	}
}

func (s *StatusBar) RepairCompleted(filesRepaired int) {
	defer s.wg.Done()
	allocUnderRepair = false
	fmt.Println("Repair file completed, Total files repaired: ", filesRepaired)
}

type StatusBar struct {
	b       *pb.ProgressBar
	wg      *sync.WaitGroup
	success bool
}
