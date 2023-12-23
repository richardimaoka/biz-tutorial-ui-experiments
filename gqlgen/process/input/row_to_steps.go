package input

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func toSteps(
	rows []Row,
	finder *StepIdFinder,
	repo *git.Repository,
) ([]state.Step, error) {
	currentColumns := &ColumnInfo{}

	var allSteps []state.Step

	for _, fromRow := range rows {
		var steps []state.Step

		mode, err := toMode(fromRow.Mode)
		if err != nil {
			return nil, fmt.Errorf("mode = '%s' is invalid", fromRow.Mode)
		}

		switch mode {
		case "slideshow":
			slideType, err := toSlideType(fromRow.Column)
			if err != nil {
				return nil, fmt.Errorf("column = '%s' is invalid", fromRow.Column)
			}

			switch slideType {
			// case TutorialTitle:
			// 	steps, currentColumns, err = toTerminalSteps(&fromRow, finder, currentColumns)
			// case SectionTitle:
			// 	steps, currentColumns, err = toSourceSteps(&fromRow, finder, currentColumns, repo, currentCommit)
			// case Markdown:
			// 	steps, currentColumns, err = toBrowserSteps(&fromRow, finder, currentColumns)
			// case Image:
			// 	steps, currentColumns, err = toBrowserSteps(&fromRow, finder, currentColumns)
			default:
				err = fmt.Errorf("slide = '%s' is not implemented", fromRow.Column)
			}

			if err != nil {
				return nil, fmt.Errorf("toSteps failed for step = %s, %s ", fromRow.StepId, err)
			}

		case "handson":
			column, err := toColumnType(fromRow.Column)
			if err != nil {
				return nil, fmt.Errorf("column = '%s' is invalid", fromRow.Column)
			}

			switch column {
			case TerminalColumn:
				steps, currentColumns, err = toTerminalSteps(&fromRow, finder, currentColumns)
			case SourceColumn:
				steps, currentColumns, err = toSourceSteps(&fromRow, finder, currentColumns)
			case BrowserColumn:
				steps, currentColumns, err = toBrowserSteps(&fromRow, finder, currentColumns)
			default:
				err = fmt.Errorf("column = '%s' is not implemented", fromRow.Column)
			}

			if err != nil {
				return nil, fmt.Errorf("toSteps failed for step = %s, %s ", fromRow.StepId, err)
			}
		default:
			return nil, fmt.Errorf("toSteps failed for step = %s, mode = '%s' is invalid", fromRow.StepId, fromRow.Mode)
		}

		allSteps = append(allSteps, steps...)
	}

	return allSteps, nil
}
