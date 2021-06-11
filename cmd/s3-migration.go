package cmd

import (
	"fmt"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/zboxcli/config"
	"github.com/0chain/zboxcli/model"
	"github.com/0chain/zboxcli/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
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

func migrateFromS3(cmd *cobra.Command, args []string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting migration from s3")

	// get migration configurations
	migrationConfig, err := config.GetMigrationConfig(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// use a region specific s3 session to fetch files from s3
	sess := config.GetS3Session(migrationConfig.Region)

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
func MigrateFromS3UsingStream(sess *session.Session, migrationConfig *model.MigrationConfig) error {
	// Create S3 service client
	svc := s3.New(sess)

	// use existing file list to exclude files that already exists in remote directory from being processed
	existingFIleList, err := getExistingFileList(migrationConfig)
	if err != nil {
		log.Println(err)
		return err
	}

	wg := sync.WaitGroup{}
	uploadQueue := model.UploadQueue{
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

				uploadQueueItem := model.UploadQueueItem{
					MigrationConfig: migrationConfig.DeepCopy(),
					FileConfig:      model.SourceFileConfig{RemoteFilePath: remoteFilePath, Incomplete: incomplete},
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
func getExistingFileList(uploadConfig *model.MigrationConfig) (map[string]int64, error) {
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

func enQueueFileToBeMigrated(svc *s3.S3, uploadQueueItem *model.UploadQueueItem) {
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
func sendToStorage(svc *s3.S3, uploadQueueItem *model.UploadQueueItem) error {
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

	uploadService := service.NewUploadService(&uploadQueueItem.MigrationConfig, &uploadQueueItem.FileConfig)
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
