package state

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type File struct {
	Repo        *git.Repository
	PrevFile    *object.File
	CurrentFile *object.File
}

func NewFile(repo *git.Repository, prevFile *object.File, currentFile *object.File) (*File, error) {
	if repo == nil {
		return nil, fmt.Errorf("failed in NewFile, repo is nil")
	}
	if currentFile == nil {
		return nil, fmt.Errorf("failed in NewFile, currentFile is nil")
	}

	return &File{
		Repo:        repo,
		PrevFile:    prevFile,
		CurrentFile: currentFile,
	}, nil
}

func (s *File) ToGraphQLOpenFileNode() *model.FileNode {
	//copy to avoid mutation effects afterwards
	fileType := model.FileNodeTypeFile
	filePath := s.CurrentFile.Name
	split := strings.Split(filePath, "/")
	fileName := split[len(split)-1]
	offset := len(split) - 1
	isUpdated := s.PrevFile == nil || s.PrevFile.Hash != s.CurrentFile.Hash

	return &model.FileNode{
		NodeType:  &fileType,
		FilePath:  &filePath,
		Name:      &fileName,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}
