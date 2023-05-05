package gitmodel_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
)

func TestFilePath(t *testing.T) {
	type TestCase struct {
		FilePath        string
		IsUpdated       bool
		ExpectationFile string
	}

	testCases := []TestCase{
		{
			FilePath:        "graph/resolver.go",
			IsUpdated:       false,
			ExpectationFile: "testdata/resolver.go.json",
		}, {
			FilePath:        "server.go",
			IsUpdated:       true,
			ExpectationFile: "testdata/server.go.json",
		},
	}

	for _, c := range testCases {
		t.Run(c.FilePath, func(t *testing.T) {
			f, err := gitmodel.NewFileFromGit(c.FilePath, c.IsUpdated)
			if err != nil {
				t.Fatalf("error creating file: %v", err)
			}
			internal.CompareAfterMarshal(t, c.ExpectationFile, f.FileNode())
		})
	}
}
