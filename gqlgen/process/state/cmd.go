package state

import (
	"flag"
	"fmt"
	"os"

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
	// Clean up previous files
	//
	if _, err := os.Stat(targetDir); !os.IsNotExist(err) {
		if err := os.RemoveAll(targetDir); err != nil {
			return err
		}
	}
	if err := os.Mkdir(targetDir, os.ModePerm); err != nil {
		return err
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
	if err := page.ProcessCurrentStep(); err != nil {
		return err
	}
	gqlModel := page.ToGraphQL()
	targetFile := fmt.Sprintf("%s/%s.json", targetDir, page.CurrentStepId())
	if err := jsonwrap.WriteJsonToFile(gqlModel, targetFile); err != nil {
		return err
	}

	//
	// From the 2nd step
	//
	for page.HasNext() {
		if err := page.IncrementStep(); err != nil {
			return err
		}
		stepId := page.CurrentStepId()

		// PUT a debug breakpoint here for a particular step id
		// stepId == xxxxxxx
		if err := page.ProcessCurrentStep(); err != nil {
			return err
		}
		gqlModel := page.ToGraphQL()

		targetFile := fmt.Sprintf("%s/%s.json", targetDir, stepId)
		if err := jsonwrap.WriteJsonToFile(gqlModel, targetFile); err != nil {
			return err
		}
	}

	return nil
}

type MetaData struct {
	Repository string `json:"repository"`
}

func metadataFileName(dirName string) string {
	return fmt.Sprintf("%s/metadata.json", dirName)
}

func writeRepoUrl(dirName, repoUrl string) error {
	meta := MetaData{
		Repository: repoUrl,
	}

	metadataFile := metadataFileName(dirName)
	err := jsonwrap.WriteJsonToFile(&meta, metadataFile)
	if err != nil {
		return err
	}
	return nil
}

func getRepoUrlFromFile(dirName string) (string, error) {
	var meta MetaData
	metadataFile := metadataFileName(dirName)
	err := jsonwrap.Read(metadataFile, &meta)
	if err != nil {
		return "", err
	}
	return meta.Repository, nil
}

func Run(subArgs []string) error {
	errorPrefix := "state.Run() failed, "

	// Read command line arguments
	stateCmd := flag.NewFlagSet("state", flag.ExitOnError)
	tutorialName := stateCmd.String("tutorial", "", "tutorial name")
	repoUrlArg := stateCmd.String("repo", "", "(optional, if 'metadata.json' is already populated) GitHub Repository URL of the tutorial")

	if len(subArgs) < 2 /* 2 = state, tutorial */ {
		writer := stateCmd.Output()
		fmt.Fprintln(writer, "Error - insufficient options. Pass the following options:")
		stateCmd.PrintDefaults()
		os.Exit(1)
	}

	stateCmd.Parse(subArgs)

	// Prepare variables based on parsed arguments
	dirName := fmt.Sprintf("data/%s", *tutorialName)
	stepFile := fmt.Sprintf("%s/steps.json", dirName)
	targetDir := fmt.Sprintf("%s/state", dirName)

	// Validate arguments
	if *tutorialName == "" {
		return fmt.Errorf(errorPrefix + "tutorial argument is empty!")
	}

	var repoUrl string
	if *repoUrlArg == "" {
		var err error
		repoUrl, err = getRepoUrlFromFile(dirName)
		if err != nil {
			metadataFile := metadataFileName(dirName)
			return fmt.Errorf(
				errorPrefix+"cannot get repoUrl from either of 'repo' command-line argument and '%s', %w",
				metadataFile,
				err,
			)

		}
	} else {
		repoUrl = *repoUrlArg
	}
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return fmt.Errorf(errorPrefix+"%s", err)
	}

	// Write RepoUrl to tutorial's metadata file
	if repoUrlArg == nil {
		err := writeRepoUrl(dirName, repoUrl)
		if err != nil {
			return fmt.Errorf(errorPrefix+"%s", err)
		}
	}

	// Process the steps file and write to target
	err = process(repo, *tutorialName, stepFile, targetDir)
	if err != nil {
		return fmt.Errorf(errorPrefix+"%s", err)
	}

	return nil
}
