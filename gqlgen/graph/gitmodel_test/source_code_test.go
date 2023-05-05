package gitmodel_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
)

var update = flag.Bool("update", false, "update golden files")

func TestSourceCodeFromGit(t *testing.T) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/richardimaoka/gqlgensandbox",
	})
	if err != nil {
		t.Fatalf("error cloning repo: %v", err)
	}

	type TestCase struct {
		FilePath        string
		IsUpdated       bool
		CommitHash      string
		ExpectationFile string
	}

	testCases := []TestCase{
		{
			FilePath:        "graph/resolver.go",
			IsUpdated:       false,
			CommitHash:      "4dd8f51d6acbee9d61b24dc26715ecc48a5d2456",
			ExpectationFile: "testdata/source_code/4dd8f51d6acbee9d61b24dc26715ecc48a5d2456.json",
		},
	}

	for _, c := range testCases {
		t.Run(c.FilePath, func(t *testing.T) {
			sc, err := gitmodel.NewSourceCodeFromGit(repo, plumbing.NewHash(c.CommitHash))
			if err != nil {
				t.Fatalf("error creating source code: %v", err)
			}

			if *update {
				internal.WriteGoldenFile(t, c.ExpectationFile, sc)
			}
		})
	}
	fmt.Println(testCases)
}
