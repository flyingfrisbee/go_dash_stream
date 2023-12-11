package uploader

import (
	env "GithubRepository/go_dash_stream/environment"
	"fmt"
	"os"
)

type uploader struct {
	fileInfoChan chan fileInformation
	errChan      chan error
}

func (u *uploader) StartUpload() (string, error) {
	go func() {
		err := sendAllFilePathInsideDirectory(env.OutputDir, u.fileInfoChan)
		u.errChan <- err
	}()

	for {
		select {
		case fileInfo := <-u.fileInfoChan:
			FBConn.uploadFile(&fileInfo)
			os.Remove(fileInfo.path)
		case err := <-u.errChan:
			var manifestURL string
			if err == nil {
				manifestURL = fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%smanifest.mpd?alt=media", env.BucketURL, env.FirebaseDir+"%2F")
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
