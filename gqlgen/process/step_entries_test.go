package process_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process"
)

func TestReadStepEntries(t *testing.T) {
	cases := []struct {
		dirPath string
	}{
		{"testdata/basic"},
	}

	for _, c := range cases {
		effects, err := process.ReadStepEntries(c.dirPath)
		if err != nil {
			t.Fatalf("ReadStepEntries failed to read file, %s", err)
		}

		internal.CompareWitGoldenFile(t, *updateFlag, c.dirPath+"/golden/step_entries_golden.json", effects)
	}
}

func TestToGraphQLPages(t *testing.T) {
	cases := []struct {
		dirPath string
	}{
		{"testdata/basic"},
	}

	for _, c := range cases {
		effects, err := process.ReadStepEntries(c.dirPath)
		if err != nil {
			t.Fatalf("ReadStepEntries failed to read file, %s", err)
		}

		pages := effects.ToGraphQLPages()
		internal.CompareWitGoldenFile(t, *updateFlag, c.dirPath+"/golden/pages_golden.json", pages)
	}
}
