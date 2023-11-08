package input

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
)

func process(inputFile, targetFile string) error {
	finder, err := NewFinder(targetFile)
	if err != nil {
		return fmt.Errorf("process failed, %s", err)
	}

	currentColumns := ColumnInfo{}

	var rows []Row
	err = jsonwrap.Read(inputFile, &rows)
	if err != nil {
		return fmt.Errorf("process failed, %v", err)
	}

	var resultSteps []result.Step
	for _, r := range rows {
		steps, updatedColumns, err := convert(&r, finder, currentColumns)
		if err != nil {
			return fmt.Errorf("process failed, %s", err)
		}
		resultSteps = append(resultSteps, steps...)
		currentColumns = updatedColumns
	}

	return nil
}

func convert(
	fromRow *Row,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
) ([]result.Step, ColumnInfo, error) {

	column := strings.ToUpper(fromRow.Column)
	switch column {
	case TerminalType:
		steps, currentColumns, err := toTerminalSteps(fromRow, finder, prevColumns)
		if err != nil {
			return nil, prevColumns, fmt.Errorf("convert failed, %s", err)
		}
		return steps, currentColumns, nil
	default:
		return nil, prevColumns, fmt.Errorf("convert failed, column = '%s' is invalid", fromRow.Column)
	}
}
