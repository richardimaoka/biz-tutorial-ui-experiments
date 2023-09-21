package rough_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestUUIDGenerator(t *testing.T) {
	cases := []struct {
		expectedUUID  string
		roughStepFile string
		subId         string
	}{
		{"73388488-7f44-4617-9d6f-5cacad5bcaf8", "testdata/rough-steps/terminal1.json", "terminalCommandStep"},
		{"96d5c1df-5488-4601-9ed2-5387f3e8d5f8", "testdata/rough-steps/terminal1.json", "fileTreeStep"},
		{"d748e6d7-124a-4889-92eb-6a2226272ea8", "testdata/rough-steps/terminal1.json", "openFileStep-0"},
		{"d748e6d7-124a-4889-92eb-6a2226272ea8", "testdata/rough-steps/terminal1.json", "non-existeint-subid"},
	}

	targetFile := "testdata/target-detailed-steps.json"
	gen, err := rough.NewUUIDGenerator(targetFile)
	if err != nil {
		t.Fatalf("failed to create UUIDGenerator: %v", err)
	}

	for _, c := range cases {
		t.Run(c.roughStepFile, func(t *testing.T) {
			var roughStep rough.RoughStep
			err := internal.JsonRead2(c.roughStepFile, &roughStep)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			id := gen.FindOrGenerateUUID(&roughStep, c.subId)
			if id != c.expectedUUID {
				t.Fatalf("expected %s, but got %s", c.expectedUUID, id)
			}
		})
	}
}
