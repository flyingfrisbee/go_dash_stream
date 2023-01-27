package main

import (
	enc "GithubRepository/go_dash_stream/encoder"
	"GithubRepository/go_dash_stream/environment"
	up "GithubRepository/go_dash_stream/uploader"
	"log"
)

func main() {
	err := environment.InitEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = enc.EncodeVideo("tsurunes2", "4")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = up.Uploader.StartUpload()
	if err != nil {
		log.Fatal(err)
		return
	}
}
