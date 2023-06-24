package internal_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing/internal"
)

func TestFileHighlight(t *testing.T) {
	cases := []struct {
		before   string
		after    string
		expected []internal.FileHighlight
	}{{
		before: "testdata/file_highlight/1-1.txt", after: "testdata/file_highlight/1-2.txt", expected: []internal.FileHighlight{
			{FromLine: 2, ToLine: 3},
			{FromLine: 5, ToLine: 5},
		},
	}, {
		before: "testdata/file_highlight/1-2.txt", after: "testdata/file_highlight/1-3.txt", expected: []internal.FileHighlight{
			{FromLine: 3, ToLine: 3},
			{FromLine: 5, ToLine: 7},
		},
	}, {
		before: "testdata/file_highlight/1-3.txt", after: "testdata/file_highlight/1-4.txt", expected: []internal.FileHighlight{
			{FromLine: 1, ToLine: 3},
		},
	}, {
		before: "testdata/file_highlight/2-1.txt", after: "testdata/file_highlight/2-2.txt", expected: []internal.FileHighlight{
			{FromLine: 1, ToLine: 4},
			{FromLine: 6, ToLine: 7},
		},
	}, {
		before: "testdata/file_highlight/3-1.txt", after: "testdata/file_highlight/3-2.txt", expected: []internal.FileHighlight{
			{FromLine: 1, ToLine: 2},
			{FromLine: 4, ToLine: 6},
		},
	}, {
		before: "testdata/file_highlight/4-1.txt", after: "testdata/file_highlight/4-2.txt", expected: []internal.FileHighlight{
			{FromLine: 2, ToLine: 8},
		},
	}, {
		before: "testdata/file_highlight/4-2.txt", after: "testdata/file_highlight/4-3.txt", expected: []internal.FileHighlight{
			{FromLine: 7, ToLine: 10},
		},
	}, {
		before: "testdata/file_highlight/5-1.txt", after: "testdata/file_highlight/5-2.txt", expected: []internal.FileHighlight{
			{FromLine: 3, ToLine: 5},
		},
	}, {
		before: "testdata/file_highlight/6-1.go", after: "testdata/file_highlight/6-2.go", expected: []internal.FileHighlight{
			{FromLine: 25, ToLine: 25},
			{FromLine: 67, ToLine: 68},
			{FromLine: 71, ToLine: 72},
		},
	}, {
		before: "testdata/file_highlight/7-1.go.mod", after: "testdata/file_highlight/7-2.go.mod", expected: []internal.FileHighlight{
			{FromLine: 4, ToLine: 20},
		},
	}}

	for _, c := range cases {
		folder := "testdata/file_highlight/"
		before := strings.Replace(strings.Replace(c.before, folder, "", 1), ".txt", "", 1)
		after := strings.Replace(strings.Replace(c.after, folder, "", 1), ".txt", "", 1)
		name := fmt.Sprintf("%s vs. %s", before, after)

		t.Run(name, func(t *testing.T) {
			beforeBytes, err := os.ReadFile(c.before)
			if err != nil {
				t.Fatalf("Failed to read file %s, %v", c.before, err)
			}
			afterBytes, err := os.ReadFile(c.after)
			if err != nil {
				t.Fatalf("Failed to read file %s, %v", c.after, err)
			}

			results := internal.CalcHighlight(string(beforeBytes), string(afterBytes))
			if diff := cmp.Diff(c.expected, results); diff != "" {
				t.Fatalf("mismatch (-expected +result):\n%s", diff)
			}
		})
	}
}
