package input

import (
	"fmt"
	"strings"
)

func convert(fromRow *Row) error {
	column := strings.ToUpper(fromRow.Column)
	switch column {
	case TerminalType:
		toTerminalRow(fromRow)
		return nil
	default:
		return fmt.Errorf("column = '%s' is invalid", fromRow.Column)
	}
}
