package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/0chain/gosdk/zboxcore/logger"
	"github.com/0chain/gosdk/zboxcore/sdk"
)

// M3u8Downloader download files from blobber's dir, and build them into a local m3u8 playlist
type M3u8Downloader struct {
	sync.RWMutex
	delay int

	localDir     string
	localPath    string
	remotePath   string
	authTicket   string
	allocationID string
	rxPay        bool

	allocationObj *sdk.Allocation

	lookupHash     string
	items          []MediaItem
	waitToDownload chan MediaItem
	playlist       *sdk.MediaPlaylist
	done           chan error
}

func createM3u8Downloader(localPath, remotePath, authTicket, allocationID, lookupHash string, rxPay bool, delay int) (*M3u8Downloader, error) {
	if len(remotePath) == 0 && (len(authTicket) == 0) {
		return nil, errors.New("Error: remotepath / authticket flag is missing")
	}

	if len(localPath) == 0 {
		return nil, errors.New("Error: localpath is missing")
	}

	dir := filepath.Dir(localPath)

	file, err := os.Create(localPath)

	if err != nil {
		return nil, err
	}

	downloader := &M3u8Downloader{
		localDir:       dir,
		localPath:      localPath,
		remotePath:     remotePath,
		authTicket:     authTicket,
		allocationID:   allocationID,
		rxPay:          rxPay,
		waitToDownload: make(chan MediaItem, 100),
		playlist:       sdk.NewMediaPlaylist(delay, dir, file),
		done:           make(chan error, 1),
	}

	if len(remotePath) > 0 {
		if len(allocationID) == 0 { // check if the flag "path" is set
			return nil, errors.New("Error: allocation flag is missing") // If not, we'll let the user know
		}

		allocationObj, err := sdk.GetAllocation(allocationID)

		if err != nil {
			return nil, fmt.Errorf("Error fetching the allocation: %s", err)
		}

		downloader.allocationObj = allocationObj

	} else if len(authTicket) > 0 {
		allocationObj, err := sdk.GetAllocationFromAuthTicket(authTicket)
		if err != nil {
			return nil, fmt.Errorf("Error fetching the allocation: %s", err)
		}

		downloader.allocationObj = allocationObj

		at := sdk.InitAuthTicket(authTicket)
		isDir, err := at.IsDir()
		if isDir && len(lookupHash) == 0 {
			lookupHash, err = at.GetLookupHash()
			if err != nil {
				return nil, fmt.Errorf("Error getting the lookuphash from authticket: %s", err)
			}

			downloader.lookupHash = lookupHash
		}
		if !isDir {
			return nil, fmt.Errorf("invalid operation. Auth ticket is not for a directory: %s", err)
		}

	}

	return downloader, nil
}

// Start start to download ,and build playlist
func (d *M3u8Downloader) Start() error {

	go d.autoDownload()
	go d.autoRefreshList()
	go d.playlist.Play()

	err := <-d.done

	return err
}

func (d *M3u8Downloader) addToDownload(item MediaItem) {
	d.waitToDownload <- item
}

func (d *M3u8Downloader) autoDownload() {
	for {
		item := <-d.waitToDownload
		//fmt.Println("download: ", item.Name)
		for i := 0; i < 3; i++ {
			if path, err := d.download(item); err == nil {
				d.playlist.Append(path)
				break
			}
		}
	}
}

func (d *M3u8Downloader) autoRefreshList() {
	for {
		list, err := d.getList()
		if err != nil {
			logger.Logger.Error("[m3u8]", err)
		} else {

			d.Lock()
			n := len(d.items)
			max := len(list)
			// there is new ts file
			if n < max {
				sort.Sort(SortedListResult(list))
				//	buf, _ := json.Marshal(list)
				//	fmt.Println(string(buf))
				for i := n; i < max; i++ {
					item := MediaItem{
						Name: list[i].Name,
						Path: list[i].Path,
					}
					d.items = append(d.items, item)

					//	fmt.Println("Added:", item.Name)
					d.addToDownload(item)
				}

			}
			d.Unlock()

		}

		time.Sleep(1 * time.Second)
	}
}

func (d *M3u8Downloader) download(item MediaItem) (string, error) {

	wg := &sync.WaitGroup{}
	statusBar := &StatusBar{wg: wg}
	wg.Add(1)

	localPath := filepath.Join(d.localDir, item.Name)
	remotePath := item.Path

	if len(d.remotePath) > 0 {

		err := d.allocationObj.DownloadFile(localPath, remotePath, statusBar)

		if err != nil {
			return "", err
		}

		wg.Wait()
	}

	//TODO: add download ts files from auth ticket
	// allocationObj, err = sdk.GetAllocationFromAuthTicket(authticket)
	// if err != nil {
	// 	PrintError("Error fetching the allocation", err)
	// 	os.Exit(1)
	// }
	// at := sdk.InitAuthTicket(authticket)
	// filename, err := at.GetFileName()
	// if err != nil {
	// 	PrintError("Error getting the filename from authticket", err)
	// 	os.Exit(1)
	// }
	// if len(lookuphash) == 0 {
	// 	lookuphash, err = at.GetLookupHash()
	// 	if err != nil {
	// 		PrintError("Error getting the lookuphash from authticket", err)
	// 		os.Exit(1)
	// 	}
	// }

	// return d.allocationObj.DownloadFromAuthTicket(localpath,
	// 	authticket, lookuphash, filename, rxPay, statusBar)

	//}

	return item.Name, nil
}

func (d *M3u8Downloader) getList() ([]*sdk.ListResult, error) {

	//get list from remoete allocations's path
	if len(d.remotePath) > 0 {

		ref, err := d.allocationObj.ListDir(d.remotePath)
		if err != nil {
			return nil, err
		}

		return ref.Children, nil

	}

	//get list from authticket
	ref, err := d.allocationObj.ListDirFromAuthTicket(d.authTicket, d.lookupHash)
	if err != nil {
		return nil, err
	}

	return ref.Children, nil

}

// SortedListResult sort files order by time
type SortedListResult []*sdk.ListResult

func (a SortedListResult) Len() int {
	return len(a)
}
func (a SortedListResult) Less(i, j int) bool {

	l := a[i]
	r := a[j]

	if len(l.Name) < len(r.Name) {
		return true
	}

	if len(l.Name) > len(r.Name) {
		return false
	}

	return l.Name < r.Name
}
func (a SortedListResult) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// MediaItem ts file
type MediaItem struct {
	Name string
	Path string
	// Type          string             `json:"type"`
	// Size          int64              `json:"size"`
	// Hash          string             `json:"hash,omitempty"`
	// MimeType      string             `json:"mimetype,omitempty"`
	// NumBlocks     int64              `json:"num_blocks"`
	// LookupHash    string             `json:"lookup_hash"`
	// EncryptionKey string             `json:"encryption_key"`
	// Attributes    fileref.Attributes `json:"attributes"`
}
