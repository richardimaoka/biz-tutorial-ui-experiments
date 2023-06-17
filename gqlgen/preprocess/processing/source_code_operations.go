package processing

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type SourceCodeOperation interface {
	IsSourceCodeOperation()
}

type SourceCodeFileOperation struct {
	FileOps []FileOperation
}

type SourceCodeGitOperation struct {
	CommitHash string
}

func (o SourceCodeFileOperation) IsSourceCodeOperation() {}
func (o SourceCodeGitOperation) IsSourceCodeOperation()  {}

func FileOpsFromCommit(repo *git.Repository, commit *object.Commit) []FileOperation {
	return nil
}
