package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "processing" {
		dirName := "data/gqlgensandbox"
		repoUrl := "https://github.com/richardimaoka/gqlgensandbox"

		// dirName := "data/protoc-go-experiments"
		// repoUrl := ""

		if repoUrl == "" {
			if err := EffectProcessing(dirName); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := GitEffectProcessing(dirName, repoUrl); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		Server()
	}
}
