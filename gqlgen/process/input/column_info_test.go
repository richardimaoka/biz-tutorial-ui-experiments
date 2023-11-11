package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state2"
)

func TestAppendIfNotExists(t *testing.T) {
	columns := UsedColumns{
		state2.TerminalColumnType,
		state2.SourceColumnType,
	}
	expected := UsedColumns{
		state2.TerminalColumnType,
		state2.SourceColumnType,
		state2.BrowserColumnType,
	}

	newColumns := appendIfNotExists(columns, state2.BrowserColumnType)

	if columns[2] != "" {
		t.Errorf("columns in `appendIfNotExists(columns, )` should not be modified, but got colums[2] = '%s'", columns[2])
	}

	if newColumns != expected {
		t.Errorf("expected %v, but got '%v'", expected, newColumns)
	}
}
