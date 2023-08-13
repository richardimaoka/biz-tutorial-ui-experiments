package main

import (
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "process" {
		err := process.Process2("sign-in-with-google", "https://github.com/richardimaoka/sign-in-with-google-experiment.git")
		if err != nil {
			panic(err)
		}
	} else if len(os.Args) > 1 && os.Args[1] == "processing" {
		process.Process("data/sign-in-with-google-1st-half")
	} else {
		Server()
	}
}
