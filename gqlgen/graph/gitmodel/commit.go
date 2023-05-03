package gitmodel

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

type File struct {
	path       string
	commitHash plumbing.Hash
	Size       int64
}

func NewFile(repoUrl string, blobHash plumbing.Hash) *File {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repoUrl,
	})
	if err != nil {
		log.Println("error cloning repo")
		return nil
	}

	blob, err := r.BlobObject(blobHash)
	if err != nil {
		log.Println("error getting blog")
		return nil
	}

	log.Println("size", blob.Size)

	return &File{Size: blob.Size}
}

func (f *File) Contents() *string {
	return nil
}
