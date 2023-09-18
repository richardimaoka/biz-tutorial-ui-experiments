package rough_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestRoughCommands(t *testing.T) {
	// repoURL := "https://github.com/richardimaoka/gqlgen-getting-started.git"

	cases := []struct {
		roughStepFile string
		goldenFile    string
	}{
		{"testdata/rough-step-command1.json", "testdata/detailed-step-command-golden1.json"},
	}

	for _, c := range cases {
		// 1. read rough step from file
		bytes, err := os.ReadFile(c.roughStepFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		var roughStep rough.RoughStep
		err = json.Unmarshal(bytes, &roughStep)
		if err != nil {
			t.Fatalf("failed to unmarshal json: %v", err)
		}

		// 3. convert to detailed step and verify
		result, err := roughStep.TerminalConvert(&rough.InnerState{})
		if err != nil {
			t.Fatalf("failed to convert rough step: %v", err)
		}
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
	}
}
