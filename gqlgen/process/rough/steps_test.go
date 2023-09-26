package rough_test

import (
	"reflect"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/test_util"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestTerminalSteps(t *testing.T) {
	cases := []struct {
		inputFile       string
		goldenFile      string
		prevColumn      string
		prevCommit      string
		seqNo           int
		expectedColumns []string
	}{
		{"testdata/rough-steps/terminal1.json", "testdata/golden/terminal1.json", "", "", 0, []string{"Terminal"}},
		{"testdata/rough-steps/terminal2.json", "testdata/golden/terminal2.json", "", "", 0, []string{"Terminal"}},
		{"testdata/rough-steps/terminal3.json", "testdata/golden/terminal3.json", "", "", 0, []string{"Terminal", "Source Code"}},
		{"testdata/rough-steps/terminal4.json", "testdata/golden/terminal4.json", "", "", 0, []string{"Terminal"}},
	}

	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo, err := test_util.GitOpenOrClone(repoUrl)
	if err != nil {
		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			err := internal.JsonRead2(c.inputFile, &roughStep)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			// convert to detailed step and verify
			uuidFinder := rough.StaticUUIDFinder("")
			converted, usedColumns, err := rough.TerminalConvertInternal(&roughStep, repo, uuidFinder, c.prevColumn, c.prevCommit, c.seqNo)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)

			if !reflect.DeepEqual(c.expectedColumns, usedColumns) {
				t.Fatalf("expected %v, but got %v", c.expectedColumns, usedColumns)
			}
		})
	}
}

func TestCommitSteps(t *testing.T) {
	cases := []struct {
		inputFile       string
		goldenFile      string
		expectedColumns []string
	}{
		{"testdata/rough-steps/manual1.json", "testdata/golden/manual1.json", []string{"Source Code"}},
	}

	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo, err := test_util.GitOpenOrClone(repoUrl)
	if err != nil {
		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			err := internal.JsonRead2(c.inputFile, &roughStep)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			// convert to detailed step and verify
			uuidFinder := rough.StaticUUIDFinder("")
			converted, usedColumns, err := rough.CommitConvertInternal(&roughStep, repo, uuidFinder, "", "")
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)

			if !reflect.DeepEqual(c.expectedColumns, usedColumns) {
				t.Fatalf("expected %v, but got %v", c.expectedColumns, usedColumns)
			}
		})
	}
}

func TestSourceErrorSteps(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/rough-steps/source_error1.json", "testdata/golden/source_error1.json"},
	}

	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo, err := test_util.GitOpenOrClone(repoUrl)
	if err != nil {
		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			err := internal.JsonRead2(c.inputFile, &roughStep)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			// convert to detailed step and verify
			uuidFinder := rough.StaticUUIDFinder("")
			converted, err := rough.SourceErrorConvertInternal(&roughStep, repo, uuidFinder)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestBrowserSteps(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/rough-steps/browser1.json", "testdata/golden/browser1.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// 1. read rough step from file
			var roughStep rough.RoughStep
			err := internal.JsonRead2(c.inputFile, &roughStep)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			// 3. convert to detailed step and verify
			uuidFinder := rough.StaticUUIDFinder("")
			converted, err := rough.BrowserConvertInternal(&roughStep, uuidFinder)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestMarkdownSteps(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/rough-steps/markdown1.json", "testdata/golden/markdown1.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// 1. read rough step from file
			var roughStep rough.RoughStep
			err := internal.JsonRead2(c.inputFile, &roughStep)
			if err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			// 3. convert to detailed step and verify
			uuidFinder := rough.StaticUUIDFinder("")
			converted, err := rough.MarkdownConvertInternal(&roughStep, uuidFinder)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

// func TestRoughSteps(t *testing.T) {
// 	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
// 	repo, err := test_util.GitOpenOrClone(repoUrl)
// 	if err != nil {
// 		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
// 	}

// 	cases := []struct {
// 		roughStepFile string
// 		goldenFile    string
// 		InnerState    *rough.InnerState
// 	}{
// 		{"testdata/rough-steps/terminal1.json", "testdata/golden/terminal1.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/rough-steps/terminal2.json", "testdata/golden/terminal2.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/rough-steps/terminal3.json", "testdata/golden/terminal3.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/rough-steps/terminal4.json", "testdata/golden/terminal4.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/rough-steps/manual1.json", "testdata/golden/manual1.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/rough-steps/source_error1.json", "testdata/golden/source_error1.json", rough.PredictableInnerState("Source Code", "", repo)},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.roughStepFile, func(t *testing.T) {
// 			// 1. read rough step from file
// 			var roughStep rough.RoughStep
// 			err := internal.JsonRead2(c.roughStepFile, &roughStep)
// 			if err != nil {
// 				t.Fatalf("failed to unmarshal json: %v", err)
// 			}

// 			// 3. convert to detailed step and verify
// 			converted, err := c.InnerState.Conversion(&roughStep, repo)
// 			if err != nil {
// 				t.Fatalf("failed to convert rough step: %v", err)
// 			}
// 			result := rough.ToOmitEmptyStructs(converted)
// 			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
// 		})
// 	}
// }

func TestRoughStepSequence(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/rough-steps/rough-steps1.json", "testdata/golden/detailed-steps1.json"},
	}

	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
	}

	state := rough.PredictableInnerState("", "", repo)

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			converted, err := state.GenerateTarget(c.inputFile)
			if err != nil {
				t.Fatalf("failed to generate detailed steps: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}
