package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

type SourceCode struct {
	repo   *git.Repository
	commit plumbing.Hash
}

func NewSourceCode(repoUrl, currentCommitHash string) (*SourceCode, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot clone repo %s, %s", repoUrl, err)
	}

	commitHash := plumbing.NewHash(currentCommitHash)
	if commitHash.String() != currentCommitHash {
		return nil, fmt.Errorf("failed in gitFileFromCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", currentCommitHash, commitHash.String())
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get commit = %s, %s", currentCommitHash, err)
	}

	rootTree, err := commit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in gitFileFromCommit, cannot get tree for commit = %s, %s", currentCommitHash, err)
	}

	for _, e := range rootTree.Entries {
		fmt.Println(e.Name)
	}

	return nil, nil
}
