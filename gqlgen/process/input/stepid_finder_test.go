package input_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/input"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
)

func TestFinder(t *testing.T) {
	cases := []struct {
		expectedUUID string
		inputFile    string
		subId        string
	}{
		{"73388488-7f44-4617-9d6f-5cacad5bcaf8", "testdata/finder/terminal1.json", "terminalCommandStep"},
		{"96d5c1df-5488-4601-9ed2-5387f3e8d5f8", "testdata/finder/terminal1.json", "fileTreeStep"},
		{"d748e6d7-124a-4889-92eb-6a2226272ea8", "testdata/finder/terminal1.json", "openFileStep-0"},
		{"", "testdata/finder/terminal1.json", "non-existeint-subid"},
	}

	targetFile := "testdata/finder/target-detailed-steps.json"
	finder := input.PredictableFinder(t, targetFile)

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var result result.Step
			err := jsonwrap.Read(c.inputFile, &result)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			id := finder.StepIdFor(result.StepId, c.subId)
			if id != c.expectedUUID {
				t.Fatalf("expected %s, but got %s", c.expectedUUID, id)
			}
		})
	}
}
