package state

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Directory struct {
	dirPath   string
	offset    int
	dirName   string
	isUpdated bool
	isAdded   bool
	isDeleted bool
}

func (d *Directory) FilePath() string {
	return d.dirPath
}

func NewDirectory(dirPath string, prevTree *object.Tree, currentTree *object.Tree) (*Directory, error) {
	if currentTree == nil && prevTree == nil {
		return nil, fmt.Errorf("failed in NewDirectory, currentTree and prevTree are both nil")
	}

	split := strings.Split(dirPath, "/")
	dirName := split[len(split)-1]

	offset := len(split) - 1

	return &Directory{
		dirPath: dirPath,
		dirName: dirName,
		offset:  offset,
	}, nil
}

func (s *Directory) ToGraphQLFileNode() *model.FileNode {
	//copy to avoid mutation effects afterwards
	dirType := model.FileNodeTypeDirectory
	dirPath := s.dirPath
	name := s.dirName
	offset := s.offset
	falseValue := false

	return &model.FileNode{
		NodeType:  &dirType,
		FilePath:  &dirPath,
		Name:      &name,
		Offset:    &offset,
		IsUpdated: &falseValue, //git doesn't track standalone dir, so changes are always in contained files
	}
}
