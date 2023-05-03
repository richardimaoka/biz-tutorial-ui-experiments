package gitmodel

import (
	"io"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type File struct {
	path       string
	commitHash plumbing.Hash
	blobHash   plumbing.Hash
	repo       *git.Repository
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

	return &File{Size: blob.Size, blobHash: blobHash, repo: r}
}

func (f *File) Contents() *string {
	blob, err := f.repo.BlobObject(f.blobHash)
	if err != nil {
		return nil
	}

	reader, err := blob.Reader()
	if err != nil {
		return nil
	}

	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil
	}

	contents := string(bytes)
	return &contents
}

// File method to return model.FileNode
func (f *File) File() *model.FileNode {
	return nil
}
