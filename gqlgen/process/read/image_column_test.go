package read_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

func TestReadImgColumn(t *testing.T) {
	filepath := "testdata/bg_columns.json"
	entries, err := read.ReadImageDescriptionColumns(filepath)
	if err != nil {
		t.Fatalf("ReadImageDescriptionColumns failed to read file, %s", err)
	}

	for i, e := range entries {
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/golden/img_col_entry_golden%d.json", i), e)
	}
}
