package gitmodel

import (
	"io"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type FileFromGit struct {
	filePath string
	blob     *object.Blob
	commit   *object.Commit
}

func NewFileFromGit(repo *git.Repository, filePath string, blobHash, commitHash plumbing.Hash) (*FileFromGit, error) {
	blob, err := repo.BlobObject(blobHash)
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, err
	}

	return &FileFromGit{
		blob:     blob,
		commit:   commit,
		filePath: filePath}, nil
}

func (f *FileFromGit) offset() int {
	split := strings.Split(f.filePath, "/")
	return len(split) - 1
}

func (f *FileFromGit) name() string {
	split := strings.Split(f.filePath, "/")
	return split[len(split)-1]
}

func (f *FileFromGit) Contents() *string {
	reader, err := f.blob.Reader()
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
	name := f.name()
	nodeType := model.FileNodeTypeFile
	offset := f.offset()

	return &model.FileNode{
		FilePath: &f.filePath, //pointer is safe here, as f.filePath is effectively immutable
		Name:     &name,
		NodeType: &nodeType,
		Offset:   &offset,
	}
}

func (f *FileFromGit) OpenFile() *model.OpenFile {
	return nil
}
