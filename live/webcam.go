// Package live open and capture camera,  save video to disk, and return io.ReadCloser
package live

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/0chain/zboxcli/util"
// 	"github.com/hybridgroup/mjpeg"
// 	"gocv.io/x/gocv"
// )

// var (
// 	// ErrWebcamClosed camera is closed
// 	ErrWebcamClosed = errors.New("[webcam]camera is closed")
// )

// // Webcam Live streaming from webcam
// type Webcam struct {
// 	webcam *gocv.VideoCapture
// 	player *mjpeg.Stream
// 	err    error

// 	videoFile   string
// 	videoWriter *gocv.VideoWriter
// 	videoReader *os.File
// }

// // OpenWebcam open and capture webcam
// func OpenWebcam(options ...WebcamOption) (*Webcam, error) {

// 	deviceID := 0
// 	webcam, err := gocv.OpenVideoCapture(deviceID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dir := util.GetConfigDir() + string(os.PathSeparator) + "webcam"

// 	os.MkdirAll(dir, 0744)

// 	s := &Webcam{
// 		webcam:    webcam,
// 		player:    mjpeg.NewStream(),
// 		videoFile: dir + string(os.PathSeparator) + time.Now().Format("2006-01-02_15-04-05") + ".avi",
// 	}

// 	for _, option := range options {
// 		option(s)
// 	}

// 	frame := gocv.NewMat()
// 	defer frame.Close()
// 	for {
// 		if ok := s.webcam.Read(&frame); !ok {
// 			return nil, ErrWebcamClosed
// 		}
// 		if frame.Empty() {
// 			continue
// 		}

// 		//MJPG
// 		// In Fedora: DIVX, XVID, MJPG, X264, WMV1, WMV2. (XVID is more preferable. MJPG results in high size video. X264 gives very small size video)
// 		// In Windows: DIVX (More to be tested and added)
// 		// In OSX: MJPG (.mp4), DIVX (.avi), X264 (.mkv).
// 		s.videoWriter, err = gocv.VideoWriterFile(s.videoFile, "MJPG", 25, frame.Cols(), frame.Rows(), true)
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = s.videoWriter.Write(frame)

// 		if err != nil {
// 			return nil, err
// 		}

// 		s.videoReader, err = os.Open(s.videoFile)
// 		if err != nil {
// 			return nil, err
// 		}

// 		break

// 	}

// 	go s.goLive()
// 	go s.startWebserver()

// 	return s, nil

// }

// // startWebserver start web server for playing video
// func (s *Webcam) startWebserver() {
// 	if s.player != nil {
// 		host := ":8445"
// 		log.Println("[live]Webcam is running on http://127.0.0.1" + host)
// 		http.Handle("/", s.player)
// 		log.Fatal(http.ListenAndServe(host, nil))
// 	}
// }

// // goLive capture video, flush content into stream
// func (s *Webcam) goLive() {
// 	img := gocv.NewMat()
// 	defer img.Close()

// 	var err error

// 	for {
// 		if ok := s.webcam.Read(&img); !ok {
// 			log.Println("Webcam closed")
// 			s.err = ErrWebcamClosed
// 			break
// 		}
// 		if img.Empty() {
// 			continue
// 		}

// 		go s.play(img)

// 		err = s.videoWriter.Write(img)
// 		if err != nil {
// 			log.Println("[webcam]", err.Error())
// 			s.err = err
// 			break
// 		}

// 	}
// }

// func (s *Webcam) play(frame gocv.Mat) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("[webcame] ", r)
// 		}
// 	}()
// 	buf, _ := gocv.IMEncode(".jpg", frame)
// 	s.player.UpdateJPEG(buf)

// }

// // Read implements io.Raader
// func (s *Webcam) Read(p []byte) (int, error) {

// 	n := len(p)

// 	readLen := 0

// 	for i := 0; i < n; i++ {
// 		//loop read bytes till ready
// 		for {

// 			if s.err != nil {
// 				return readLen, s.err
// 			}

// 			if s.videoReader != nil {
// 				buf := make([]byte, 1)
// 				m, _ := s.videoReader.Read(buf)

// 				if m == 1 {
// 					readLen++
// 					p[i] = buf[0]
// 					break
// 				}
// 			}
// 			time.Sleep(1 * time.Second)
// 		}

// 	}

// 	return readLen, nil
// }

// // Close implements io.Closer
// func (s *Webcam) Close() error {

// 	if s != nil {
// 		if s.videoWriter != nil {
// 			s.videoWriter.Close()
// 		}

// 		if s.videoReader != nil {
// 			s.videoReader.Close()
// 		}
// 	}

// 	return nil
// }

// // GetVideoFile get full path of video file
// func (s *Webcam) GetVideoFile() string {
// 	return s.videoFile
// }
