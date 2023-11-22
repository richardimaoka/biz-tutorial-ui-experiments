package input

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func process(repo *git.Repository, inputFile, targetFile string) error {
	finder, err := NewFinder(targetFile)
	if err != nil {
		return fmt.Errorf("process failed, %s", err)
	}

	var rows []Row
	err = jsonwrap.Read(inputFile, &rows)
	if err != nil {
		return fmt.Errorf("process failed, %v", err)
	}

	steps, err := toSteps(rows, finder, repo)
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
	// Read command line arguments
	inputCmd := flag.NewFlagSet("input", flag.ExitOnError)
	tutorialName := inputCmd.String("tutorial", "", "tutorial name")
	repoUrl := inputCmd.String("repo", "", "GitHub Repository URL of the tutorial")

	if len(subArgs) < 1 {
		writer := inputCmd.Output()
		fmt.Fprintln(writer, "Error - insufficient options. Pass the following options:")
		inputCmd.PrintDefaults()
		os.Exit(1)
	}

	inputCmd.Parse(subArgs)

	// Prepare variables based on parsed arguments
	dirName := fmt.Sprintf("data/%s", *tutorialName)
	inputFile := fmt.Sprintf("%s/input.json", dirName)
	targetFile := fmt.Sprintf("%s/steps.json", dirName)

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: *repoUrl})
	if err != nil {
		return fmt.Errorf("input.Process() failed, %s", err)
	}

	// Process the input file and write to target
	err = process(repo, inputFile, targetFile)
	if err != nil {
		return fmt.Errorf("state.Run() failed, %s", err)
	}

	log.Printf("input.Run() successfully written file = '%s'", targetFile)
	return nil
}
