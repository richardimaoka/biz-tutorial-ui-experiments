package gitmodel_test

import (
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
)

func TestCommitTest(t *testing.T) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/richardimaoka/gqlgensandbox",
	})
	if err != nil {
		t.Fatalf("error cloning repo: %v", err)
	}

	f := gitmodel.NewFile(
		repo,
		plumbing.NewHash("046d10917933eadce9b04880b6bc5d99c1ce9637"))

	bytes, err := os.ReadFile("./testdata/go.046d10917933eadce9b04880b6bc5d99c1ce9637.mod")
	if err != nil {
		t.Errorf("error reading ./testdata/go.046d10917933eadce9b04880b6bc5d99c1ce9637.mod")
	}

	contents := f.Contents()
	t.Error(*contents)
	t.Error(string(bytes))

	if contents := f.Contents(); *contents != string(bytes) {
		t.Errorf("contents mismatched")
	}
}
