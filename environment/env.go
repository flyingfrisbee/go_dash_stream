package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MP4BoxPath          string
	FFMPEGPath          string
	OutputDir           string
	TempDir             string
	FirebaseJSONPath    string
	BucketURL           string
	ExtractSubtitleFmt  string
	ExtractAudioFmt     string
	ExtractVideo720Fmt  string
	ExtractVideo240Fmt  string
	GenerateManifestFmt string
	SubtitleExt         string
	Mp4Ext              string
	GenericExt          string
	MetadataKey         string
	VideoPath           string
	FirebaseDir         string
)

func InitEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	MP4BoxPath = os.Getenv("MP4BOX_PATH")
	FFMPEGPath = os.Getenv("FFMPEG_PATH")
	OutputDir = os.Getenv("OUT_DIR")
	TempDir = os.Getenv("TEMP_DIR")
	FirebaseJSONPath = os.Getenv("FIREBASE_JSON_PATH")
	BucketURL = os.Getenv("BUCKET_URL")
	ExtractSubtitleFmt = os.Getenv("EXTRACT_SUBTITLE_FMT")
	ExtractAudioFmt = os.Getenv("EXTRACT_AUDIO_FMT")
	ExtractVideo720Fmt = os.Getenv("EXTRACT_VIDEO_720_FMT")
	ExtractVideo240Fmt = os.Getenv("EXTRACT_VIDEO_240_FMT")
	GenerateManifestFmt = os.Getenv("GENERATE_MANIFEST_FMT")
	SubtitleExt = os.Getenv("SUBTITLE_EXT")
	Mp4Ext = os.Getenv("MP4_EXT")
	GenericExt = os.Getenv("GENERIC_EXT")
	MetadataKey = os.Getenv("METADATA_KEY")

	VideoPath = os.Getenv("VIDEO_PATH")
	FirebaseDir = os.Getenv("FIREBASE_DIR")
}
