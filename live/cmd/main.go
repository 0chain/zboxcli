package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jtguibas/cinema"
)

func main() {
	// webcam, err := live.OpenWebcam()

	// if err != nil {
	// 	log.Panic(err)
	// }

	// for {
	// 	buf := make([]byte, 1)
	// 	_, err = webcam.Read(buf)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// }

	video, err := cinema.Load("/Users/geax/.zcn/youtubedl/2021-07-06_21-53-43_pJ0VgJloR9E.mp4")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(video.End())

	video.SetStart(1 * time.Second)
	video.SetEnd(30 * time.Second)
	video.Render("/Users/geax/.zcn/youtubedl/split.mp4")

	fileReader, _ := os.Open("/Users/geax/.zcn/youtubedl/2021-07-06_21-53-43_pJ0VgJloR9E.mp4")

	fileWriter, _ := os.Create("/Users/geax/.zcn/youtubedl/test.mp4")

	defer fileReader.Close()
	defer fileWriter.Close()

	var offset int64
	for {

		// if offset >= 10*1024*1024 {
		// 	fmt.Println("done")
		// 	break
		// }

		//fi, _ := fileReader.Stat()

		buf := make([]byte, 10*1024*1024)

		//wantRead := int64(len(buf))

		//	if offset+wantRead < fi.Size() {
		//fileReader.Seek(offset, 0)
		readLen, _ := fileReader.Read(buf)

		offset += int64(readLen)

		fmt.Println(offset, readLen)

		if readLen == 0 {
			fmt.Println("io.EOF")
			break
		}

		fileWriter.Write(buf[0:readLen])
		//	}

		//time.Sleep(1 * time.Second)
	}

}
