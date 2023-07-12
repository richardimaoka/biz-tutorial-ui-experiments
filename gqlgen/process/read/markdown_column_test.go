package read_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

func TestReadMarkdownColumn(t *testing.T) {
	filepath := "testdata/md_columns.json"
	elements, err := read.ReadMarkdownColumns(filepath)
	if err != nil {
		t.Fatalf("ReadMarkdownColumns failed to read file, %s", err)
	}

	for i, e := range elements {
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/md_col_entry_golden%d.json", i), e)
	}
}
