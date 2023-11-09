package input

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func Process(repo *git.Repository, inputFile, targetFile string) error {
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
