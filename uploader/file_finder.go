package uploader

import (
	"os"
	"path/filepath"
	"strings"
)

type fileInformation struct {
	path    string
	extType fileExtensionType
}

type fileExtensionType int

const (
	manifest fileExtensionType = iota
	audio
	video
)

func sendAllFilePathInsideDirectory(directoryPath string, fileInfoChan chan<- fileInformation) error {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			childDirPath := filepath.Join(directoryPath, file.Name())
			err = sendAllFilePathInsideDirectory(childDirPath, fileInfoChan)
			if err != nil {
				return err
			}
			continue
		}

		filePath := filepath.Join(directoryPath, file.Name())
		fileInfo := fileInformation{
			path: filePath,
		}

		switch {
		case strings.Contains(filePath, audioStr):
			fileInfo.extType = audio
		case strings.Contains(filePath, videoStr):
			fileInfo.extType = video
		default:
			fileInfo.extType = manifest
		}

		fileInfoChan <- fileInfo
	}

	return nil
}
