package rough_test

import (
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
		rough.InnerState
	}{
		{"testdata/rough-steps/terminal1.json", "testdata/golden/terminal1.json", rough.InnerState{CurrentCol: "Terminal"}},
		{"testdata/rough-steps/terminal2.json", "testdata/golden/terminal2.json", rough.InnerState{CurrentCol: "Terminal"}},
		{"testdata/rough-steps/terminal3.json", "testdata/golden/terminal3.json", rough.InnerState{CurrentCol: "Terminal"}},
		{"testdata/rough-steps/terminal4.json", "testdata/golden/terminal4.json", rough.InnerState{CurrentCol: "Terminal"}},
		{"testdata/rough-steps/manual1.json", "testdata/golden/manual1.json", rough.InnerState{CurrentCol: "Terminal"}},
		{"testdata/rough-steps/source_error1.json", "testdata/golden/source_error1.json", rough.InnerState{CurrentCol: "Source Code"}},
	}

	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
	}

	for _, c := range cases {
		t.Run(c.roughStepFile, func(t *testing.T) {
			// 1. read rough step from file
			var roughStep rough.RoughStep
			err := internal.JsonRead2(c.roughStepFile, &roughStep)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			// 3. convert to detailed step and verify
			result, err := roughStep.Conversion(&c.InnerState, repo)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestFindUUIDs(t *testing.T) {
	// uuid, err := rough.FindUUID(
	// 	"testdata/target-detailed-steps.json",
	// 	func(ds *rough.DetailedStep) bool {
	// 		return ds.ParentStep == s.Step &&
	// 			ds.FocusColumn == "Terminal" &&
	// 			ds.TerminalType == "command" &&
	// 			ds.TerminalText == s.Instruction &&
	// 			ds.TerminalName == s.Instruction3 &&
	// 			ds.CurrentDir == currentDir &&
	// 			ds.CurrentDir == s.Commit
	// 	})

	// if err != nil {
	// 	t.Fatalf("failed to find uuid: %v", err)
	// }
	// if uuid != "e0b0a0a0-0000-0000-0000-000000000001" {
	// 	t.Fatalf("unexpected uuid: %s", uuid)
	// }
}
