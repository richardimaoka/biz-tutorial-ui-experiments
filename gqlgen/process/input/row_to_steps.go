package input

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
)

func toSteps(
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

		column := strings.ToLower(fromRow.Column)
		switch column {
		case TerminalType:
			steps, currentColumns, err = toTerminalSteps(&fromRow, finder, currentColumns)
		case SourceType:
			steps, currentColumns, err = toSourceSteps(&fromRow, finder, currentColumns, repo, currentCommit)
		case BrowserType:
			steps, currentColumns, err = toBrowserSteps(&fromRow, finder, currentColumns)
		default:
			err = fmt.Errorf("column = '%s' is invalid", fromRow.Column)
		}

		if err != nil {
			return nil, fmt.Errorf("mainLoop failed for step = %s, %s ", fromRow.StepId, err)
		}

		allSteps = append(allSteps, steps...)
	}

	return allSteps, nil
}
