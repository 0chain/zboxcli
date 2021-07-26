package integrationtest

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	dirPath      = "/home/zboxcli"
	remotePath   = "/"
	fileName     = "1.txt"
	allocationID = "bb1c56abdb92d033a200d17c1d625fd83f37358fab9b4ca2a73a2640e915d4a0"
	publicKey    = "13d46abf7f400bba6a013cfed6ff500d9b90080c67f1183a4ddde3e3c767fa202fa4efed01841c75291c6e9ff0b00c2d2fca033e3cc2199038db74c25e204b10"
)

//Registers the wallet with the blockchain
func Test_Register(t *testing.T) {
	cmd := exec.Command("./zbox", "register") // or whatever the program is
	cmd.Dir = dirPath                         // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
	}
	require.Contains(t, "Wallet registered\n", string(out))
}

//Show active blobbers in storage SC.
func Test_ListBlobbers(t *testing.T) {
	cmd := exec.Command("./zbox", "ls-blobbers") // or whatever the program is
	cmd.Dir = dirPath                            // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
	require.Contains(t, string(out), "d64ff8605e16d89b1a78fc6d90c7657daa053e4f95c554f00527258494f5647c")

}

//List allocations for the client
func Test_ListAllocations(t *testing.T) {
	cmd := exec.Command("./zbox", "listallocations") // or whatever the program is
	cmd.Dir = dirPath                                // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
	require.Contains(t, string(out), allocationID)
}

//list files from blobbers
func Test_ListFileFromBlobber(t *testing.T) {
	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/")
	cmd := exec.Command("./zbox", "list", "--remotepath", remotePath, "--allocation", allocationID) // or whatever the program is
	cmd.Dir = dirPath                                                                               // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
	DeleteFileFromBlobber("/1.txt", allocationID)
	require.Contains(t, string(out), "/1.txt")
}
func UploadFile(filepath string, allocationID string, remotePath string) {
	cmd := exec.Command("./zbox", "upload", "--localpath", filepath, "--remotepath", remotePath, "--allocation", allocationID) // or whatever the program is
	cmd.Dir = dirPath                                                                                                          // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
}
func DeleteLocalFile(filepath string) {
	cmd := exec.Command("rm", filepath) // or whatever the program is
	cmd.Dir = dirPath                   // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
}

func ShareFile(remotePath string, allocationID string) string {
	cmd := exec.Command("./zbox", "share", "--remotepath", remotePath, "--allocation", allocationID) // or whatever the program is
	cmd.Dir = dirPath                                                                                // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return ""
	}
	token := strings.Split(string(out), ":")
	return token[1]
}
func LockTokensIntoReadPool(allocationID string, token string, duration string) {
	cmd := exec.Command("./zbox", "rp-lock", "--allocation", allocationID, "--duration", duration, "--tokens", token) // or whatever the program is
	cmd.Dir = dirPath                                                                                                 // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)

		log.Fatal(err)

	} else {
		fmt.Printf("%s", out)
	}
}
func GetAuth(remotePath string, allocationID string) {
	cmd := exec.Command("./zbox", "share", "--remotepath", remotePath, "--allocation", allocationID) // or whatever the program is
	cmd.Dir = dirPath                                                                                // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
}
func DeleteFileFromBlobber(remotePath string, allocationID string) {
	cmd := exec.Command("./zbox", "delete", "--remotepath", remotePath, "--allocation", allocationID) // or whatever the program is
	cmd.Dir = dirPath                                                                                 // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s", out)
	}
}

//list all files from blobbers
func Test_ListAllFileFromBlobber(t *testing.T) {
	// DeleteFileFromBlobber("/destname.txt", allocationID)

	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/")

	cmd := exec.Command("./zbox", "list-all", "--allocation", allocationID)
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
	DeleteFileFromBlobber("/1.txt", allocationID)

	require.Contains(t, string(out), "/1.txt")
}

func Test_CreatesANewAllocation(t *testing.T) {

	cmd := exec.Command("./zbox", "newallocation", "--lock", "0.01")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}

	require.Contains(t, string(out), "Allocation created")
}
func Test_RenameAnObject(t *testing.T) {
	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/")

	cmd := exec.Command("./zbox", "rename", "--allocation", allocationID, "--remotepath", "/1.txt", "--destname", "destname.txt")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
	DeleteFileFromBlobber("/destname.txt", allocationID)

	require.Contains(t, string(out), "/1.txt renamed")
}
func Test_RenameAnObjectFail(t *testing.T) {

	cmd := exec.Command("./zbox", "rename", "--allocation", allocationID, "--remotepath", "/1.txt", "--destname", "destname.txt")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}

	require.Contains(t, string(out), "Can not find path")
}
func Test_ShareFilesFromBlobbers(t *testing.T) {
	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/")

	cmd := exec.Command("./zbox", "share", "--allocation", allocationID, "--remotepath", "/1.txt")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
	DeleteFileFromBlobber("/1.txt", allocationID)

	require.Contains(t, string(out), "Auth token")
}

func Test_ShareFilesFromBlobbersFail(t *testing.T) {

	cmd := exec.Command("./zbox", "share", "--allocation", allocationID, "--remotepath", "/1.txt")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}

	require.Contains(t, string(out), "Can not find path")
}
func Test_GetWalletInformation(t *testing.T) {
	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/")

	cmd := exec.Command("./zbox", "getwallet")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}
	DeleteFileFromBlobber("/1.txt", allocationID)

	require.Contains(t, string(out), publicKey)
}
func Test_AllocCancelFail(t *testing.T) {

	cmd := exec.Command("alloc-cancel", "--allocation",allocationID)
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		return
	}

	require.Contains(t, string(out), "Error creating allocation")
}
func Test_DownloadFileFromBlobbers(t *testing.T) {
	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/2.txt")
	token := ShareFile("/2.txt", allocationID)

	cmd := exec.Command("./zbox", "download", "--allocation", allocationID, "--authticket", token, "--localpath", dirPath+"/integrationtest")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		// go LockTokensIntoReadPool(allocationID, "1", "40m")
	} else {
		fmt.Printf("%s", out)
	}
	DeleteFileFromBlobber("/2.txt", allocationID)
	DeleteLocalFile("integrationtest/2.txt")
	require.Contains(t, string(out), "Status completed callback. Type = application/octet-stream. Name = 2.txt")
}
func Test_DownloadFileFromBlobbersFailTimeOut(t *testing.T) {
	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/2.txt")
	token := ShareFile("/2.txt", allocationID)

	cmd := exec.Command("./zbox", "download", "--allocation", allocationID, "--authticket", token, "--localpath", dirPath+"/integrationtest")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		LockTokensIntoReadPool(allocationID, "1", "40m")
	} else {
		fmt.Printf("%s", out)
	}
	DeleteFileFromBlobber("/2.txt", allocationID)
	DeleteLocalFile("integrationtest/2.txt")
	require.Contains(t, string(out), "Status completed callback. Type = application/octet-stream. Name = 2.txt")
}
func Test_DownloadFileFromBlobbersFailCannotFindPath(t *testing.T) {
	token := ShareFile("/2.txt", allocationID)

	cmd := exec.Command("./zbox", "download", "--allocation", allocationID, "--authticket", token, "--localpath", dirPath+"/integrationtest")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		// go LockTokensIntoReadPool(allocationID, "1", "40m")
	} else {
		fmt.Printf("%s", out)
	}
	DeleteFileFromBlobber("/2.txt", allocationID)
	DeleteLocalFile("integrationtest/2.txt")
	require.Contains(t, string(out), "Status completed callback. Type = application/octet-stream. Name = 2.txt")
}
func Test_DownloadFileFromBlobbersFailEmptyToken(t *testing.T) {
	UploadFile(dirPath+"/integrationtest/1.txt", allocationID, "/2.txt")
	token := ""

	cmd := exec.Command("./zbox", "download", "--allocation", allocationID, "--authticket", token, "--localpath", dirPath+"/integrationtest")
	cmd.Dir = dirPath // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		// go LockTokensIntoReadPool(allocationID, "1", "40m")
	} else {
		fmt.Printf("%s", out)
	}
	DeleteFileFromBlobber("/2.txt", allocationID)
	DeleteLocalFile("integrationtest/2.txt")
	require.Contains(t, string(out), "Status completed callback. Type = application/octet-stream. Name = 2.txt")
}

func TestRegister(t *testing.T) {
	cmd := exec.Command("./zbox", "register")         // or whatever the program is
	cmd.Dir = "/Users/dev/Desktop/0chain/zboxcli" // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}
	require.Equal(t, "Wallet registered\n", string(out))
}

