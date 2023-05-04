package gitmodel_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
)

func TestFilePath(t *testing.T) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/richardimaoka/gqlgensandbox",
	})
	if err != nil {
		t.Fatalf("error cloning repo: %v", err)
	}

	filePath := "graph/resolver.go"

	f := gitmodel.NewFile(
		repo,
		filePath,
		plumbing.NewHash("a25c09c619f9ed0db2ef05ece3429624b22a0f59"))

	fmt.Println(f)

}
