package rough

import (
	"fmt"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func gitFilesForCommit(repo *git.Repository, commitHashStr string) ([]string, error) {
	commitHash := plumbing.NewHash(commitHashStr)
	if commitHash.String() != commitHashStr {
		return nil, fmt.Errorf("failed in gitFilesForCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in gitFilesForCommit, cannot get commit = %s, %s", commitHashStr, err)
	}

	fileIter, err := commit.Files()
	if err != nil {
		return nil, fmt.Errorf("failed in gitFilesForCommit, cannot get file iterator for commit = %s, %s", commitHashStr, err)
	}

	var files []string
	for {
		file, err := fileIter.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error in commit file traversal in commit %s, %v", commitHash, err)
		}
		files = append(files, file.Name)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("failed to get files for commit = %s, no files found", commitHashStr)
	}

	return files, nil
}
