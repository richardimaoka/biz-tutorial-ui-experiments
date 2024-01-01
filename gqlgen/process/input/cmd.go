package input

import (
	"flag"
	"fmt"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func process(inputFile, targetFile string) error {
	finder, err := NewFinder(targetFile)
	if err != nil {
		return fmt.Errorf("process failed, %s", err)
	}

	var rows []Row
	err = jsonwrap.Read(inputFile, &rows)
	if err != nil {
		return fmt.Errorf("process failed, %v", err)
	}

	steps, err := toSteps(rows, finder)
	if err != nil {
		return fmt.Errorf("process failed, %v", err)
	}

	err = jsonwrap.WriteJsonToFile(steps, targetFile)
	if err != nil {
		return fmt.Errorf("process failed, %v", err)
	}

	return nil
}

func Run(subArgs []string) error {
	errorPrefix := "input.Run() failed, "

	// Read command line arguments
	inputCmd := flag.NewFlagSet("input", flag.ExitOnError)
	tutorialName := inputCmd.String("tutorial", "", "tutorial name")
	// repoUrl := inputCmd.String("repo", "", "GitHub Repository URL of the tutorial")

	if len(subArgs) < 2 /* 2 = input, tutorial */ {
		writer := inputCmd.Output()
		fmt.Fprintln(writer, "Error - insufficient options. Pass the following options:")
		inputCmd.PrintDefaults()
		os.Exit(1)
	}

	inputCmd.Parse(subArgs)

	// Validate arguments
	if *tutorialName == "" {
		return fmt.Errorf(errorPrefix + "tutorial argument is empty!")
	}

	// Prepare variables based on parsed arguments
	dirName := fmt.Sprintf("data/%s", *tutorialName)
	inputFile := fmt.Sprintf("%s/input.json", dirName)
	targetFile := fmt.Sprintf("%s/steps.json", dirName)

	// Process the input file and write to target
	if err := process(inputFile, targetFile); err != nil {
		return fmt.Errorf(errorPrefix+"%s", err)
	}

	return nil
}
