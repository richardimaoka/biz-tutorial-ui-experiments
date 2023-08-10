package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func gitCommit(repoUrl, commitHashStr string) (*object.Commit, error) {
	repo, err := gitOpenOrClone(repoUrl)
	if err != nil {
		return nil, fmt.Errorf("failed in gitCommit for, %s", err)
	}

	commitHash := plumbing.NewHash(commitHashStr)
	if commitHash.String() != commitHashStr {
		return nil, fmt.Errorf("failed in gitCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in gitCommit, cannot get commit = %s, %s", commitHashStr, err)
	}

	return commit, nil
}

func gitFileFromCommit(repoUrl, commitHashStr, filePath string) (*object.File, error) {
	commit, err := gitCommit(repoUrl, commitHashStr)
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

func fileStateFromCommit(repoUrl, commitHashStr, filePath string) (*state.File, error) {
	gitFile, err := gitFileFromCommit(repoUrl, commitHashStr, filePath)
	if err != nil {
		return nil, fmt.Errorf("failed in stateFileFromCommit, %s", err)
	}

	fileState, err := state.FileUnChanged(gitFile, "")
	if err != nil {
		return nil, fmt.Errorf("failed in stateFileFromCommit, cannot create file state for file = %s in commit = %s, %s", filePath, commitHashStr, err)
	}

	return fileState, nil
}

func TestFileUnchanged(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"

	cases := []struct {
		commit             string
		filePath           string
		goldenFileOpenFile string
		goldenFileFileNode string
	}{
		{
			// text file
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"next/package.json",
			"testdata/file_unchanged_openfile_golden1.json",
			"testdata/file_unchanged_filenode_golden1.json",
		},
		{
			// another text file to see difference from case 1
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"next/pages/api/hello.ts",
			"testdata/file_unchanged_openfile_golden2.json",
			"testdata/file_unchanged_filenode_golden2.json",
		},
		{
			// binary file
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"next/public/favicon.ico",
			"testdata/file_unchanged_openfile_golden3.json",
			"testdata/file_unchanged_filenode_golden3.json",
		},
		{
			// svg file
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"next/public/next.svg",
			"testdata/file_unchanged_openfile_golden4.json",
			"testdata/file_unchanged_filenode_goldenr.json",
		},
	}

	for _, c := range cases {
		gitFile, err := gitFileFromCommit(repoUrl, c.commit, c.filePath)
		if err != nil {
			t.Fatalf("failed in TestFileUnchanged to get gitFile, %s", err)
		}

		s, err := state.FileUnChanged(gitFile, "") //curretDir = "", as gitFile is retrieved with respect to the rootDir
		if err != nil {
			t.Fatalf("failed in TestFileUnchanged to create state.File, %s", err)
		}

		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileOpenFile, s.ToGraphQLOpenFile())
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileFileNode, s.ToGraphQLFileNode())
	}
}

func TestFileAdded(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"

	cases := []struct {
		// prevCommit         string
		currentCommit      string
		filePath           string
		goldenFileOpenFile string
		goldenFileFileNode string
	}{
		{
			// "55c98498a85f4503e3922586ceeb86ab5100e91f", //cleanup
			"8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"next/package.json",
			"testdata/file_added_openfile_golden1.json",
			"testdata/file_added_filenode_golden1.json",
		},
	}

	for _, c := range cases {
		gitFile, err := gitFileFromCommit(repoUrl, c.currentCommit, c.filePath)
		if err != nil {
			t.Fatalf("failed in TestFileAdded to get gitFile, %s", err)
		}

		u, err := state.FileUnChanged(gitFile, "") //curretDir = "", as gitFile is retrieved with respect to the rootDir
		if err != nil {
			t.Fatalf("failed in TestFileAdded to create state.File, %s", err)
		}

		s := u.ToFileAdded()

		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileOpenFile, s.ToGraphQLOpenFile())
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileFileNode, s.ToGraphQLFileNode())
	}
}

func TestFileDeleted(t *testing.T) {
	// repoUrl := "https://github.com/richardimaoka/next-sandbox.git"

	cases := []struct {
		// prevCommit         string
		// currentCommit      string
		filePath           string
		goldenFileOpenFile string
		goldenFileFileNode string
	}{
		{
			// text file
			// "3b452151c8a567e2d42a133c255e85d81ea5912a", //getStaticProps
			// "55c98498a85f4503e3922586ceeb86ab5100e91f", //cleanup
			".gitignore",
			"testdata/file_deleted_openfile_golden1.json",
			"testdata/file_deleted_filenode_golden1.json",
		},
		{
			// "3b452151c8a567e2d42a133c255e85d81ea5912a", //getStaticProps
			// "55c98498a85f4503e3922586ceeb86ab5100e91f", //cleanup
			"pages/posts/[id].js",
			"testdata/file_deleted_openfile_golden2.json",
			"testdata/file_deleted_filenode_golden2.json",
		},
	}

	for _, c := range cases {
		s := state.FileDeleted(c.filePath)
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileOpenFile, s.ToGraphQLOpenFile())
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileFileNode, s.ToGraphQLFileNode())
	}
}

func TestFileUpdatd(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"

	cases := []struct {
		// prevCommit         string
		currentCommit      string
		filePath           string
		goldenFileOpenFile string
		goldenFileFileNode string
	}{
		{
			// TODO: calculate highlights
			// "8adac375628219e020d4b5957ff24f45954cbd3f", //npx create-next-app@latest
			"fa2e1e5edb4379ceaaa9b9250e11c06c1fdbf4ad", //npm install --save @emotion/react
			"next/package.json",
			"testdata/file_updated_openfile_golden1.json",
			"testdata/file_updated_filenode_golden1.json",
		},
		// //rename
		// "e4a7c6509d5ff90da22612a66128eb38d418cb3e",
		// "e5784f193c5e62a840bbfb96a2b9a5662d1361c1", //next to nextjs
		// "next/pages/index.tsx",                     //renamed to "nextjs/pages/index.tsx",
		// "testdata/file_renamed_openfile_golden.json",
		// "testdata/file_renamed_filenode_golden.json",
	}

	for _, c := range cases {
		currentFile, err := gitFileFromCommit(repoUrl, c.currentCommit, c.filePath)
		if err != nil {
			t.Fatalf("failed in TestFileUpdatd to get currentFile, %s", err)
		}

		u, err := state.FileUnChanged(currentFile, "") //curretDir = "", as gitFile is retrieved with respect to the rootDir
		if err != nil {
			t.Fatalf("failed in TestFileUpdatd to create state.File, %s", err)
		}

		s := u.ToFileUpdated()

		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileOpenFile, s.ToGraphQLOpenFile())
		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFileFileNode, s.ToGraphQLFileNode())
	}
}
