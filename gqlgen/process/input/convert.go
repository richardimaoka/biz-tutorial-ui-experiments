package input

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/input/column"
)

func convert(fromRow *Row) error {
	colType := strings.ToUpper(fromRow.Column)
	switch colType {
	case column.Terminal:
		toTerminalRow(fromRow)
		return nil
	default:
		return fmt.Errorf("column = '%s' is invalid", fromRow.Column)
	}
}
