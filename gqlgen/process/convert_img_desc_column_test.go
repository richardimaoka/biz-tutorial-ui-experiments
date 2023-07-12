package process_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

func TestToStateImgDescColumn(t *testing.T) {
	filepath := "testdata/img_columns.json"
	entries, err := read.ReadImageDescriptionColumns(filepath)
	if err != nil {
		t.Fatalf("TestToStateBgColumn failed to read file, %s", err)
	}

	for i, e := range entries {
		col := process.ToStateImgDescColumn(e)
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/golden/state_img_desc_col_golden%d.json", i), col)
	}
}
