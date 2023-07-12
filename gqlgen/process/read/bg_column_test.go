package read_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

// testing the field set - changes to the spreadsheed or structure will be captured by this
func TestReadBgColumn(t *testing.T) {
	filepath := "testdata/bg_columns.json"
	entries, err := read.ReadBackgroundImageColumns(filepath)
	if err != nil {
		t.Fatalf("ReadBackgroundImageColumns failed to read file, %s", err)
	}

	for i, e := range entries {
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/bg_col_eff_golden%d.json", i), e)
	}
}
