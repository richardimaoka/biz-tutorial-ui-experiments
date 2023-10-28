package gitwrap

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func errroMessage(prefix, leadingMessage string, underlyingError error) error {
	return fmt.Errorf("%s - %s, %s", prefix, leadingMessage, underlyingError)
}

func ValidateCommitHash(hashStr string) (plumbing.Hash, error) {
	commitHash := plumbing.NewHash(hashStr)
	if commitHash.String() != hashStr {
		return plumbing.ZeroHash, fmt.Errorf("commit hash = %s mismatched with re-calculated hash = %s", hashStr, commitHash.String())
	}

	return commitHash, nil
}

// Get git commit object from hash string
// bit easier than go-get's equivalent, as this function works with string, not plumbing.Hash
func GetCommit(repo *git.Repository, hashStr string) (*object.Commit, error) {
	funcName := "gitwrap.GetCommit"

	commitHash, err := ValidateCommitHash(hashStr)
	if err != nil {
		return nil, errroMessage(funcName, "validation error", err)
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get commit for %s", hashStr), err)
	}

	return commit, nil
}

// Get git patch object from hash strings
// bit easier than go-get's equivalent, as this function works with string, not plumbing.Hash
func GetPatch(repo *git.Repository, fromCommitHash, toCommitHash string) (*object.Patch, error) {
	funcName := "gitwrap.GetPatch"

	fromCommit, err := GetCommit(repo, fromCommitHash)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get commit for %s", fromCommitHash), err)
	}

	toCommit, err := GetCommit(repo, toCommitHash)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get commit for %s", toCommitHash), err)
	}

	patch, err := fromCommit.Patch(toCommit)
	if err != nil {
		return nil, errroMessage(funcName, fmt.Sprintf("cannot get patch from = %s to = %s", fromCommitHash, toCommitHash), err)
	}

	return patch, nil
}

func FindFilePatch(patch *object.Patch, fileFullPath string) (string, diff.FilePatch) {
	for _, filePatch := range patch.FilePatches() {
		// Files returns the from and to Files, with all the necessary metadata
		// about them. If the patch creates a new file, "from" will be nil.
		// If the patch deletes a file, "to" will be nil.
		from, to := filePatch.Files()

		if from == nil /* (i.e.) to != nil */ {
			if to.Path() == fileFullPath {
				return "Add", filePatch
			}
		} else if to == nil /* (i.e.) from != nil */ {
			if from.Path() == fileFullPath {
				return "Add", filePatch
			}
		} else if from != nil && to != nil {
			if from.Path() == fileFullPath {
				if to.Path() == fileFullPath {
					return "Update", filePatch
				} else {
					return "Rename-From", filePatch
				}
			} else if to.Path() == fileFullPath {
				if from.Path() == fileFullPath {
					return "Update", filePatch
				} else {
					return "Rename-To", filePatch
				}
			}
		}
	}

	return "", nil
}
