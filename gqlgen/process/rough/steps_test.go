package rough_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/test_util"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestTerminalSteps(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo := test_util.GitOpenOrClone(t, repoUrl)

	cases := []struct {
		inputFile  string
		goldenFile string
		state      *rough.InnerState
	}{
		{"testdata/single/input/terminal1.json", "testdata/single/golden/terminal1.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
		{"testdata/single/input/terminal2.json", "testdata/single/golden/terminal2.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
		{"testdata/single/input/terminal3.json", "testdata/single/golden/terminal3.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
		{"testdata/single/input/terminal4.json", "testdata/single/golden/terminal4.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			test_util.JsonRead(t, c.inputFile, &roughStep)

			// convert to detailed step
			converted, err := c.state.Conversion(&roughStep)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)

			// verify results
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestCommitSteps(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo := test_util.GitOpenOrClone(t, repoUrl)

	cases := []struct {
		inputFile  string
		goldenFile string
		state      *rough.InnerState
	}{
		{"testdata/single/input/manual-commit1.json", "testdata/single/golden/manual-commit1.json", rough.InitStateForUnitTest(repo)},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			test_util.JsonRead(t, c.inputFile, &roughStep)

			// convert to detailed step
			converted, err := c.state.Conversion(&roughStep)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)

			// verify results
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestSourceErrorSteps(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo := test_util.GitOpenOrClone(t, repoUrl)

	cases := []struct {
		inputFile  string
		goldenFile string
		state      *rough.InnerState
	}{
		{"testdata/single/input/source_error1.json", "testdata/single/golden/source_error1.json", rough.InitStateForUnitTest(repo)},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			test_util.JsonRead(t, c.inputFile, &roughStep)

			// convert to detailed step
			converted, err := c.state.Conversion(&roughStep)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)

			// verify results
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestBrowserSteps(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
		state      *rough.InnerState
	}{
		{"testdata/single/input/browser1.json", "testdata/single/golden/browser1.json", rough.InitStateForUnitTest(nil)},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			test_util.JsonRead(t, c.inputFile, &roughStep)

			// convert to detailed step
			converted, err := c.state.Conversion(&roughStep)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)

			// verify results
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestMarkdownSteps(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
		state      *rough.InnerState
	}{
		{"testdata/single/input/markdown1.json", "testdata/single/golden/markdown1.json", rough.InitStateForUnitTest(nil)},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.RoughStep
			test_util.JsonRead(t, c.inputFile, &roughStep)

			// convert to detailed step and verify
			converted, err := c.state.Conversion(&roughStep)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)

			// verify results
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
// 		{"testdata/single/input/terminal1.json", "testdata/single/golden/terminal1.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/single/input/terminal2.json", "testdata/single/golden/terminal2.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/single/input/terminal3.json", "testdata/single/golden/terminal3.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/single/input/terminal4.json", "testdata/single/golden/terminal4.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/single/input/manual-commit1.json", "testdata/single/golden/manual-commit1.json", rough.PredictableInnerState("Terminal", "", repo)},
// 		{"testdata/single/input/source_error1.json", "testdata/single/golden/source_error1.json", rough.PredictableInnerState("Source Code", "", repo)},
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
		{"testdata/rough-steps1.json", "testdata/detailed-steps-golden1.json"},
		{"testdata/rough-steps2.json", "testdata/detailed-steps-golden2.json"},
	}

	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("cannot clone repo %s, %s", repoUrl, err)
	}

	state := rough.InitStateForUnitTest(repo)

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
