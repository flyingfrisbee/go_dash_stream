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

func runCreateDashWithFFMPEG(srcVidPath, outputDir string) error {
	ffmpegPath := environment.Env.FFMPEGPath()
	outputManifestPath := fmt.Sprintf("%s/stream.mpd", outputDir)
	args := generateRunCreateDashWithFFMPEGArgs(ffmpegPath, srcVidPath, outputManifestPath)

	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("ini?")
		return err
	}

	return nil
}

func generateRunCreateDashWithFFMPEGArgs(ffmpegPath, srcVidPath, outputPath string) []string {
	text := fmt.Sprintf(createDashFFMPEG, ffmpegPath, srcVidPath, outputPath)
	return strings.Fields(text)
}
