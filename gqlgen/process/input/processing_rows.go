package input

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
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

	steps, err := mainLoop(rows, finder, repo)
	if err != nil {
		return fmt.Errorf("process failed, %v", err)
	}

	err = jsonwrap.WriteJsonToFile(steps, targetFile)
	if err != nil {
		return fmt.Errorf("process failed, %v", err)
	}

	return nil
}

func mainLoop(
	rows []Row,
	finder *StepIdFinder,
	repo *git.Repository,
) ([]result.Step, error) {
	currentColumns := ColumnInfo{}
	currentCommit := ""

	var allSteps []result.Step

	for _, fromRow := range rows {
		var steps []result.Step
		var err error

		column := strings.ToUpper(fromRow.Column)
		switch column {
		case TerminalType:
			steps, currentColumns, err = toTerminalSteps(&fromRow, finder, currentColumns)
		case SourceType:
			steps, currentColumns, err = toSourceSteps(&fromRow, finder, currentColumns, repo, currentCommit)
		default:
			err = fmt.Errorf("column = '%s' is invalid", fromRow.Column)
		}

		if err != nil {
			return nil, fmt.Errorf("mainLoop faield %s ", err)
		}

		allSteps = append(allSteps, steps...)
	}

	return allSteps, nil
}
