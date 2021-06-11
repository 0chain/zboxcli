package service

import (
	"fmt"
	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/model"
	"log"
	"sync"
)

type UploadService struct {
	UploadConfig *model.MigrationConfig
	FileConfig   *model.SourceFileConfig
}

func NewUploadService(MigrationConfig *model.MigrationConfig, FileConfig *model.SourceFileConfig) *UploadService {
	return &UploadService{UploadConfig: MigrationConfig, FileConfig: FileConfig}
}

func (u *UploadService) UploadStreamToDStorage() error {
	log.Printf("uploading: from to remote '%s'", u.FileConfig.RemoteFilePath)
	allocationObj, err := sdk.GetAllocation(u.UploadConfig.AppConfig.AllocationID)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error fetching the allocation. %s", err.Error())
	}

	//todo: delete incomplete upload and re-upload
	if u.FileConfig.Incomplete {
		err = deleteIncompleteUpload(allocationObj, u.FileConfig.RemoteFilePath)
		if err != nil {
			return err
		}
	}

	wg := &sync.WaitGroup{}
	statusBar := &StatusBar{wg: wg}
	wg.Add(1)

	var attrs fileref.Attributes
	if u.UploadConfig.AppConfig.WhoPays != "" {
		var wp common.WhoPays
		if err = wp.Parse(u.UploadConfig.AppConfig.WhoPays); err != nil {
			return fmt.Errorf("error parssing who-pays value. %s", err.Error())
		}
		attrs.WhoPaysForReads = wp // set given value
	}
	if u.UploadConfig.AppConfig.Encrypt {
		//err = allocationObj.EncryptAndUploadFile(u.UploadConfig.LocalTempFilePath, u.UploadConfig.RemoteFilePath, attrs, statusBar)
		log.Println("encryption has not been implemented for direct upload")
	} else {
		input := sdk.UploadStreamInput{
			Size: u.FileConfig.SourceFileSize,
			Data: u.FileConfig.SourceFileReader,
			Type: u.FileConfig.SourceFileType,
		}
		err = allocationObj.UploadStream(&input, u.FileConfig.RemoteFilePath, attrs, statusBar)
	}

	if err != nil {
		return fmt.Errorf("upload failed. %s", err.Error())
	}
	wg.Wait()
	if !statusBar.success {
		log.Println("upload failed to complete.")
		return fmt.Errorf("upload failed. statusbar. success : %v", statusBar.success)
	}

	if u.UploadConfig.AppConfig.Commit {
		statusBar.wg.Add(1)
		err = commitMetaTxn(u.FileConfig.RemoteFilePath, "Upload", "", "", allocationObj, nil, statusBar)
		statusBar.wg.Wait()
	}

	return nil
}

func deleteIncompleteUpload(allocationObj *sdk.Allocation, filePath string) error {
	err := allocationObj.DeleteFile(filePath)
	if err != nil {
		log.Println("Delete failed:", err.Error())
		return err
	}

	return nil
}

func commitMetaTxn(path, crudOp, authTicket, lookupHash string, a *sdk.Allocation, fileMeta *sdk.ConsolidatedFileMeta, status *StatusBar) error {
	err := a.CommitMetaTransaction(path, crudOp, authTicket, lookupHash, fileMeta, status)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
