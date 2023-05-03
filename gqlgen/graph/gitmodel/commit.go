package gitmodel

import (
	"github.com/go-git/go-git/v5/plumbing"
)

type File struct {
	path string
	hash plumbing.Hash
}

func (f *File) Contents() *string {
	return nil
}
