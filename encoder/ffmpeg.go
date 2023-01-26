package encoder

import (
	"GithubRepository/go_dash_stream/environment"
	"fmt"
	"os/exec"
	"strings"
)

func runEncodeToMP4(srcVidPath, outputPath string) error {
	ffmpegPath := environment.Env.FFMPEGPath()
	args := generateEncodeToMP4Args(ffmpegPath, srcVidPath, outputPath)

	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func generateEncodeToMP4Args(ffmpegPath, srcVidPath, outputPath string) []string {
	text := fmt.Sprintf(encodeToMP4, ffmpegPath, srcVidPath, srcVidPath, outputPath)
	return strings.Fields(text)
}
