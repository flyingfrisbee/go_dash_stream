package encoder

import (
	"GithubRepository/go_dash_stream/environment"
	"fmt"
	"os/exec"
	"strings"
)

func runCreateFrMP4(srcVidPath, outputPath string) error {
	bentoPath := environment.Env.BentoPath()
	args := generateCreateFrMP4Args(bentoPath, srcVidPath, outputPath)

	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func generateCreateFrMP4Args(bentoPath, srcVidPath, outputPath string) []string {
	text := fmt.Sprintf(createFrMP4, bentoPath, srcVidPath, outputPath)
	return strings.Fields(text)
}

func runCreateDash(srcVidPath, outputDir string) error {
	dashPath := environment.Env.DashPath()
	args := generateCreateDashArgs(dashPath, srcVidPath, outputDir)

	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func generateCreateDashArgs(dashPath, srcVidPath, outputDir string) []string {
	text := fmt.Sprintf(createDash, dashPath, outputDir, srcVidPath)
	return strings.Fields(text)
}
