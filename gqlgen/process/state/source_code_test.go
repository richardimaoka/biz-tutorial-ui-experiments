package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func gitTreeForCommit(repoUrl, commitHashStr, dirPath string) (*object.Tree, error) {
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

	gitTree, err := rootTree.Tree(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get tree = %s in commit = %s, %s", dirPath, commitHashStr, err)
	}

	return gitTree, nil
}

func TestTreeFilesDirs(t *testing.T) {
	tree, err := gitTreeForCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next",
	)
	if err != nil {
		t.Fatalf("failed in TestTreeDirectories to get git tree, %s", err)
	}

	files, dirs := state.TreeFilesDirs(tree)
	expectedFiles := []string{
		".eslintrc.json",
		".gitignore",
		"README.md",
		"next.config.js",
		"package-lock.json",
		"package.json",
		"tsconfig.json",
	}
	expectedDirs := []string{
		"pages",
		"public",
		"styles",
	}

	if len(files) != len(expectedFiles) {
		t.Fatalf("len(files) = %d mismatched with len(expectedFiles) = %d", len(files), len(expectedFiles))
	}

	for i, f := range files {
		if f.Name != expectedFiles[i] {
			t.Fatalf("files[%d] = %s mismatched with expectedFiles[%d] = %s", i, f.Name, i, expectedFiles[i])
		}
	}

	if len(dirs) != len(expectedDirs) {
		t.Fatalf("len(dirs) = %d mismatched with len(expectedDirs) = %d", len(dirs), len(expectedDirs))
	}

	for i, d := range dirs {
		if d.Name != expectedDirs[i] {
			t.Fatalf("dirs[%d] = %s mismatched with expectedDirs[%d] = %s", i, d.Name, i, expectedDirs[i])
		}
	}
}

func TestSourceCodePatterns(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	currentCommitHash := "8adac375628219e020d4b5957ff24f45954cbd3f"
	_, err := state.NewSourceCode(repoUrl, currentCommitHash)
	if err != nil {
		t.Fatalf("failed in TestSourceCodePatterns to create state.SourceCode, %s", err)
	}

	t.Fatalf("intentionally failed")
}
