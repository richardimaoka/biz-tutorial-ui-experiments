package state

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type File struct {
	repo        *git.Repository
	prevFile    *object.File
	currentFile *object.File
}

func NewFile(repo *git.Repository, prevFile *object.File, currentFile *object.File) (*File, error) {
	if repo == nil {
		return nil, fmt.Errorf("failed in NewFile, repo is nil")
	}
	if currentFile == nil && prevFile == nil {
		return nil, fmt.Errorf("failed in NewFile, currentFile and prevFile are both nil")
	}

	return &File{
		repo:        repo,
		prevFile:    prevFile,
		currentFile: currentFile,
	}, nil
}

func (s *File) ToGraphQLOpenFile() (*model.OpenFile, error) {
	//copy to avoid mutation effects afterwards
	filePath := s.currentFile.Name
	split := strings.Split(filePath, "/")
	fileName := split[len(split)-1]
	contents, err := s.currentFile.Contents()
	if err != nil {
		return nil, fmt.Errorf("failed in ToGraphQLFileNode for file = %s, cannot get file contents, %s", filePath, err)
	}
	trueValue := true

	return &model.OpenFile{
		FilePath:      &filePath,
		FileName:      &fileName,
		IsFullContent: &trueValue,
		Content:       &contents,
	}, nil
}

func (s *File) ToGraphQLFileNode() *model.FileNode {
	//copy to avoid mutation effects afterwards
	fileType := model.FileNodeTypeFile
	filePath := s.currentFile.Name
	split := strings.Split(filePath, "/")
	fileName := split[len(split)-1]
	offset := len(split) - 1
	isUpdated := s.prevFile == nil || s.prevFile.Hash != s.currentFile.Hash

	return &model.FileNode{
		NodeType:  &fileType,
		FilePath:  &filePath,
		Name:      &fileName,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}
