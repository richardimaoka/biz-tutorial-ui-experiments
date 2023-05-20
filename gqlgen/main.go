package main

import (
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "processing" {
		dirName := "data/gqlgensandbox"
		repoUrl := "https://github.com/richardimaoka/gqlgensandbox"

		// dirName := "data/protoc-go-experiments"
		// repoUrl := ""

		if repoUrl == "" {
			if err := graph.EffectProcessing(dirName); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := graph.GitEffectProcessing(dirName, repoUrl); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		Server()
	}
}
