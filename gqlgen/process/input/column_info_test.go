package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestAppendIfNotExists(t *testing.T) {
	columns := UsedColumns{
		state.TerminalColumnType,
		state.SourceColumnType,
	}
	expected := UsedColumns{
		state.TerminalColumnType,
		state.SourceColumnType,
		state.BrowserColumnType,
	}

	newColumns := appendIfNotExists(columns, state.BrowserColumnType)

	if columns[2] != "" {
		t.Errorf("columns in `appendIfNotExists(columns, )` should not be modified, but got colums[2] = '%s'", columns[2])
	}

	if newColumns != expected {
		t.Errorf("expected %v, but got '%v'", expected, newColumns)
	}
}
