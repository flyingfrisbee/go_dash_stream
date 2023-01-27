package environment

import (
	"os"

	"github.com/joho/godotenv"
)

type envVar interface {
	FFMPEGPath() string
	BentoPath() string
	DashPath() string
	SrcVidPath() string
	OutputDir() string
	loadVariablesFromDotEnv() error
}

type envVarImpl struct {
	ffmpegFullPath  string
	bentoFullPath   string
	dashFullPath    string
	srcVidFullPath  string
	outputDirectory string
}

var (
	Env envVar
)

func InitEnvironmentVariables() error {
	Env = createEnvVar()

	err := Env.loadVariablesFromDotEnv()
	if err != nil {
		return err
	}

	return nil
}

func (e *envVarImpl) FFMPEGPath() string {
	return e.ffmpegFullPath
}

func (e *envVarImpl) BentoPath() string {
	return e.bentoFullPath
}

func (e *envVarImpl) DashPath() string {
	return e.dashFullPath
}

func (e *envVarImpl) SrcVidPath() string {
	return e.srcVidFullPath
}

func (e *envVarImpl) OutputDir() string {
	return e.outputDirectory
}

func (e *envVarImpl) loadVariablesFromDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	e.ffmpegFullPath = os.Getenv("FFMPEG_PATH")
	e.bentoFullPath = os.Getenv("BENTO_PATH")
	e.dashFullPath = os.Getenv("DASH_PATH")
	e.srcVidFullPath = os.Getenv("SRC_VID_PATH")
	e.outputDirectory = os.Getenv("OUTPUT_DIR")

	return nil
}

func createEnvVar() envVar {
	var env envVar = &envVarImpl{}
	return env
}
