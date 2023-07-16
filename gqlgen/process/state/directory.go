package state

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
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
	files     []*File
	subDirs   []*Directory
}

func (this *Directory) FilePath() string {
	return this.dirPath
}

func NewDirectory(repo *git.Repository, dirPath string, currentTree *object.Tree, doRecurse bool) (*Directory, error) {
	if currentTree == nil {
		return nil, fmt.Errorf("failed in NewDirectory, currentTree and is nil")
	}

	split := strings.Split(dirPath, "/")
	dirName := split[len(split)-1]
	offset := len(split) - 1

	var subDirs []*Directory
	var files []*File

	if doRecurse {
		fileEntries, subDirEntries := TreeFilesDirs(currentTree)
		SortEntries(fileEntries)
		SortEntries(subDirEntries)

		for _, d := range subDirEntries {
			subDirPath := FilePathInDir(dirPath, d.Name)
			subTree, err := object.GetTree(repo.Storer, d.Hash)
			if err != nil {
				return nil, fmt.Errorf("failed in NewDirectory, cannot get tree = %s, %s", subDirPath, err)
			}

			subDir, err := NewDirectory(repo, subDirPath, subTree, true)
			if err != nil {
				return nil, fmt.Errorf("failed in recursive, cannot create directory = %s, %s", subDirPath, err)
			}
			subDirs = append(subDirs, subDir)
		}

		for _, f := range fileEntries {
			fileObj, err := currentTree.File(f.Name)
			if err != nil {
				return nil, fmt.Errorf("failed in NewDirectory, cannot get file = %s in dir = %s, %s", f.Name, dirPath, err)
			}

			file, err := NewFileUnChanged(fileObj, dirPath)
			if err != nil {
				return nil, fmt.Errorf("failed in NewDirectory, cannot create file = %s in dir = %s, %s", f.Name, dirPath, err)
			}

			files = append(files, file)
		}
	}

	return &Directory{
		dirPath: dirPath,
		dirName: dirName,
		offset:  offset,
		files:   files,
		subDirs: subDirs,
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
