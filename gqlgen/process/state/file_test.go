package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func fileFromCommit(repoUrl, commitHashStr, filePath string) (*state.File, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCodeColumn, cannot clone repo %s, %s", repoUrl, err)
	}

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

	gitFile, err := rootTree.File(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed in FindCommitFile, cannot get file = %s in commit = %s, %s", filePath, commitHashStr, err)
	}

	fileState, err := state.NewFile(repo, nil, gitFile)
	if err != nil {
		return nil, fmt.Errorf("failed in FindCommitFile, cannot create file state for file = %s in commit = %s, %s", filePath, commitHashStr, err)
	}

	return fileState, nil
}

// since state.File is effectively immutable, no need to test the state mutation, but only the GraphQL model mutation
func TestFileMutation1(t *testing.T) {
	s, err := fileFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestFileMutation1, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel, err := s.ToGraphQLOpenFile()
	if err != nil {
		t.Fatalf("failed in TestFileMutation1, %s", err)
	}
	goldenFile1 := "testdata/file_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	*gqlModel.FilePath = "next/README.md"
	*gqlModel.FileName = "README.md"
	*gqlModel.IsFullContent = false
	*gqlModel.Content = "updated contents"
	//*gqlModel.Language = "markdown"
	line100 := 100
	line200 := 200
	highlight := model.FileHighlight{
		FromLine: &line100,
		ToLine:   &line200,
	}
	gqlModel.Highlight = append(gqlModel.Highlight, &highlight)

	// ... has NO effect on a RE-materialized GraphQL model
	gqlModelReMat, err := s.ToGraphQLOpenFile()
	if err != nil {
		t.Fatalf("failed in TestFileMutation1, %s", err)
	}
	internal.CompareAfterMarshal(t, goldenFile1, gqlModelReMat)

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/file_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}

// since state.File is effectively immutable, no need to test the state mutation, but only the GraphQL model mutation
func TestFileMutation2(t *testing.T) {
	s, err := fileFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in FindCommitFile, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLFileNode()
	goldenFile1 := "testdata/file_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	*gqlModel.Name = "README.md"
	*gqlModel.FilePath = "next/README.md"
	*gqlModel.Offset = 5
	*gqlModel.IsUpdated = !*gqlModel.IsUpdated
	*gqlModel.NodeType = model.FileNodeTypeDirectory

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLFileNode())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/file_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
