package environment

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVar interface {
	FFMPEGPath() string
	BentoPath() string
	DashPath() string
	SrcVidPath() string
	OutputDir() string
	loadVariablesFromDotEnv() error
}

type EnvVarImpl struct {
	FFMPEGFullPath  string
	BentoFullPath   string
	DashFullPath    string
	SrcVidFullPath  string
	OutputDirectory string
}

var (
	Env EnvVar
)

func InitEnvironmentVariables() error {
	Env = createEnvVar()

	err := Env.loadVariablesFromDotEnv()
	if err != nil {
		return err
	}

	return nil
}

func (e *EnvVarImpl) FFMPEGPath() string {
	return e.FFMPEGFullPath
}

func (e *EnvVarImpl) BentoPath() string {
	return e.BentoFullPath
}

func (e *EnvVarImpl) DashPath() string {
	return e.DashFullPath
}

func (e *EnvVarImpl) SrcVidPath() string {
	return e.SrcVidFullPath
}

func (e *EnvVarImpl) OutputDir() string {
	return e.OutputDirectory
}

func (e *EnvVarImpl) loadVariablesFromDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	e.FFMPEGFullPath = os.Getenv("FFMPEG_PATH")
	e.BentoFullPath = os.Getenv("BENTO_PATH")
	e.DashFullPath = os.Getenv("DASH_PATH")
	e.SrcVidFullPath = os.Getenv("SRC_VID_PATH")
	e.OutputDirectory = os.Getenv("OUTPUT_DIR")

	return nil
}

func createEnvVar() EnvVar {
	var env EnvVar = &EnvVarImpl{}
	return env
}
