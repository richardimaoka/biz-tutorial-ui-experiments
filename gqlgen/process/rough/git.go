package rough

import (
	"fmt"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
)

func gitFilesForCommit(repo *git.Repository, commitHashStr string) ([]string, error) {
	commit, err := gitwrap.GetCommit(repo, commitHashStr)

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
			return nil, fmt.Errorf("error in commit file traversal in commit %s, %v", commitHashStr, err)
		}
		files = append(files, file.Name)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("failed to get files for commit = %s, no files found", commitHashStr)
	}

	return files, nil
}

func filesForDiffInternal(repo *git.Repository, currentCommit, prevCommit *object.Commit) ([]string, error) {
	patch, err := prevCommit.Patch(currentCommit)
	if err != nil {
		return nil, fmt.Errorf("failed in filesForDiff, cannot get patch between commit = %s and commit = %s, %s", prevCommit.Hash.String(), currentCommit.Hash.String(), err)
	}

	var files []string
	for _, filePatch := range patch.FilePatches() {
		from, to := filePatch.Files()
		if from == nil {
			//added
			files = append(files, to.Path())
		} else if to == nil {
			// deleted
			// files = append(files, from.Path())
			panic("file deletion is not implemented")
		} else if from.Path() != to.Path() {
			// renamed
			files = append(files, to.Path())
		} else {
			// updated
			files = append(files, to.Path())
		}
	}

	return files, nil
}

func filesForDiff(repo *git.Repository, currentCommitHash, prevCommitHash string) ([]string, error) {
	currentCommit, err := gitwrap.GetCommit(repo, currentCommitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in GitFilesForDiff, %s", err)
	}

	prevCommit, err := gitwrap.GetCommit(repo, prevCommitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in GitFilesForDiff, %s", err)
	}

	return filesForDiffInternal(repo, currentCommit, prevCommit)
}

func CommitFiles(repo *git.Repository, currentCommitHash, prevCommitHash string) ([]string, error) {
	if prevCommitHash == "" {
		return gitFilesForCommit(repo, currentCommitHash)
	} else {
		return filesForDiff(repo, currentCommitHash, prevCommitHash)
	}
}
