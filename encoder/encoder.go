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
		shouldModifyText := strings.Contains(text, keyword1)
		if shouldModifyText {
			text = strings.Replace(text, keyword1, replacementText1, -1)
			text = strings.Replace(text, keyword2, replacementText2, -1)
		}

		shouldModifyVideoDirectory := strings.Contains(text, keyword3)
		if shouldModifyVideoDirectory {
			leftIndex := strings.Index(text, keyword5) + len(keyword5)
			leftText := text[:leftIndex]

			rightIndex := strings.Index(text, keyword6)
			rightText := text[rightIndex:]

			text = fmt.Sprintf(
				textFormat,
				leftText, movieTitle,
				separator, episode, separator,
				video, rightText,
			)
		}

		shouldModifyAudioDirectory := strings.Contains(text, keyword4)
		if shouldModifyAudioDirectory {
			leftIndex := strings.Index(text, keyword5) + len(keyword5)
			leftText := text[:leftIndex]

			rightIndex := strings.Index(text, keyword6)
			rightText := text[rightIndex:]

			text = fmt.Sprintf(
				textFormat,
				leftText, movieTitle,
				separator, episode, separator,
				audio, rightText,
			)
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
