package uploader

import (
	"GithubRepository/go_dash_stream/environment"
	"fmt"
	"os"
)

type uploader struct {
	fileInfoChan chan fileInformation
	errChan      chan error
}

func (u *uploader) StartUpload(movieTitle, episode string) (string, error) {
	go func() {
		err := sendAllFilePathInsideDirectory(environment.Env.OutputDir(), u.fileInfoChan)
		u.errChan <- err
	}()

	for {
		select {
		case fileInfo := <-u.fileInfoChan:
			FBConn.uploadFile(movieTitle, episode, &fileInfo)
			os.Remove(fileInfo.path)
		case err := <-u.errChan:
			var manifestURL string
			if err == nil {
				manifestURL = fmt.Sprintf(downloadURLFormat, bucketURL, movieTitle, separator, episode, separator, manifestFileName)
			}
			return manifestURL, err
		}
	}
}

var (
	Uploader uploader = initUploader()
)

func initUploader() uploader {
	up := uploader{
		fileInfoChan: make(chan fileInformation),
		errChan:      make(chan error),
	}

	return up
}
