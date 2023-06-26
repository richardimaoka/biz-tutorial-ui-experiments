package preprocess2

import "github.com/go-git/go-git/v5/plumbing"

type Directory struct {
	dirs     []Directory
	files    []File
	treeHash plumbing.Hash //assuming every transition is a git commit
}
