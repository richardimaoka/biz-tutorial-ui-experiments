package main

import (
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/input"
)

func main() {
	subCmdError := "Sub-command not provided. Specify 'process', 'input', 'rough', 'csv', 'commits', or 'server'."
	if len(os.Args) < 2 {
		log.Fatalf(subCmdError)
	}

	subCmd := os.Args[1]

	switch subCmd {
	case "input":
		err := input.Run(os.Args[2:])
		if err != nil {
			log.Print(err)
		}
	case "server":
		Server()
	default:
		log.Fatalf(subCmdError)
	}
}
