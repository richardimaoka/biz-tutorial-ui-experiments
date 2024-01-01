package main

import (
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/input"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func main() {
	subCmdError := "Sub-command not provided. Specify 'process', 'input', 'rough', 'csv', 'commits', or 'server'."
	if len(os.Args) < 2 {
		log.Fatalf(subCmdError)
	}

	subCmd := os.Args[1]

	switch subCmd {
	case "input":
		err := input.Run(os.Args[2:]) //[:2] omit the first two args, main command and subcommand
		if err != nil {
			log.Print(err)
		}
		log.Printf("input.Run() successfully written file")
	case "state":
		err := state.Run(os.Args[2:]) //[:2] omit the first two args, main command and subcommand
		if err != nil {
			log.Print(err)
		}
		log.Printf("state.Run() successfully written files")
	case "process":
		err := input.Run(os.Args[2:]) //[:2] omit the first two args, main command and subcommand
		if err != nil {
			log.Print(err)
		}
		log.Printf("input.Run() successfully written file")
		err = state.Run(os.Args[2:]) //[:2] omit the first two args, main command and subcommand
		if err != nil {
			log.Print(err)
		}
		log.Printf("state.Run() successfully written files")
	case "server":
		Server()
	default:
		log.Fatalf(subCmdError)
	}
}
