package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

var storage *memory.Storage
var repoCache map[string]*git.Repository

func gitOpenOrClone(repoUrl string) (*git.Repository, error) {
	if repo, ok := repoCache[repoUrl]; ok {
		return repo, nil
	}

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("cannot clone repo %s, %s", repoUrl, err)
	}

	repoCache[repoUrl] = repo
	return repo, nil
}

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

func gitTreeForCommit(repoUrl, commitHash, dirPath string) (*object.Tree, error) {
	commit, err := gitCommit(repoUrl, commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get commit = %s, %s", commitHash, err)
	}

	rootTree, err := commit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get tree for commit = %s, %s", commitHash, err)

	}

	gitTree, err := rootTree.Tree(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get tree = %s in commit = %s, %s", dirPath, commitHash, err)
	}

	return gitTree, nil
}

func TestMain(m *testing.M) {
	repoCache = make(map[string]*git.Repository)

	fmt.Println("before all tests")
	m.Run()
}
