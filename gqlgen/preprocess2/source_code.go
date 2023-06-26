package preprocess2

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type SourceCode struct {
	currentCommit plumbing.Hash
	RootDir       Directory
}

func (s *SourceCode) Transition(commit object.Commit) string {
	// if commit has two parents, error
	// if commit' only parent is NOT current commit, error
	return ""
}
