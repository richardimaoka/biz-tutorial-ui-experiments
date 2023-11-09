package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
)

func TestToSteps(t *testing.T) {
	cases := []struct {
		name       string
		inputFile  string
		targetFile string
		goldenFile string
	}{
		{"successful", "testdata/rows/docker-tutorial.json", "", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Prepare arguments
			repo := testio.GitOpenOrClone(t, "https://github.com/richardimaoka/article-docker-cmd-entrypoint.git")
			finder := PredictableFinder(t, c.targetFile)
			var rows []Row
			testio.JsonRead(t, c.inputFile, &rows)

			// Function to test
			steps, err := toSteps(rows, finder, repo)
			if err != nil {
				t.Errorf("TestToSteps failed, %s", err)
			}

			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, steps)
		})
	}
}
