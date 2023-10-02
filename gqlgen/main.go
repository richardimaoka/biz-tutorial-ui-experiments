package main

import (
	"fmt"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/commits"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "process" {
		if len(os.Args) != 3 {
			panic("process sub command needs extra argument")
		}

		tutorial := os.Args[2]
		err := process.Process2(tutorial, fmt.Sprintf("https://github.com/richardimaoka/article-%s.git", tutorial))
		if err != nil {
			panic(err)
		}
	} else if len(os.Args) > 1 && os.Args[1] == "rough" {
		if len(os.Args) != 3 {
			panic("rough sub command needs extra argument")
		}

		tutorial := os.Args[2]
		err := rough.Process("data/"+tutorial, fmt.Sprintf("https://github.com/richardimaoka/article-%s.git", tutorial))
		if err != nil {
			panic(err)
		}
	} else if len(os.Args) > 1 && os.Args[1] == "csv" {
		err := rough.ConvertBoolean("data/gqlgen-getting-started/steps2.json", "data/gqlgen-getting-started/steps3.json")
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
