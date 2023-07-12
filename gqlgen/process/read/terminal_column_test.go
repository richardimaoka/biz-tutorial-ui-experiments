package read_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

func TestReadTerminalColumn(t *testing.T) {
	filepath := "testdata/terminal_columns.json"
	elements, err := read.ReadTerminalColumns(filepath)
	if err != nil {
		t.Fatalf("ReadTerminalColumns failed to read file, %s", err)
	}

	for i, e := range elements {
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/golden/terminal_col_entry_golden%d.json", i), e)
	}
}
