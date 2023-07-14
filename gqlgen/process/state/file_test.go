package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func FindCommitFile(repo *git.Repository, commitHashStr, filePath string) (*object.File, error) {
	commitHash := plumbing.NewHash(commitHashStr)
	if commitHash.String() != commitHashStr {
		return nil, fmt.Errorf("failed in FindCommitFile, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())
	}

	//commit 'npx create-next-app@latest' =
	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in FindCommitFile, cannot get commit = %s, %s", commitHashStr, err)
	}

	rootTree, err := commit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in FindCommitFile, cannot get tree for commit = %s, %s", commitHashStr, err)

	}

	file, err := rootTree.File(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed in FindCommitFile, cannot get file = %s in commit = %s, %s", filePath, commitHashStr, err)
	}

	return file, nil
}

func TestFileMutation1(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("failed in NewSourceCodeColumn, cannot clone repo %s, %s", repoUrl, err)
	}

	// commit = 'npx create-next-app@latest'
	gitFile1, err := FindCommitFile(repo, "8adac375628219e020d4b5957ff24f45954cbd3f", "next/package.json")
	if err != nil {
		t.Fatalf("failed in FindCommitFile, %s", err)
	}

	// commit = 'npm install --save @emotion/react', which updates package.json
	gitFile2, err := FindCommitFile(repo, "8adac375628219e020d4b5957ff24f45954cbd3f", "next/README.md")
	if err != nil {
		t.Fatalf("failed in FindCommitFile, %s", err)
	}

	s, err := state.NewFile(repo, nil, gitFile1)
	if err != nil {
		t.Fatalf("failed in NewFile, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLOpenFileNode()
	goldenFile1 := "testdata/file_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.PrevFile = gitFile1
	s.CurrentFile = gitFile2

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/file_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLOpenFileNode())
}

func TestFileMutation2(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("failed in NewSourceCodeColumn, cannot clone repo %s, %s", repoUrl, err)
	}

	// commit = 'npx create-next-app@latest'
	gitFile1, err := FindCommitFile(repo, "8adac375628219e020d4b5957ff24f45954cbd3f", "next/package.json")
	if err != nil {
		t.Fatalf("failed in FindCommitFile, %s", err)
	}

	s, err := state.NewFile(repo, nil, gitFile1)
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
