package process_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
)

func TestReadStepEntries(t *testing.T) {
	effects, err := process.ReadStepEntries("testdata/basic")
	if err != nil {
		t.Fatalf("ReadStepEntries failed to read file, %s", err)
	}

	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/golden/basic/step_entries_golden.json", effects)
}

func TestToGraphQLPages(t *testing.T) {
	effects, err := process.ReadStepEntries("testdata/basic")
	if err != nil {
		t.Fatalf("ReadStepEntries failed to read file, %s", err)
	}

	pages := effects.ToGraphQLPages()
	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/basic/golden/pages_golden.json", pages)
}
