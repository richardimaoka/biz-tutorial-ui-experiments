package processing_test

import (
	"os"
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
			{FromLine: 2, ToLine: 3},
			{FromLine: 5, ToLine: 5},
		},
	}}

	for _, c := range cases {
		t.Run(c.before, func(t *testing.T) {
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
