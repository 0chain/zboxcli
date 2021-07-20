package cmd

import (
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/model"
	"log"
	"strings"
	"testing"
)

type TestConfig struct {
}

func TestSingleUploadUsingReaderToDStorage(t *testing.T) {
	initConfig()
	migrationConfig := model.MigrationConfig{
		AppConfig: model.AppConfig{
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
	fileConfig := model.SourceFileConfig{
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
