package state_test

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestRootDirectory(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	repo, err := gitOpenOrClone(repoUrl)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		currentCommit string
		goldenFile    string
	}{
		{
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"testdata/directory_golden.json",
		},
	}

	for _, c := range cases {
		currentCommit, err := repo.CommitObject(plumbing.NewHash(c.currentCommit))
		if err != nil {
			t.Fatal(err)
		}
		currentRoot, err := currentCommit.Tree()
		if err != nil {
			t.Fatal(err)
		}
		rootDir := state.EmptyDirectory(repo, "")
		if err := rootDir.Recurse("", currentRoot); err != nil {
			t.Fatal(err)
		}

		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, rootDir.ToGraphQLFileNodeSlice())
	}
}
