package gitmodel

import (
	"io"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type FileFromGit struct {
	path           string
	commitHash     plumbing.Hash
	blobHash       plumbing.Hash
	repo           *git.Repository
	Size           int64
	_Contents      string
	IsFullContents bool
}

func NewFile(repo *git.Repository, blobHash plumbing.Hash) *FileFromGit {
	blob, err := repo.BlobObject(blobHash)
	if err != nil {
		log.Println("error getting blog")
		return nil
	}

	log.Println("size", blob.Size)

	return &FileFromGit{Size: blob.Size, blobHash: blobHash, repo: repo}
}

func (f *FileFromGit) Contents() *string {
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
func (f *FileFromGit) IsUpdated() bool {
	return false
}

func (f *FileFromGit) NodeType() string {
	return ""
}

func (f *FileFromGit) FilePath() string {
	return ""
}

func (f *FileFromGit) FileNode() *model.FileNode {
	return nil
}

func (f *FileFromGit) OpenFile() *model.OpenFile {
	return nil
}
