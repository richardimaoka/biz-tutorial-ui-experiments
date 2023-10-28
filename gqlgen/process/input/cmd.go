package input

import (
	"flag"
	"fmt"
	"os"
)

func Run(subArgs []string) {
	inputCmd := flag.NewFlagSet("input", flag.ExitOnError)
	inputFileName := inputCmd.String("name", "", "input .json file name")

	if len(subArgs) < 1 {
		writer := inputCmd.Output()
		fmt.Fprintln(writer, "Error - insufficient options. Pass the following options:")
		inputCmd.PrintDefaults()
		os.Exit(1)
	}

	inputCmd.Parse(subArgs)

	fmt.Printf("input file name '%s'\n", *inputFileName)
}
