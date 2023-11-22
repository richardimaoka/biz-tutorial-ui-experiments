package input

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

type ColumnType string

const (
	// Lower cases since they are from manual entries
	SourceColumn   ColumnType = "source"
	TerminalColumn ColumnType = "terminal"
	BrowserColumn  ColumnType = "browser"
)

func toColumnType(s string) (ColumnType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(SourceColumn):
		return SourceColumn, nil
	case string(TerminalColumn):
		return TerminalColumn, nil
	case string(BrowserColumn):
		return BrowserColumn, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid column type", s)
	}
}

type UsedColumns [10]state.ColumnType
type CurrentColumn = state.ColumnType

type ColumnInfo struct {
	AllUsed UsedColumns
	Focus   CurrentColumn
}

// similar to append() for slice
func appendIfNotExists(columns UsedColumns, colName state.ColumnType) UsedColumns {
	for _, col := range columns {
		if col == colName {
			// if already exists, do nothing
			return columns
		}
	}

	// here we didn't find the column, so append it
	for i, col := range columns {
		if col == "" {
			// columns is copied as an argument, so we can modify it without affecting the caller
			columns[i] = colName
			break
		}
	}

	return columns
}

func resultColumns(current CurrentColumn, prevColumns UsedColumns) state.ColumnFields {
	columns := appendIfNotExists(prevColumns, current)
	return state.ColumnFields{
		FocusColumn: current,
		Column1:     columns[0],
		Column2:     columns[1],
		Column3:     columns[2],
		Column4:     columns[3],
		Column5:     columns[4],
		Column6:     columns[5],
		Column7:     columns[6],
		Column8:     columns[7],
		Column9:     columns[8],
		Column10:    columns[9],
	}
}
