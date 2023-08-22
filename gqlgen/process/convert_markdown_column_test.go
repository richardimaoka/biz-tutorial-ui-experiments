package process_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

// func TestToStateMarkdownColumn(t *testing.T) {
// 	filepath := "testdata/basic/md_columns.json"
// 	elements, err := read.ReadMarkdownColumns(filepath)
// 	if err != nil {
// 		t.Fatalf("ReadMarkdownColumns failed to read file, %s", err)
// 	}

// 	for i, e := range elements {
// 		col := process.ToStateMarkdownColumn(&e)
// 		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/basic/golden/state_md_col_golden%d.json", i), col)
// 	}
// }

func TestToGraphQLMarkdownColumn(t *testing.T) {
	filepath := "testdata/basic/md_columns.json"
	elements, err := read.ReadMarkdownColumns(filepath)
	if err != nil {
		t.Fatalf("ReadMarkdownColumns failed to read file, %s", err)
	}

	for i, e := range elements {
		col := process.ToGraphQLMarkdownColumn(&e)
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/basic/golden/graphql_md_col_golden%d.json", i), col)
	}
}
