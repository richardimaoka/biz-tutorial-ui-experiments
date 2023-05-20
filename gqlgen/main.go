package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "processing" {
		cases := []struct {
			dirName string
			repoUrl string
		}{
			// {
			// 	dirName: "data/apollo-client-getting-started",
			// 	repoUrl: "https://github.com/richardimaoka/apollo-client-getting-started",
			// },
			{
				dirName: "data/gqlgensandbox",
				repoUrl: "https://github.com/richardimaoka/gqlgensandbox",
			},
			{
				dirName: "data/protoc-go-experiments",
			},
		}

		for _, input := range cases {
			if input.repoUrl == "" {
				if err := EffectProcessing(input.dirName); err != nil {
					log.Fatal(err)
				}
			} else {
				if err := GitEffectProcessing(input.dirName, input.repoUrl); err != nil {
					log.Fatal(err)
				}
			}
		}
	} else {
		Server()
	}
}
