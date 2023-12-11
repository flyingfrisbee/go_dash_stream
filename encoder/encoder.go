package encoder

import (
	env "GithubRepository/go_dash_stream/environment"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func StartEncoder() {
	err := extractSubtitle()
	if err != nil {
		log.Fatal(err)
	}

	err = extractAudio()
	if err != nil {
		log.Fatal(err)
	}

	err = extractVideo()
	if err != nil {
		log.Fatal(err)
	}

	err = generateManifestAndChunks()
	if err != nil {
		log.Fatal(err)
	}

	err = wipeDirectory(env.TempDir)
	if err != nil {
		log.Fatal(err)
	}

	err = modifyManifestFile()
	if err != nil {
		log.Fatal(err)
	}
}

func extractSubtitle() error {
	args := strings.Fields(fmt.Sprintf(env.ExtractSubtitleFmt, env.VideoPath))
	err := runCommandFromCustomDir(env.OutputDir, env.FFMPEGPath, args...)
	if err != nil {
		return err
	}

	return nil
}

func extractAudio() error {
	args := strings.Fields(fmt.Sprintf(env.ExtractAudioFmt, env.VideoPath, env.TempDir))
	err := runCommandFromCustomDir(env.TempDir, env.FFMPEGPath, args...)
	if err != nil {
		return err
	}

	return nil
}

func extractVideo() error {
	args := strings.Fields(fmt.Sprintf(env.ExtractVideo720Fmt, env.VideoPath, env.TempDir))
	err := runCommandFromCustomDir(env.TempDir, env.FFMPEGPath, args...)
	if err != nil {
		return err
	}

	args = strings.Fields(fmt.Sprintf(env.ExtractVideo240Fmt, env.VideoPath, env.TempDir))
	err = runCommandFromCustomDir(env.TempDir, env.FFMPEGPath, args...)
	if err != nil {
		return err
	}

	return nil
}

func generateManifestAndChunks() error {
	args := strings.Fields(fmt.Sprintf(env.GenerateManifestFmt, env.OutputDir, env.TempDir, env.TempDir, env.TempDir))
	err := runCommandFromCustomDir(env.OutputDir, env.MP4BoxPath, args...)
	if err != nil {
		return err
	}

	return nil
}

func runCommandFromCustomDir(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func wipeDirectory(dir string) error {
	f, err := os.Open(dir)
	if err != nil {
		return err
	}

	files, err := f.ReadDir(-1)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := path.Join(dir, file.Name())
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func modifyManifestFile() error {
	manifestFilePath := path.Join(env.OutputDir, "manifest.mpd")
	outBytes, err := os.ReadFile(manifestFilePath)
	if err != nil {
		return err
	}
	stringSlice := strings.Split(string(outBytes), "\n")
	for i := 0; i < len(stringSlice); i++ {
		if strings.Contains(stringSlice[i], "<SegmentTemplate") {
			text := stringSlice[i]
			text = addPrefixAndSuffix("media=", text)
			text = addPrefixAndSuffix("initialization=", text)
			stringSlice[i] = text
			continue
		}
		if strings.Contains(stringSlice[i], "</Period>") {
			textsToAppend := generateSubtitleTemplateForManifestFile()
			textsToAppend = append(textsToAppend, stringSlice[i:]...)
			stringSlice = append(stringSlice[:i], textsToAppend...)
			break
		}
	}

	f, err := os.OpenFile(manifestFilePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(strings.Join(stringSlice, "\n"))
	if err != nil {
		return err
	}

	return nil
}

func addPrefixAndSuffix(keyword, text string) string {
	idxPlaceholder := strings.Index(text, keyword)
	firstIdx := idxPlaceholder + len(keyword) + 1
	lastIdx := 0

	counter := 0
	for i, v := range text[idxPlaceholder:] {
		if v == '"' {
			counter++
		}
		if counter == 2 {
			lastIdx = idxPlaceholder + i
			break
		}
	}

	var sb strings.Builder
	sb.WriteString(env.FirebaseDir)
	sb.WriteString(`%2F`)
	sb.WriteString(text[firstIdx:lastIdx])
	sb.WriteString("?alt=media")

	return strings.Replace(text, text[firstIdx:lastIdx], sb.String(), 1)
}

func generateSubtitleTemplateForManifestFile() []string {
	text := fmt.Sprintf(
		`<AdaptationSet mimeType="text/vtt" lang="en">
	<Representation id="caption" bandwidth="123">
		<BaseURL>https://firebasestorage.googleapis.com/v0/b/%s/o/%ssub.vtt?alt=media</BaseURL>
	</Representation>
</AdaptationSet>`,
		env.BucketURL,
		env.FirebaseDir+"%2F",
	)

	return strings.Split(text, "\n")
}
