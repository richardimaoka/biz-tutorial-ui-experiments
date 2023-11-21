package state

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

type initStep struct {
	InitialStep string `json:"initialStep"`
}

func process(repo *git.Repository, tutorial, stepFile, targetDir string) error {
	// Read step file
	var steps []Step
	err := jsonwrap.Read(stepFile, &steps)
	if err != nil {
		return fmt.Errorf("result.Process failed, %v", err)
	}

	//
	// Initialize the page state
	//
	page := NewPage(repo, tutorial, steps)
	// Initial step file
	initStepFile := fmt.Sprintf("%s/_initial_step.json", targetDir)
	if err := jsonwrap.WriteJsonToFile(initStep{page.CurrentStepId()}, initStepFile); err != nil {
		return err
	}
	// Initial step page
	gqlModel := page.ToGraphQL()
	targetFile := fmt.Sprintf("%s/%s.json", targetDir, page.CurrentStepId())
	if err := jsonwrap.WriteJsonToFile(gqlModel, targetFile); err != nil {
		return err
	}

	//
	// From the 2nd step
	//
	for page.HasNext() {
		if err := page.ToNextStep(); err != nil {
			return err
		}
		gqlModel := page.ToGraphQL()

		targetFile := fmt.Sprintf("%s/%s.json", targetDir, page.CurrentStepId())
		if err := jsonwrap.WriteJsonToFile(gqlModel, targetFile); err != nil {
			return err
		}
	}

	return nil
}

func Run(subArgs []string) error {
	// Read command line arguments
	stateCmd := flag.NewFlagSet("state", flag.ExitOnError)
	dirName := stateCmd.String("dir", "", "directory name where steps.json is located and state/{stepId}.json files will be written")
	repoUrl := stateCmd.String("repo", "", "GitHub Repository URL of the tutorial")

	if len(subArgs) < 1 {
		writer := stateCmd.Output()
		fmt.Fprintln(writer, "Error - insufficient options. Pass the following options:")
		stateCmd.PrintDefaults()
		os.Exit(1)
	}

	stateCmd.Parse(subArgs)

	// Prepare variables based on parsed arguments
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: *repoUrl})
	if err != nil {
		return fmt.Errorf("state.Run() failed, %s", err)
	}

	split := strings.Split(*dirName, "/")
	tutorial := split[len(split)-1]
	stepFile := fmt.Sprintf("%s/steps.json", *dirName)
	targetDir := fmt.Sprintf("%s/state", *dirName)

	// Process the steps file and write to target
	err = process(repo, tutorial, stepFile, targetDir)
	if err != nil {
		return fmt.Errorf("state.Run() failed, %s", err)
	}

	log.Printf("state.Run() successfully written files = '%s/{step-id}.json'", targetDir)
	return nil
}
