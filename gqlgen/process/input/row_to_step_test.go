package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
)

func TestToSteps(t *testing.T) {
	cases := []struct {
		name       string
		inputFile  string
		goldenFile string
	}{
		{"successful", "testdata/rows/docker-tutorial.json", "testdata/rows/docker-tutorial-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Prepare arguments
			repoUrl := "https://github.com/richardimaoka/article-docker-cmd-entrypoint.git"
			repo := testio.GitOpenOrClone(t, repoUrl)

			// Generate non-empty step IDs to make sure target file's IDs are indeed picked up
			// Golden file is the target file
			finder, _ := NewFinder(c.goldenFile)

			var rows []Row
			testio.JsonRead(t, c.inputFile, &rows)

			// Function to test
			_, err := toSteps(rows, finder, repo)
			if err != nil {
				t.Errorf("TestToSteps failed, %s", err)
			}

			// do not check the results. don't want to check the implementation
			// if no error, then it's good enough
		})
	}
}
