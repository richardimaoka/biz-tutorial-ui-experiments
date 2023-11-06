package input

import "testing"

func TestAppendIfNotExists(t *testing.T) {
	columns := UsedColumns{"Terminal", "Source Code"}

	expected := UsedColumns{"Terminal", "Source Code", "Browser"}
	newColumns := appendIfNotExists(columns, expected[2])

	if columns[2] != "" {
		t.Errorf("columns in `appendIfNotExists(columns, )` should not be modified, but got colums[2] = '%s'", columns[2])
	}

	if newColumns != expected {
		t.Errorf("expected %v, but got '%v'", expected, newColumns)
	}
}
