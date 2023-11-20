package state

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func process(repo *git.Repository, tutorial, stepFile, targetDir string) error {
	var steps []Step
	err := jsonwrap.Read(stepFile, &steps)
	if err != nil {
		return fmt.Errorf("result.Process failed, %v", err)
	}

	page := NewPage(repo, tutorial)
	for _, s := range steps {
		if err := page.Update(&s); err != nil {
			return err
		}
		gqlModel := page.ToGraphQL()

		targetFile := fmt.Sprintf("%s/state/%s.json", targetDir, s.StepId)
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
		return fmt.Errorf("input.Process() failed, %s", err)
	}

	split := strings.Split(*dirName, "/")
	tutorial := split[len(split)-1]
	stepFile := fmt.Sprintf("%s/steps.json", *dirName)

	// Process the steps file and write to target
	return process(repo, tutorial, stepFile, *dirName)
}
