package input

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
)

func toSteps(
	rows []Row,
	finder *StepIdFinder,
	repo *git.Repository,
) ([]result.Step, error) {
	currentColumns := &ColumnInfo{}
	currentCommit := ""

	var allSteps []result.Step

	for _, fromRow := range rows {
		var steps []result.Step

		column, err := toColumnType(fromRow.Column)
		if err != nil {
			return nil, fmt.Errorf("column = '%s' is invalid", fromRow.Column)
		}

		switch column {
		case TerminalColumn:
			steps, currentColumns, err = toTerminalSteps(&fromRow, finder, currentColumns)
		case SourceColumn:
			steps, currentColumns, err = toSourceSteps(&fromRow, finder, currentColumns, repo, currentCommit)
		case BrowserColumn:
			steps, currentColumns, err = toBrowserSteps(&fromRow, finder, currentColumns)
		default:
			err = fmt.Errorf("column = '%s' is not implemented", fromRow.Column)
		}

		if err != nil {
			return nil, fmt.Errorf("mainLoop failed for step = %s, %s ", fromRow.StepId, err)
		}

		allSteps = append(allSteps, steps...)
	}

	return allSteps, nil
}
