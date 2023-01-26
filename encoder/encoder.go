package encoder

import (
	"GithubRepository/go_dash_stream/environment"
	"os"
	"path/filepath"
)

func EncodeVideo() error {
	srcVidPath := environment.Env.SrcVidPath()
	outputDir := environment.Env.OutputDir()
	mp4OutDst := filepath.Join(outputDir, "res.mp4")
	frMP4OutDst := filepath.Join(outputDir, "res-fr.mp4")

	err := runEncodeToMP4(srcVidPath, mp4OutDst)
	if err != nil {
		return err
	}

	err = runCreateFrMP4(mp4OutDst, frMP4OutDst)
	if err != nil {
		return err
	}

	err = runCreateDash(frMP4OutDst, outputDir)
	if err != nil {
		return err
	}

	err = removeUnusedFiles(mp4OutDst, frMP4OutDst)

	return nil
}

func removeUnusedFiles(filePaths ...string) error {
	for _, filePath := range filePaths {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
