package uploader

import (
	"GithubRepository/go_dash_stream/environment"
	"fmt"
	"time"
)

type uploader struct {
	fileInfoChan chan fileInformation
	errChan      chan error
}

var (
	Uploader uploader = initUploader()
)

func (u *uploader) StartUpload() error {
	go func() {
		err := sendAllFilePathInsideDirectory(environment.Env.OutputDir(), u.fileInfoChan)
		u.errChan <- err
	}()

	for {
		select {
		case fileInfo := <-u.fileInfoChan:
			// upload to firebase
			<-time.After(1 * time.Second)
			fmt.Printf("sent %v to firebase\n", fileInfo)
		case err := <-u.errChan:
			return err
		}
	}
}

func initUploader() uploader {
	up := uploader{
		fileInfoChan: make(chan fileInformation),
		errChan:      make(chan error),
	}

	return up
}
