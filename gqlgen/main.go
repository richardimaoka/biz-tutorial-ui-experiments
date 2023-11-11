package main

import (
	"fmt"
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/commits"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/input"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func main() {
	subCmdError := "Sub-command not provided. Specify 'process', 'input', 'rough', 'csv', 'commits', or 'server'."
	if len(os.Args) < 2 {
		log.Fatalf(subCmdError)
	}

	subCmd := os.Args[1]

	switch subCmd {
	case "input":
		input.Run(os.Args[2:])
	case "rough":
		if len(os.Args) != 3 {
			panic("rough sub command needs extra argument")
		}

		tutorial := os.Args[2]
		err := rough.Process("data/"+tutorial, fmt.Sprintf("https://github.com/richardimaoka/article-%s.git", tutorial))
		if err != nil {
			panic(err)
		}
	case "csv":
		err := rough.ConvertBoolean("data/gqlgen-getting-started/steps2.json", "data/gqlgen-getting-started/steps3.json")
		if err != nil {
			panic(err)
		}
	case "commits":
		err := commits.Committtssss("gqlgen-getting-started")
		if err != nil {
			panic(err)
		}
	case "server":
		Server()
	default:
		log.Fatalf(subCmdError)
	}
}
