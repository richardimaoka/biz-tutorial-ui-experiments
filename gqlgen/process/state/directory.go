package state

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Directory struct {
	repo      *git.Repository
	dirPath   string
	offset    int
	dirName   string
	isUpdated bool
	isAdded   bool
	isDeleted bool
	files     Files
	subDirs   Directories
}

type Directories []*Directory

func (this *Directory) FilePath() string {
	return this.dirPath
}

func TreeFilesDirs(tree *object.Tree) ([]object.TreeEntry, []object.TreeEntry) {
	var files []object.TreeEntry
	var dirs []object.TreeEntry

	for _, e := range tree.Entries {
		if e.Mode.IsFile() {
			files = append(files, e)
		} else {
			dirs = append(dirs, e)
		}
	}
	return files, dirs
}

func SortEntries(entries []object.TreeEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name) < strings.ToLower(entries[j].Name)
	})
}

func NewDirectory(repo *git.Repository, dirPath string) *Directory {
	split := strings.Split(dirPath, "/")
	dirName := split[len(split)-1]
	offset := len(split) - 1

	dir := Directory{
		repo:    repo,
		dirPath: dirPath,
		dirName: dirName,
		offset:  offset,
	}

	return &dir
}

func (s *Directory) Recurse(dirPath string, tree *object.Tree) error {
	if tree == nil {
		return fmt.Errorf("failed in recurse, tree is nil")
	}

	fileEntries, subDirEntries := TreeFilesDirs(tree)
	SortEntries(fileEntries)
	SortEntries(subDirEntries)

	for _, d := range subDirEntries {
		subDirPath := FilePathInDir(dirPath, d.Name)
		subTree, err := object.GetTree(s.repo.Storer, d.Hash)
		if err != nil {
			return fmt.Errorf("failed in recurse, cannot get subtree = %s, %s", subDirPath, err)
		}

		subDir := NewDirectory(s.repo, subDirPath)
		if err := subDir.Recurse(subDirPath, subTree); err != nil {
			return fmt.Errorf("failed in recurse, cannot create directory = %s, %s", subDirPath, err)
		}
		s.subDirs = append(s.subDirs, subDir)
	}

	for _, f := range fileEntries {
		fileObj, err := tree.File(f.Name)
		if err != nil {
			return fmt.Errorf("failed in recurse, cannot get file = %s in dir = %s, %s", f.Name, dirPath, err)
		}

		file, err := FileUnChanged(fileObj, dirPath)
		if err != nil {
			return fmt.Errorf("failed in recurse, cannot create file = %s in dir = %s, %s", f.Name, dirPath, err)
		}

		s.files = append(s.files, file)
	}

	return nil
}

func (s *Directory) InsertFileDeleted(dirPath, relativeFilePath string, deletedFile diff.File) error {
	split := strings.Split(relativeFilePath, "/")
	if len(split) == 1 {
		file, err := FileDeleted(deletedFile)
		if err != nil {
			return fmt.Errorf("failed in InsertFileDeleted, cannot delete = %s, %s", deletedFile.Path(), err)
		}

		s.files = append(s.files, file)
		s.files.Sort()
		return nil
	} else {
		subDirName := split[0]
		for _, subdir := range s.subDirs {
			if subdir.dirName == subDirName {
				var newDirPath string
				if dirPath == "" {
					newDirPath = subDirName
				} else {
					newDirPath = dirPath + "/" + subDirName
				}
				newFilePath := strings.Join(split[1:], "/")
				return subdir.InsertFileDeleted(newDirPath, newFilePath, deletedFile)
			}
		}

		// if no matching subdir found
		var subDirPath string
		if dirPath == "" {
			subDirPath = subDirName
		} else {
			subDirPath = dirPath + "/" + subDirName
		}
		subDir := NewDirectory(s.repo, subDirPath)
		relativeFilePath := strings.Join(split[1:], "/")
		if err := subDir.InsertFileDeleted(subDirPath, relativeFilePath, deletedFile); err != nil {
			return fmt.Errorf("failed in InsertFileDeleted, cannot mark deletion file = %s, %s", deletedFile.Path(), err)
		}
		s.subDirs = append(s.subDirs, subDir)
		s.subDirs.Sort()
		return nil
	}
}

func (s *Directory) MarkFileUpdated(filePath string, previFile *object.File) error {
	return nil
}

func (s *Directory) MarkFileRenamed(filePath string, previFile *object.File) error {
	return nil
}

func (s *Directory) MarkFileAdded(filePath string) error {
	split := strings.Split(filePath, "/")
	if len(split) == 1 {
		for i, f := range s.files {
			if f.fileName == filePath {
				added, err := f.ToFileAdded()
				if err != nil {
					return fmt.Errorf("failed in MarkFileAdded, cannot mark file = %s as added, %s", filePath, err)
				}
				s.files[i] = added
				return nil
			}
		}
	} else {
		subDirName := split[0]
		for _, subdir := range s.subDirs {
			if subdir.dirName == subDirName {
				newFilePath := strings.Join(split[1:], "/")
				return subdir.MarkFileAdded(newFilePath)
			}
		}
	}

	return fmt.Errorf("failed in MarkFileAdded, cannot find file = %s", filePath)
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

func (s *Directory) ToGraphQLFileNodeSlice() []*model.FileNode {
	var fileNodes []*model.FileNode

	if s.dirPath != "" {
		fileNodes = append(fileNodes, s.ToGraphQLFileNode())
	}

	for _, subdir := range s.subDirs {
		subFileNodes := subdir.ToGraphQLFileNodeSlice()
		fileNodes = append(fileNodes, subFileNodes...)
	}

	for _, file := range s.files {
		fileNodes = append(fileNodes, file.ToGraphQLFileNode())
	}

	return fileNodes
}

func (dirs Directories) Sort() {
	sort.Slice(dirs, func(i, j int) bool {
		return strings.ToLower(dirs[i].dirName) < strings.ToLower(dirs[j].dirName)
	})
}
