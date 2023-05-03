package gitmodel_test

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
)

func TestActionCommandMarshal(t *testing.T) {
	f := gitmodel.NewFile(
		"https://github.com/richardimaoka/gqlgensandbox",
		plumbing.NewHash("18ebd6486e9d929f614aba39bd0a2f7bb074d34d"))

	t.Log(f.Size)
	t.Errorf("test")

	if f.Size == 0 {
		t.Errorf("size is 0")
	}

}
