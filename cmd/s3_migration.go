package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

var region string

// s3Cmd is the migrateFromS3 sub command to start the api server
var s3Cmd = &cobra.Command{
	Use:   "migrate-s3",
	Short: "migrate user data from S3 to dStorage",
	RunE:  migrateFromS3,
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.PersistentFlags().StringVarP(&region, "region", "r", "", "s3 region")
	s3Cmd.PersistentFlags().StringSlice("bucket", []string{}, "specific s3 buckets to use")
	s3Cmd.PersistentFlags().String("prefix", "", "s3 file prefix to use during migration")
	s3Cmd.Flags().Bool("delete-source", false, "pass this option to remove migrated files files from source (s3)")
	s3Cmd.PersistentFlags().Int("concurrency", 10, "number of concurrent files to process concurrently during migration")
	s3Cmd.PersistentFlags().String("allocation", "", "allocation ID for dStorage")
	//s3Cmd.PersistentFlags().String("thumbnailpath", "", "Local thumbnail path of file to upload")
	s3Cmd.Flags().Bool("encrypt", false, "pass this option to encrypt and upload the file")
	s3Cmd.Flags().Bool("commit", false, "pass this option to commit the metadata transaction")
	s3Cmd.MarkFlagRequired("allocation")
}

// Models
type AppConfig struct {
	AllocationID string
	Commit       bool
	Encrypt      bool
	WhoPays      string
}

// MigrationConfig contains configurations to upload files from S3
type MigrationConfig struct {
	AppConfig    AppConfig
	Region       string
	Buckets      []string
	Prefix       string
	Concurrency  int
	DeleteSource bool
}

func (m *MigrationConfig) ToString() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *MigrationConfig) DeepCopy() MigrationConfig {
	b, _ := json.Marshal(m)
	newConfig := MigrationConfig{}
	json.Unmarshal(b, &newConfig)

	return newConfig
}

// SourceFileConfig contains a single uploadable file's configs
type SourceFileConfig struct {
	SourceFileReader io.Reader
	SourceFileType   string
	SourceFileSize   int64
	RemoteFilePath   string
	Incomplete       bool
}

func (u *SourceFileConfig) ToString() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func (u *SourceFileConfig) DeepCopy() SourceFileConfig {
	b, _ := json.Marshal(u)
	q := SourceFileConfig{}
	json.Unmarshal(b, &q)

	return q
}

// UploadQueueItem contains a single queue item's configurations to upload files from S3
type UploadQueueItem struct {
	MigrationConfig MigrationConfig
	FileConfig      SourceFileConfig
	Bucket          string
	FileKey         string
	UploadQueue     *UploadQueue
}

func (u *UploadQueueItem) ToString() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func (u *UploadQueueItem) DeepCopy() UploadQueueItem {
	b, _ := json.Marshal(u)
	q := UploadQueueItem{}
	json.Unmarshal(b, &q)

	return q
}

// UploadQueue contains queue configurations to upload files from S3
type UploadQueue struct {
	WaitGroup        *sync.WaitGroup
	CurrentQueueSize int64
	MaxQueueSize     int64
}

func (u *UploadQueue) ToString() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func (u *UploadQueue) DeepCopy() UploadQueue {
	b, _ := json.Marshal(u)
	q := UploadQueue{}
	json.Unmarshal(b, &q)

	return q
}

// migration

func migrateFromS3(cmd *cobra.Command, args []string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting migration from s3")

	// get migration configurations
	migrationConfig, err := GetMigrationConfig(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// use a region specific s3 session to fetch files from s3
	sess := GetS3Session(migrationConfig.Region)

	// list all the buckets
	if len(migrationConfig.Buckets) == 0 {
		buckets, err := ListS3Buckets(sess)
		if err != nil {
			log.Println(err)
			return err
		}

		migrationConfig.Buckets = buckets
	}

	err = MigrateFromS3UsingStream(sess, migrationConfig)
	if err != nil {
		log.Println(err)
	}

	log.Println("migration complete")
	return nil

}

func ListS3Buckets(sess *session.Session) ([]string, error) {
	// Create S3 service client
	svc := s3.New(sess)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Println("Unable to list buckets, %v" + err.Error())
		return nil, err
	}

	buckets := make([]string, 0)
	for _, b := range result.Buckets {
		buckets = append(buckets, aws.StringValue(b.Name))
	}

	return buckets, nil
}

// MigrateFromS3UsingStream steams all the data in the specified buckets to dStorage
func MigrateFromS3UsingStream(sess *session.Session, migrationConfig *MigrationConfig) error {
	// Create S3 service client
	svc := s3.New(sess)

	// use existing file list to exclude files that already exists in remote directory from being processed
	existingFIleList, err := getExistingFileList(migrationConfig)
	if err != nil {
		log.Println(err)
		return err
	}

	wg := sync.WaitGroup{}
	uploadQueue := UploadQueue{
		WaitGroup:        &wg,
		CurrentQueueSize: 0,
		MaxQueueSize:     int64(migrationConfig.Concurrency),
	}

	for _, thisBucket := range migrationConfig.Buckets {
		err := svc.ListObjectsV2Pages(&s3.ListObjectsV2Input{Bucket: aws.String(thisBucket), Prefix: aws.String(migrationConfig.Prefix)}, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
			for _, item := range page.Contents {
				if *item.Size == 0 {
					continue
				}

				remoteFilePath := fmt.Sprintf("/%s/%s", thisBucket, *item.Key)
				incomplete := false
				if fileSize, ok := existingFIleList[remoteFilePath]; ok {
					if fileSize != *item.Size {
						//log.Println("migration was incomplete for this file")
						incomplete = true
					} else {
						continue
					}
				}

				uploadQueueItem := UploadQueueItem{
					MigrationConfig: migrationConfig.DeepCopy(),
					FileConfig:      SourceFileConfig{RemoteFilePath: remoteFilePath, Incomplete: incomplete},
					Bucket:          thisBucket,
					FileKey:         *item.Key,
					UploadQueue:     &uploadQueue,
				}
				enQueueFileToBeMigrated(svc, &uploadQueueItem)
			}
			return true
		})
		if err != nil {
			log.Printf("Unable to list items in bucket %q, %v", thisBucket, err)
			return err
		}
	}

	wg.Wait()

	return nil
}

//getExistingFileList list existing files (with size) from dStorage
func getExistingFileList(uploadConfig *MigrationConfig) (map[string]int64, error) {
	allocationObj, err := sdk.GetAllocation(uploadConfig.AppConfig.AllocationID)
	if err != nil {
		PrintError("Error fetching the allocation", err)
		return nil, err
	}
	// Create filter
	filter := []string{".DS_Store", ".git"}
	exclMap := make(map[string]int)
	for idx, path := range filter {
		exclMap[strings.TrimRight(path, "/")] = idx
	}

	remoteFiles, err := allocationObj.GetRemoteFileMap(exclMap)
	if err != nil {
		PrintError("Error getting remote files.", err)
		return nil, err
	}

	fileList := map[string]int64{}

	for remoteFileName, remoteFileValue := range remoteFiles {
		fileList[remoteFileName] = remoteFileValue.ActualSize
	}

	return fileList, nil
}

func enQueueFileToBeMigrated(svc *s3.S3, uploadQueueItem *UploadQueueItem) {
	queue := uploadQueueItem.UploadQueue
	for {
		if queue.MaxQueueSize == 0 || queue.CurrentQueueSize < queue.MaxQueueSize {
			queue.CurrentQueueSize = queue.CurrentQueueSize + 1
			queue.WaitGroup.Add(1)
			go func() {
				//log.Println(uploadQueueItem.ToString())
				err := sendToStorage(svc, uploadQueueItem)
				if err != nil {
					PrintError("upload to storage failed", err)
				}
			}()
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}

//sendToStorage takes a single file from the bucket and upload it as a stream to dStorage
func sendToStorage(svc *s3.S3, uploadQueueItem *UploadQueueItem) error {
	//fmt.Printf("Migrating s3://%s/%s..\n", uploadQueueItem.Bucket, uploadQueueItem.FileKey)
	queue := uploadQueueItem.UploadQueue
	defer func() {
		//log.Println(fmt.Sprintf("Migration done for s3://%s/%s..\n", uploadQueueItem.Bucket, uploadQueueItem.FileKey))
		queue.CurrentQueueSize = queue.CurrentQueueSize - 1
		queue.WaitGroup.Done()
	}()

	out, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(uploadQueueItem.Bucket),
		Key:    aws.String(uploadQueueItem.FileKey),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	uploadQueueItem.FileConfig.SourceFileReader = out.Body
	uploadQueueItem.FileConfig.SourceFileSize = *out.ContentLength
	uploadQueueItem.FileConfig.SourceFileType = *out.ContentType

	uploadService := NewUploadService(&uploadQueueItem.MigrationConfig, &uploadQueueItem.FileConfig)
	err = uploadService.UploadStreamToDStorage()
	if err != nil {
		log.Println("upload error:", err)
		return err
	}

	if uploadQueueItem.MigrationConfig.DeleteSource {
		_, err = svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(uploadQueueItem.Bucket),
			Key:    aws.String(uploadQueueItem.FileKey),
		})
		if err != nil {
			log.Println(fmt.Sprintf("err removing source file  s3://%s/%s: %v", uploadQueueItem.Bucket, uploadQueueItem.FileKey, err))
		}
	}

	return nil
}

// Config

//GetMigrationConfig returns user defined configs to migrate data from s3
func GetMigrationConfig(cmd *cobra.Command) (*MigrationConfig, error) {
	config := MigrationConfig{}

	fflags := cmd.Flags()                 // fflags is a *flag.FlagSet
	if fflags.Changed("bucket") == true { // set bucket
		buckets, err := fflags.GetStringSlice("bucket")
		if err != nil {
			log.Println("bucket list not provided, backing up all buckets in this region:")

		} else {
			//log.Println("buckets from cmd:", buckets)
			config.Buckets = buckets
		}
	}

	if fflags.Changed("region") == true { // check if the flag "region" is set
		config.Region, _ = fflags.GetString("region")
	}
	config.Prefix, _ = fflags.GetString("prefix")
	config.Concurrency, _ = fflags.GetInt("concurrency")
	config.DeleteSource, _ = cmd.Flags().GetBool("delete-source")

	if fflags.Changed("allocation") == false { // check if the flag "allocation" is set
		return nil, fmt.Errorf("allocation flag is missing")
	}

	config.AppConfig.AllocationID = cmd.Flag("allocation").Value.String()

	config.AppConfig.Encrypt, _ = cmd.Flags().GetBool("encrypt")
	config.AppConfig.Commit, _ = cmd.Flags().GetBool("commit")
	if fflags.Changed("attr-who-pays-for-reads") {
		wps, err := fflags.GetString("attr-who-pays-for-reads")
		if err != nil {
			return nil, fmt.Errorf("getting 'attr-who-pays-for-reads' flag: %v", err)
		}
		config.AppConfig.WhoPays = wps
	}

	return &config, nil
}

// GetS3Session gets a reusable session to fetch data from s3
func GetS3Session(region string) *session.Session {
	if region == "" {
		log.Println("no region defined. using \"us-east-2\"")
		region = "us-east-2"
	}

	configP, _ := session.NewSession(&aws.Config{Region: aws.String(region)})

	return configP
}

// Upload service
type UploadService struct {
	UploadConfig *MigrationConfig
	FileConfig   *SourceFileConfig
}

func NewUploadService(MigrationConfig *MigrationConfig, FileConfig *SourceFileConfig) *UploadService {
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
		commitMetaTxn(u.FileConfig.RemoteFilePath, "Upload", "", "", allocationObj, nil, statusBar)
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
