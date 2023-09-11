package rough_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestRough(t *testing.T) {
	filename := "testdata/rough-steps.json"
	goldenFile := "testdata/detailed-steps-golden.json"

	detailedSteps := rough.GenDetailedSteps(filename)
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile, detailedSteps)
}
