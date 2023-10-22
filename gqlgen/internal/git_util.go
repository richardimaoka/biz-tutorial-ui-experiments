package internal

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func validateCommitHash(hashStr string) (plumbing.Hash, error) {
	commitHash := plumbing.NewHash(hashStr)
	if commitHash.String() != hashStr {
		return plumbing.ZeroHash, fmt.Errorf("commit hash = %s mismatched with re-calculated hash = %s", hashStr, commitHash.String())
	}

	return commitHash, nil
}

func errroMessage(prefix, leadingMessage string, underlyingError error) error {
	return fmt.Errorf("%s - %s, %s", prefix, leadingMessage, underlyingError)
}

func GetCommit(repo *git.Repository, hashStr string) (*object.Commit, error) {
	funcName := "internal.GetCommit"

	commitHash, err := validateCommitHash(hashStr)
	if err != nil {
		return nil, errroMessage(funcName, "validation error", err)
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get commit for %s", hashStr), err)
	}

	return commit, nil
}
