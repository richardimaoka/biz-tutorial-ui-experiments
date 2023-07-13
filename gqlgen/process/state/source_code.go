package state

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type SourceCode struct {
	repo   *git.Repository
	commit plumbing.Hash
}
