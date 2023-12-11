package uploader

import (
	"os"
	"path/filepath"
)

type fileInformation struct {
	name string
	path string
}

func sendAllFilePathInsideDirectory(directoryPath string, fileInfoChan chan<- fileInformation) error {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(directoryPath, file.Name())
		fileInfo := fileInformation{
			name: file.Name(),
			path: filePath,
		}

		fileInfoChan <- fileInfo
	}

	return nil
}
