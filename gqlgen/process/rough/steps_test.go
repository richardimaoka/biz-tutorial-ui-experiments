package rough_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestTerminalSteps(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo := testio.GitOpenOrClone(t, repoUrl)

	cases := []struct {
		inputFile  string
		goldenFile string
		state      *rough.InnerState
	}{
		// terminal step tests
		// {"testdata/single/input/terminal1.json", "testdata/single/golden/terminal1.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
		// {"testdata/single/input/terminal2.json", "testdata/single/golden/terminal2.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
		// {"testdata/single/input/terminal3.json", "testdata/single/golden/terminal3.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
		// {"testdata/single/input/terminal4.json", "testdata/single/golden/terminal4.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},
		// {"testdata/single/input/terminal5.json", "testdata/single/golden/terminal5.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},

		// // terminal command steps
		// {"testdata/single/input/terminal-command1.json", "testdata/single/golden/terminal-command1.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},

		// terminal output steps
		{"testdata/single/input/terminal-output1.json", "testdata/single/golden/terminal-output1.json", rough.StateWithColumnForUnitTest(repo, "Terminal")},

		// // commit step tests
		// {"testdata/single/input/manual-commit1.json", "testdata/single/golden/manual-commit1.json", rough.InitStateForUnitTest(repo)},

		// // source error step tests
		// {"testdata/single/input/source_error1.json", "testdata/single/golden/source_error1.json", rough.InitStateForUnitTest(repo)},

		// // browser step tests
		// {"testdata/single/input/browser1.json", "testdata/single/golden/browser1.json", rough.InitStateForUnitTest(nil)},

		// // markdown step tests
		// {"testdata/single/input/markdown1.json", "testdata/single/golden/markdown1.json", rough.InitStateForUnitTest(nil)},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			// read rough step from file
			var roughStep rough.ResultStep
			testio.JsonRead(t, c.inputFile, &roughStep)

			// convert to detailed step and verify
			converted, err := c.state.Conversion(&roughStep)
			if err != nil {
				t.Fatalf("failed to convert rough step: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)

			// verify results
			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestRoughStepSequence(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/article-gqlgen-getting-started"
	repo := testio.GitOpenOrClone(t, repoUrl)

	cases := []struct {
		inputFile  string
		goldenFile string
		state      *rough.InnerState
	}{
		{"testdata/rough-steps1.json", "testdata/detailed-steps-golden1.json", rough.InitStateForUnitTest(repo)},
		{"testdata/rough-steps2.json", "testdata/detailed-steps-golden2.json", rough.InitStateForUnitTest(repo)},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			converted, err := c.state.GenerateTarget(c.inputFile)
			if err != nil {
				t.Fatalf("failed to generate detailed steps: %v", err)
			}
			result := rough.ToOmitEmptyStructs(converted)
			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}
