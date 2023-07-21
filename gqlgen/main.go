package main

import (
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "process2" {
		process.Process2("data/sign-in-with-google")
	} else if len(os.Args) > 1 && os.Args[1] == "processing" {
		process.Process("data/sign-in-with-google")
	} else {
		Server()
	}
}
