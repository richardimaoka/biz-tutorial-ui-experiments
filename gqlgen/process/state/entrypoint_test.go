package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func Test(t *testing.T) {
	cases := []struct {
		tutorialName string
		repoUrl      string
	}{
		{"docker-tutorial", "https://github.com/richardimaoka/article-docker-cmd-entrypoint.git"},
	}

	for _, c := range cases {
		t.Run(c.tutorialName, func(t *testing.T) {
			// Prepare arguments
			repo := testio.GitOpenOrClone(t, c.repoUrl)
			inputFile := "testdata/" + c.tutorialName + "/steps.json"

			// read data from file
			var steps []state.Step
			testio.JsonRead(t, inputFile, &steps)

			// Function to test
			page := state.NewPage(repo, c.tutorialName)
			for _, step := range steps {
				t.Run(c.tutorialName, func(t *testing.T) {
					err := page.Update(&step)
					if err != nil {
						t.Errorf("failed to update page, %s", err)
					}
					gqlModel := page.ToGraphQL()
					goldenFile := "testdata/" + c.tutorialName + "/" + step.StepId + ".json"
					testio.CompareWithGoldenFile(t, *updateFlag, goldenFile, gqlModel)
				})
			}
		})
	}
}
