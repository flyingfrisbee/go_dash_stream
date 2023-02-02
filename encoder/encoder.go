package encoder

import (
	"GithubRepository/go_dash_stream/environment"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func EncodeVideo(movieTitle, episode string) error {
	srcVidPath := environment.Env.SrcVidPath()
	outputDir := environment.Env.OutputDir()
	mp4OutDst := filepath.Join(outputDir, "res.mp4")

	err := runEncodeToMP4(srcVidPath, mp4OutDst)
	if err != nil {
		return err
	}

	err = runCreateDashWithFFMPEG(mp4OutDst, outputDir)
	if err != nil {
		return err
	}

	err = removeUnusedFiles(mp4OutDst)
	if err != nil {
		return err
	}

	err = modifyText(movieTitle, episode)
	if err != nil {
		return err
	}

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

func modifyText(movieTitle, episode string) error {
	outputTextSlice := []string{}

	bytes, err := os.ReadFile(manifestFilePath)
	if err != nil {
		return err
	}

	sliceText := strings.Split(string(bytes), "\n")
	for _, text := range sliceText {
		shouldModifyInitDirectory := strings.Contains(text, keyword1)
		if shouldModifyInitDirectory {
			formattedText := fmt.Sprintf(replacementText1, movieTitle, separator, episode, separator)
			text = strings.Replace(text, keyword1, formattedText, 1)
		}

		shouldModifyChunkDirectory := strings.Contains(text, keyword3)
		if shouldModifyChunkDirectory {
			formattedText := fmt.Sprintf(replacementText3, movieTitle, separator, episode, separator)
			text = strings.Replace(text, keyword3, formattedText, 1)
		}

		shouldAddSuffix := strings.Contains(text, keyword2)
		if shouldAddSuffix {
			text = strings.Replace(text, keyword2, replacementText2, -1)
		}

		outputTextSlice = append(outputTextSlice, text)
	}

	modifiedText := strings.Join(outputTextSlice, "\n")
	err = os.WriteFile(manifestFilePath, []byte(modifiedText), 0644)
	if err != nil {
		return err
	}

	return nil
}
