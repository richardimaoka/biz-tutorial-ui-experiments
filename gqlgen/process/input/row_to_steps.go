package input

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func toSteps(
	rows []Row,
	finder *StepIdFinder,
) ([]state.Step, error) {
	currentColumn := state.NoColumnType

	var allSteps []state.Step

	for _, fromRow := range rows {
		errorPrefix := fmt.Sprintf("toSteps() failed for row = '%s'", fromRow.RowId)

		var steps []state.Step

		mode, err := toMode(fromRow.Mode)
		if err != nil {
			return nil, fmt.Errorf("%s, mode = '%s' is invalid", errorPrefix, fromRow.Mode)
		}

		switch mode {
		case "slideshow":
			slideType, err := toSlideType(fromRow.RowType)
			if err != nil {
				return nil, fmt.Errorf("%s, slide = '%s' is invalid", errorPrefix, fromRow.RowType)
			}

			switch slideType {
			case TutorialTitleSlide:
				steps, err = toTutorialTitleSteps(&fromRow, finder, currentColumn)
			case SectionTitleSlide:
				steps, err = toSectionTitleSteps(&fromRow, finder, currentColumn)
			case MarkdownSlide:
				steps, err = toMarkdownSteps(&fromRow, finder, currentColumn)
			case ImageSlide:
				steps, err = toImageSteps(&fromRow, finder, currentColumn)
			case TocSlide:
				steps, err = toTocSteps(&fromRow, finder, currentColumn)
			default:
				err = fmt.Errorf("%s, slide = '%s' is not implemented", errorPrefix, fromRow.RowType)
			}

			if err != nil {
				return nil, fmt.Errorf("%s, %s", errorPrefix, err)
			}

		case "handson":
			column, err := toColumnType(fromRow.RowType)
			if err != nil {
				return nil, fmt.Errorf("%s, column = '%s' is invalid", errorPrefix, fromRow.RowType)
			}

			switch column {
			case TerminalColumn:
				steps, err = toTerminalSteps(&fromRow, finder, currentColumn)
				currentColumn = state.TerminalColumnType
			case SourceColumn:
				steps, err = toSourceSteps(&fromRow, finder, currentColumn)
				currentColumn = state.SourceColumnType
			case BrowserColumn:
				steps, err = toBrowserSteps(&fromRow, finder, currentColumn)
				currentColumn = state.BrowserColumnType
			default:
				err = fmt.Errorf("%s, column = '%s' is not implemented", errorPrefix, fromRow.RowType)
			}

			if err != nil {
				return nil, fmt.Errorf("%s, %s ", errorPrefix, err)
			}

		default:
			return nil, fmt.Errorf("%s, mode = '%s' is invalid", errorPrefix, fromRow.Mode)
		}

		allSteps = append(allSteps, steps...)
	}

	return allSteps, nil
}
