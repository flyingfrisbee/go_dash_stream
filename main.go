package main

import (
	"GithubRepository/go_dash_stream/encoder"
	"GithubRepository/go_dash_stream/environment"
	"log"
)

func main() {
	err := environment.InitEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = encoder.EncodeVideo()
	if err != nil {
		log.Fatal(err)
		return
	}
}
