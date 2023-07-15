package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func gitOpenOrClone(repoUrl string) (*git.Repository, error) {
	repo, err := git.Open(storage, nil)

	// if failed to open, then try cloning
	if err != nil {
		repo, err = git.Clone(storage, nil, &git.CloneOptions{URL: repoUrl})
		if err != nil {
			return nil, fmt.Errorf("cannot clone repo %s, %s", repoUrl, err)
		}
	}

	return repo, nil
}

func gitFileFromCommit(repoUrl, commitHashStr, filePath string) (*object.File, error) {
	repo, err := gitOpenOrClone(repoUrl)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, %s", err)
	}

	commitHash := plumbing.NewHash(commitHashStr)
	if commitHash.String() != commitHashStr {
		return nil, fmt.Errorf("failed in gitFileFromCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get commit = %s, %s", commitHashStr, err)
	}

	rootTree, err := commit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get tree for commit = %s, %s", commitHashStr, err)

	}

	gitFile, err := rootTree.File(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get file = %s in commit = %s, %s", filePath, commitHashStr, err)
	}

	return gitFile, nil
}

func stateFileFromCommit(repoUrl, commitHashStr, filePath string) (*state.File, error) {
	gitFile, err := gitFileFromCommit(repoUrl, commitHashStr, filePath)
	if err != nil {
		return nil, fmt.Errorf("failed in stateFileFromCommit, %s", err)
	}

	fileState, err := state.NewFile(nil, gitFile)
	if err != nil {
		return nil, fmt.Errorf("failed in stateFileFromCommit, cannot create file state for file = %s in commit = %s, %s", filePath, commitHashStr, err)
	}

	return fileState, nil
}

func TestFilePatterns(t *testing.T) {
	cases := []struct {
		prevCommit         string
		currentCommit      string
		prevFilePath       string
		currentFilePath    string
		goldenFileOpenFile string
		goldenFileFileNode string
	}{
		{
			"55c98498a85f4503e3922586ceeb86ab5100e91f", //cleanup
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"next/package.json",
			"next/package.json",
			"testdata/file_add_openfile_golden.json",
			"testdata/file_add_filenode_golden.json",
		},
		{
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"fa2e1e5edb4379ceaaa9b9250e11c06c1fdbf4ad", //npm install --save @emotion/react
			"next/package.json",
			"next/package.json",
			"testdata/file_update_openfile_golden.json",
			"testdata/file_update_filenode_golden.json",
		},
		{
			"3b452151c8a567e2d42a133c255e85d81ea5912a",
			"55c98498a85f4503e3922586ceeb86ab5100e91f", //cleanup
			"pages/posts/[id].js",
			"pages/posts/[id].js",
			"testdata/file_delete_openfile_golden.json",
			"testdata/file_delete_filenode_golden.json",
		},
		{
			"e4a7c6509d5ff90da22612a66128eb38d418cb3e",
			"e5784f193c5e62a840bbfb96a2b9a5662d1361c1", //next to nextjs
			"next/pages/index.tsx",
			"nextjs/pages/index.tsx",
			"testdata/file_rename_openfile_golden.json",
			"testdata/file_rename_filenode_golden.json",
		},
	}

	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	for _, c := range cases {
		// gitFile = nil is ok, so ignore errors
		prevGitFile, _ := gitFileFromCommit(repoUrl, c.prevCommit, c.prevFilePath)
		currentGitFile, _ := gitFileFromCommit(repoUrl, c.currentCommit, c.currentFilePath)

		s, err := state.NewFile(prevGitFile, currentGitFile)
		if err != nil {
			t.Fatalf("failed in TestFilePatterns to create state.File, %s", err)
		}

		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileOpenFile, s.ToGraphQLOpenFile())
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileFileNode, s.ToGraphQLFileNode())
	}
}

// since state.File is effectively immutable, no need to test the state mutation, but only the GraphQL model mutation
func TestFileMutation1(t *testing.T) {
	s, err := stateFileFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestFileMutation1, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLOpenFile()
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
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLOpenFile())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/file_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}

// since state.File is effectively immutable, no need to test the state mutation, but only the GraphQL model mutation
func TestFileMutation2(t *testing.T) {
	s, err := stateFileFromCommit(
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
