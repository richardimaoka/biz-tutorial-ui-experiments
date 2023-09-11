package main

import (
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/commits"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "process" {
		// err := process.Process2("sign-in-with-google", "https://github.com/richardimaoka/sign-in-with-google-experiment.git")
		// err := process.Process2("live-server", "https://github.com/richardimaoka/tutorial-html-live-server.git")
		err := process.Process2("gqlgen-getting-started", "https://github.com/richardimaoka/article-gqlgen-getting-started.git")
		if err != nil {
			panic(err)
		}
	} else if len(os.Args) > 1 && os.Args[1] == "processing" {
		process.Process("data/sign-in-with-google-1st-half")
	} else if len(os.Args) > 1 && os.Args[1] == "commits" {
		err := commits.Committtssss("gqlgen-getting-started")
		if err != nil {
			panic(err)
		}

	} else {
		Server()
	}
}
