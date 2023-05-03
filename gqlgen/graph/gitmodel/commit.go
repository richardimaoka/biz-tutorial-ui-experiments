package gitmodel

import (
	"github.com/go-git/go-git/v5/plumbing"
)

type file struct {
	path string
	hash plumbing.Hash
}

func (f *file) Contents() string {
	return f.path
}
