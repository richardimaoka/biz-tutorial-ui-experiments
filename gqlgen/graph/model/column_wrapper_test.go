package model_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
)

func TestMarshalUnmarshal(t *testing.T) {
	cases := []struct {
		name      string
		inputFile string
	}{
		{"terminal fully empty" /*******/, "testdata/column_wrapper_terminal1.json"},
		{"terminal initial terminal" /**/, "testdata/column_wrapper_terminal2.json"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var col model.ColumnWrapper
			err := jsonwrap.Read(c.inputFile, &col)
			if err != nil {
				t.Fatalf("test %s failed, %s", c.name, err)
			}

			testio.CompareWithGoldenFile(t, *updateFlag, c.inputFile, col)
		})
	}
}
