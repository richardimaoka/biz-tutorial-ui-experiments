package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
)

func TestAppendIfNotExists(t *testing.T) {
	columns := UsedColumns{
		result.TerminalColumnType,
		result.SourceColumnType,
	}
	expected := UsedColumns{
		result.TerminalColumnType,
		result.SourceColumnType,
		result.BrowserColumnType,
	}

	newColumns := appendIfNotExists(columns, result.BrowserColumnType)

	if columns[2] != "" {
		t.Errorf("columns in `appendIfNotExists(columns, )` should not be modified, but got colums[2] = '%s'", columns[2])
	}

	if newColumns != expected {
		t.Errorf("expected %v, but got '%v'", expected, newColumns)
	}
}
