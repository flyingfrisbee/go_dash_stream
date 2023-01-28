package main

import (
	enc "GithubRepository/go_dash_stream/encoder"
	"GithubRepository/go_dash_stream/environment"
	up "GithubRepository/go_dash_stream/uploader"
	"fmt"
	"log"
)

func main() {
	err := environment.InitEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = up.InitFirebaseConn()
	if err != nil {
		log.Fatal(err)
		return
	}

	title := "tsurune_s2"
	episode := "4"

	err = enc.EncodeVideo(title, episode)
	if err != nil {
		log.Fatal(err)
		return
	}

	manifestURL, err := up.Uploader.StartUpload(title, episode)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(manifestURL)
}
