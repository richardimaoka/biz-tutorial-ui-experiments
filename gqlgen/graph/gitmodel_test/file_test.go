package gitmodel_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
)

func TestFilePath(t *testing.T) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/richardimaoka/gqlgensandbox",
	})
	if err != nil {
		t.Fatalf("error cloning repo: %v", err)
	}

	type TestCase struct {
		FilePath        string
		CommitHash      string
		BlobHash        string
		ExpectationFile string
	}

	testCases := []TestCase{
		{
			FilePath:        "graph/resolver.go",
			CommitHash:      "4dd8f51d6acbee9d61b24dc26715ecc48a5d2456",
			BlobHash:        "a25c09c619f9ed0db2ef05ece3429624b22a0f59",
			ExpectationFile: "testdata/resolver.go.json",
		}, {
			FilePath:        "server.go",
			CommitHash:      "4dd8f51d6acbee9d61b24dc26715ecc48a5d2456",
			BlobHash:        "c0f5b2c8f1bf0ebbf713cc213b378af90a14f061",
			ExpectationFile: "testdata/server.go.json",
		},
	}

	for _, c := range testCases {
		t.Run(c.FilePath, func(t *testing.T) {
			f, err := gitmodel.NewFileFromGit(
				repo,
				c.FilePath,
				plumbing.NewHash(c.BlobHash),
				plumbing.NewHash(c.CommitHash))

			if err != nil {
				t.Fatalf("error creating file: %v", err)
			}
			internal.CompareAfterMarshal(t, c.ExpectationFile, f.FileNode())
		})
	}
}
