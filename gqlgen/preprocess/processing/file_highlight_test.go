package processing_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

func TestFileHighlight(t *testing.T) {
	cases := []struct {
		before   string
		after    string
		expected []processing.FileHighlight
	}{{
		before: "testdata/file_highlight/1-1.txt", after: "testdata/file_highlight/1-2.txt", expected: []processing.FileHighlight{
			{FromLine: 2, ToLine: 2},
			{FromLine: 3, ToLine: 3},
			{FromLine: 5, ToLine: 5},
		},
	}, {
		before: "testdata/file_highlight/1-2.txt", after: "testdata/file_highlight/1-3.txt", expected: []processing.FileHighlight{
			{FromLine: 3, ToLine: 3},
			{FromLine: 5, ToLine: 7},
		},
	}, {
		before: "testdata/file_highlight/1-3.txt", after: "testdata/file_highlight/1-4.txt", expected: []processing.FileHighlight{
			{FromLine: 1, ToLine: 1},
			{FromLine: 1, ToLine: 1},
			{FromLine: 2, ToLine: 2},
			{FromLine: 3, ToLine: 3},
		},
	}, {
		before: "testdata/file_highlight/2-1.txt", after: "testdata/file_highlight/2-2.txt", expected: []processing.FileHighlight{
			{FromLine: 1, ToLine: 1},
			{FromLine: 2, ToLine: 2},
			{FromLine: 3, ToLine: 3},
			{FromLine: 4, ToLine: 4},
			{FromLine: 4, ToLine: 4},
			{FromLine: 4, ToLine: 4},
			{FromLine: 6, ToLine: 6},
			{FromLine: 7, ToLine: 7},
		},
	}, {
		before: "testdata/file_highlight/3-1.txt", after: "testdata/file_highlight/3-2.txt", expected: []processing.FileHighlight{
			{FromLine: 1, ToLine: 1},
			{FromLine: 1, ToLine: 1},
			{FromLine: 2, ToLine: 2},
			{FromLine: 4, ToLine: 6},
		},
	}, {
		before: "testdata/file_highlight/4-1.txt", after: "testdata/file_highlight/4-2.txt", expected: []processing.FileHighlight{
			{FromLine: 2, ToLine: 2},
			{FromLine: 2, ToLine: 2},
			{FromLine: 3, ToLine: 3},
			{FromLine: 4, ToLine: 4},
			{FromLine: 5, ToLine: 8},
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

			results := processing.CalcHighlight(string(beforeBytes), string(afterBytes))
			if diff := cmp.Diff(c.expected, results); diff != "" {
				t.Fatalf("mismatch (-expected +result):\n%s", diff)
			}
		})
	}
}
