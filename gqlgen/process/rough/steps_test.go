package rough_test

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestRough(t *testing.T) {
	// repoURL := "https://github.com/richardimaoka/gqlgen-getting-started.git"

	cases := []struct {
		roughStepFile string
		goldenFile    string
	}{
		{"testdata/rough-step1.json", "testdata/detailed-steps-golden1.json"},
		{"testdata/rough-step2.json", "testdata/detailed-steps-golden2.json"},
	}

	uuidFile, err := os.Open("testdata/uuids.txt")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	r := bufio.NewReader(uuidFile)
	for _, c := range cases {
		// 1. read UUID from file
		line, err := r.ReadString('\n')
		if err != nil {
			t.Fatalf("failed to read uuid: %v", err)
		}
		line = strings.TrimSuffix(line, "\n")
		stepId, err := uuid.Parse(line)
		if err != nil {
			t.Fatalf("failed to parse uuid = '%s': %v", line, err)
		}

		// 2. read rough step from file
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
		result, err := roughStep.Convert(stepId.String(), []string{})
		if err != nil {
			t.Fatalf("failed to convert rough step: %v", err)
		}
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
	}
}
