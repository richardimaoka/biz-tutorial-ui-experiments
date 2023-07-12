package read_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

func TestReadGitColumn(t *testing.T) {
	filepath := "testdata/git_columns.json"
	elements, err := read.ReadGitColumns(filepath)
	if err != nil {
		t.Fatalf("ReadGitColumns failed to read file, %s", err)
	}

	for i, e := range elements {
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/golden/git_col_entry_golden%d.json", i), e)
	}
}
