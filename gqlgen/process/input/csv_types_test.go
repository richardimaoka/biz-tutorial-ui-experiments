package input_test

import (
	"encoding/json"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/input"
)

func TestCsvString(t *testing.T) {
	cases := []struct {
		expected   input.CsvString
		inputBytes []byte
	}{
		{"1", []byte(`1`)},
		{"2", []byte(`2`)},
		{"abc", []byte(`"abc"`)},
	}

	for _, c := range cases {
		t.Run(string(c.expected), func(t *testing.T) {
			var result input.CsvString
			err := json.Unmarshal(c.inputBytes, &result)

			if err != nil {
				t.Fatalf("failed %s", err)
			}

			if result != c.expected {
				t.Fatalf("CSV string is expected = '%s' but got '%s'", c.expected, result)
			}

		})
	}
}
