package gitmodel

import (
	"io"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type FileFromGit struct {
	filePath       string
	commitHash     plumbing.Hash
	blobHash       plumbing.Hash
	repo           *git.Repository
	Size           int64
	_Contents      string
	IsFullContents bool
}

func NewFileFromGit(repo *git.Repository, filePath string, blobHash plumbing.Hash) (*FileFromGit, error) {
	blob, err := repo.BlobObject(blobHash)
	if err != nil {
		return nil, err
	}

	log.Println("size", blob.Size)

	return &FileFromGit{Size: blob.Size, blobHash: blobHash, repo: repo, filePath: filePath}, nil
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

func (f *FileFromGit) FileNode() *model.FileNode {
	return &model.FileNode{
		FilePath: &f.filePath, //pointer is safe here, as f.filePath is effectively immutable
	}
}

func (f *FileFromGit) OpenFile() *model.OpenFile {
	return nil
}
