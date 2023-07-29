package main

import (
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "process2" {
		err := process.Process2("sign-in-with-google")
		if err != nil {
			panic(err)
		}
	} else if len(os.Args) > 1 && os.Args[1] == "processing" {
		process.Process("data/sign-in-with-google")
	} else {
		Server()
	}
}
