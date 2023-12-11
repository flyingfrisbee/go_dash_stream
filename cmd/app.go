package cmd

import (
	"GithubRepository/go_dash_stream/encoder"
	env "GithubRepository/go_dash_stream/environment"
	up "GithubRepository/go_dash_stream/uploader"
	"fmt"
	"log"
)

func Run() {
	env.InitEnvironmentVariables()
	encoder.StartEncoder()

	err := up.InitFirebaseConn()
	if err != nil {
		log.Fatal(err)
	}

	manifestURL, err := up.Uploader.StartUpload()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(manifestURL)
}
