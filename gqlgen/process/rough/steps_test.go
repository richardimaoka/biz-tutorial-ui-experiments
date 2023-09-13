package rough_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestRough(t *testing.T) {
	cases := []struct {
		roughStepFile string
		goldenFile    string
	}{
		{"testdata/rough-step.json", "testdata/detailed-steps-golden.json"},
	}

	for _, c := range cases {
		bytes, err := os.ReadFile(c.roughStepFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		var roughStep rough.RoughStep
		err = json.Unmarshal(bytes, &roughStep)
		if err != nil {
			t.Fatalf("failed to unmarshal json: %v", err)
		}

		result := roughStep.Convert("c8238063-dd5a-4cd4-9d62-5c9c8ebd35ec", []string{})
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
	}
}
