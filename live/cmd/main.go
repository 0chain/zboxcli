package main

import (
	"log"

	"github.com/0chain/zboxcli/live"
)

func main() {
	webcam, err := live.OpenWebcam()

	if err != nil {
		log.Panic(err)
	}

	for {
		buf := make([]byte, 1)
		_, err = webcam.Read(buf)
		if err != nil {
			log.Panic(err)
		}
	}
}
