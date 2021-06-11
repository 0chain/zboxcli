package model

import (
	"encoding/json"
	"io"
	"sync"
)

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
