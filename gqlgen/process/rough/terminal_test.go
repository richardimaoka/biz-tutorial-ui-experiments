package rough_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestRoughCommands(t *testing.T) {
	cases := []struct {
		roughStepFile string
		goldenFile    string
	}{
		{"testdata/rough-steps/terminal1.json", "testdata/golden/terminal1.json"},
		{"testdata/rough-steps/terminal2.json", "testdata/golden/terminal2.json"},
		{"testdata/rough-steps/terminal3.json", "testdata/golden/terminal3.json"},
		{"testdata/rough-steps/terminal4.json", "testdata/golden/terminal4.json"},
		{"testdata/rough-steps/manual1.json", "testdata/golden/manual1.json"},
		{"testdata/rough-steps/source_error1.json", "testdata/golden/source_error1.json"},
	}

	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
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
		result, err := roughStep.Conversion(&rough.InnerState{}, repo)
		if err != nil {
			t.Fatalf("failed to convert rough step: %v", err)
		}
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
	}
}
