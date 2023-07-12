package read_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

func TestReadSteps(t *testing.T) {
	filepath := "testdata/steps.json"
	elements, err := read.ReadSteps(filepath)
	if err != nil {
		t.Fatalf("ReadSteps failed to read file, %s", err)
	}

	for i, e := range elements {
		internal.CompareWitGoldenFile(t, *updateFlag, fmt.Sprintf("testdata/step_entry_golden%d.json", i), e)
	}
}
