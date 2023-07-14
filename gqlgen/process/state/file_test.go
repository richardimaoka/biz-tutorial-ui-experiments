package state_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

// since state.File is effectively immutable, no need to test the state mutation, but only the GraphQL model mutation
func TestFileMutation(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("failed in NewSourceCodeColumn, cannot clone repo %s, %s", repoUrl, err)
	}

	//commit 'npx create-next-app@latest' =
	commitHash := "8adac375628219e020d4b5957ff24f45954cbd3f"
	commit, err := repo.CommitObject(plumbing.NewHash(commitHash))
	if err != nil {
		t.Fatalf("failed in FindCommitFile, cannot get commit = %s, %s", commit.Hash, err)
	}

	rootTree, err := commit.Tree()
	if err != nil {
		t.Fatalf("failed in FindCommitFile, cannot get tree for commit = %s, %s", commitHash, err)

	}

	filePath := "next/package.json"
	gitFile, err := rootTree.File(filePath)
	if err != nil {
		t.Fatalf("failed in FindCommitFile, cannot get file = %s in commit = %s, %s", filePath, commitHash, err)
	}

	s, err := state.NewFile(repo, nil, gitFile)
	if err != nil {
		t.Fatalf("failed in NewFile, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLOpenFileNode()
	goldenFile1 := "testdata/file_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	*gqlModel.Name = "README.md"
	*gqlModel.FilePath = "next/README.md"

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLOpenFileNode())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/file_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
