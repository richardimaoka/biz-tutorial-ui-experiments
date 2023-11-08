package gitwrap

import (
	"fmt"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func errorMessage(prefix, leadingMessage string, underlyingError error) error {
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
		return nil, errorMessage(funcName, "validation error", err)
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, errorMessage(funcName, fmt.Sprintf("cannot get commit for %s", hashStr), err)
	}

	return commit, nil
}

// Name of the returning object.File should be full path, as it is retrieved from the
// root tree of the commit.
func GetCommitFiles(repo *git.Repository, commitHashStr string) ([]object.File, error) {
	funcName := "gitwrap.GetCommitFiles"
	commit, err := GetCommit(repo, commitHashStr)
	if err != nil {
		return nil, errorMessage(funcName, "failed", err)
	}

	fileIter, err := commit.Files()
	if err != nil {
		return nil, errorMessage(funcName, "failed to get files for commit = "+commitHashStr, err)
	}

	var files []object.File
	for {
		file, err := fileIter.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errorMessage(funcName, "failed to traverse files in commit = "+commitHashStr, err)
		}
		files = append(files, *file)
	}

	return files, nil
}

// Get git patch object from hash strings
// bit easier than go-get's equivalent, as this function works with string, not plumbing.Hash
func GetPatch(repo *git.Repository, fromCommitHash, toCommitHash string) (*object.Patch, error) {
	funcName := "gitwrap.GetPatch"

	fromCommit, err := GetCommit(repo, fromCommitHash)
	if err != nil {
		return nil, errorMessage(funcName, fmt.Sprintf("cannot get commit for %s", fromCommitHash), err)
	}

	toCommit, err := GetCommit(repo, toCommitHash)
	if err != nil {
		return nil, errorMessage(funcName, fmt.Sprintf("cannot get commit for %s", toCommitHash), err)
	}

	patch, err := fromCommit.Patch(toCommit)
	if err != nil {
		return nil, errorMessage(funcName, fmt.Sprintf("cannot get patch from = %s to = %s", fromCommitHash, toCommitHash), err)
	}

	return patch, nil
}

// Returns []diff.File, not []object.File
func GetPatchFiles(repo *git.Repository, fromCommitHash, toCommitHash string) ([]diff.File, error) {
	funcName := "gitwrap.GetPatchFiles"

	patch, err := GetPatch(repo, fromCommitHash, toCommitHash)
	if err != nil {
		return nil, errorMessage(funcName, "failed", err)
	}

	var files []diff.File
	for _, filePatch := range patch.FilePatches() {
		from, to := filePatch.Files()
		if from == nil {
			//added
			files = append(files, to)
		} else if to == nil {
			// deleted
			// files = append(files, from.Path())
			return nil, errorMessage(funcName, "file deletion is not implemented", err)
		} else if from.Path() != to.Path() {
			// renamed
			files = append(files, to)
		} else {
			// updated
			files = append(files, to)
		}
	}

	return files, nil
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
