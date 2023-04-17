package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zboxcore/zboxutil"

	// "github.com/0chain/zboxcli/util"

	// "github.com/0chain/zboxcli/util"
	"github.com/spf13/cobra"
)

func getUp(localPath string, remotePath string) (sdk.FileMeta, io.Reader) {
	
		fileReader, err := os.Open(localPath)
		if err != nil {
			fmt.Println(err)
		}
		// defer fileReader.Close()

		fileInfo, err := fileReader.Stat()
		if err != nil {
			// return err
			fmt.Println(err)
		}
		mimeType, err := zboxutil.GetFileContentType(fileReader)
		if err != nil {
			// return err
			fmt.Println(err)
		}

		remotePath, fileName, err := fullPathAndFileNameForUpload(localPath, remotePath)
		if err != nil {
			// return err
			fmt.Println(err)
		}

		fileMeta := sdk.FileMeta{
			Path:       localPath,
			ActualSize: fileInfo.Size(),
			MimeType:   mimeType,
			RemoteName: fileName,
			RemotePath: remotePath,
		}
		return fileMeta, fileReader
}

// multiCmd represents move command
var multiCmd = &cobra.Command{
	Use:   "multi",
	Short: "Temp change to check multi operation workflow",
	Long:  `Temp change to check multi operation workflow`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                      // fflags is a *flag.FlagSet
		if fflags.Changed("allocation") == false { // check if the flag "path" is set
			fmt.Println("Error: allocation flag is missing") // If not, we'll let the user know
			return                                           // and return
		}
		// if fflags.Changed("remotepath") == false {
		// 	fmt.Println("Error: remotepath flag is missing")
		// 	return
		// }

		// if fflags.Changed("destpath") == false {
		// 	fmt.Println("Error: destpath flag is missing")
		// 	return
		// }
		allocationID := cmd.Flag("allocation").Value.String()
		allocationObj, err := sdk.GetAllocation(allocationID)
		if err != nil {
			fmt.Println("Error fetching the allocation", err)
			return
		}
		idx := 0
		_ = allocationObj;
		oReqs := make([]sdk.OperationRequest, 2)

		idx = 1
		oReqs[idx].OperationType = "delete"
		oReqs[idx].RemotePath = "/f2_again.txt"
		// oReqs[idx].DestPath = "/moved_folder/"
		// idx = 1
		// oReqs[idx].OperationType = "delete"
		// oReqs[idx].RemotePath = "/f2.txt/"
		// oReqs[idx].DestPath = "/moved_folder2/"

		idx = 0
		oReqs[idx].OperationType = "rename"
		oReqs[idx].RemotePath = "/f2_rn.txt"
		oReqs[idx].DestName =  "f2_again.txt"

		// idx = 1
		// oReqs[idx].OperationType = "rename"
		// oReqs[idx].RemotePath = "/f3.txt"
		// oReqs[idx].DestName =  "f3_rn.txt"
		
		
		// idx = 2
		// oReqs[idx].OperationType = "delete"
		// oReqs[idx].RemotePath = "/t8.txt"
		// oReqs[idx].DestPath = "/renamed_folder/"

		// idx = 2
		// oReqs[idx].OperationType = "copy"
		// oReqs[idx].RemotePath = "/f3.txt/"
		// oReqs[idx].DestPath = "/copied_folder4/"

		

		// oReqs[4].OperationType = "copy"
		// oReqs[4].RemotePath = "/t5.txt/"
		// oReqs[4].DestPath = "/renamed_folder6/"

		
		
		// oReqs[2].OperationType = "move"
		// oReqs[2].RemotePath = "/t2_renamed.txt"
		// oReqs[2].DestPath = "/renamed_folder/"
		
		// oReqs[3].OperationType = "move"
		// oReqs[3].RemotePath = "/t4.txt"
		// oReqs[3].DestPath = "/newfolder/"
		
		// localPath := "files_upload_4/f1.txt"
		// remotePath := "/"
		// fileMeta, fileReader := getUp(localPath, remotePath);
		
		// _ = fileMeta;
		// _ = util.GetHomeDir()

		// idx = 0

		// oReqs[idx].OperationType = "insert"
		// oReqs[idx].FileMeta = fileMeta
		// oReqs[idx].FileReader = fileReader
		// oReqs[idx].Opts = make([]sdk.ChunkedUploadOption, 0)
		// oReqs[idx].Workdir = util.GetHomeDir()

		
		// localPath = "files_upload_4/f2.txt"
		// remotePath = "/"
		// fileMeta, fileReader2 := getUp(localPath, remotePath);
		// idx = 1
		// oReqs[idx].OperationType = "insert"
		// oReqs[idx].FileMeta = fileMeta
		// oReqs[idx].FileReader = fileReader2
		// oReqs[idx].Opts = make([]sdk.ChunkedUploadOption, 0)
		// oReqs[idx].Workdir = util.GetHomeDir()

		// defer fileReader.Close()



		err = allocationObj.DoMultiOperation(oReqs)
		if err != nil {
			fmt.Println("Error : ", err);
		}
		// remotePath := cmd.Flag("remotepath").Value.String()
		// destPath := cmd.Flag("destpath").Value.String()

		// err = allocationObj.MoveObject(remotePath, destPath)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }

		// fmt.Println(remotePath + " moved")

	},
}

func init() {
	rootCmd.AddCommand(multiCmd)
	multiCmd.PersistentFlags().String("allocation", "", "Allocation ID")
	// multiCmd.PersistentFlags().String("remotepath", "", "Remote path of object to move")
	// multiCmd.PersistentFlags().String("destpath", "", "Destination path for the object. Existing directory the object should be copied to")

	multiCmd.MarkFlagRequired("allocation")
	// multiCmd.MarkFlagRequired("remotepath")
	// multiCmd.MarkFlagRequired("destpath")
}
