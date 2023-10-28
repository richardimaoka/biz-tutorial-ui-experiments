package input

import (
	"flag"
	"fmt"
)

func Run(subArgs []string) {
	inputCmd := flag.NewFlagSet("input", flag.ExitOnError)
	inputFileName := inputCmd.String("name", "", "input .json file name")

	fmt.Println(inputFileName)
}
