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

func NewDirectory(repo *git.Repository, dirPath string, currentTree *object.Tree, doRecurse bool) (*Directory, error) {
	split := strings.Split(dirPath, "/")
	dirName := split[len(split)-1]
	offset := len(split) - 1

	var subDirs []*Directory
	var files []*File

	if doRecurse {
		if currentTree == nil {
			return nil, fmt.Errorf("failed in NewDirectory, currentTree is nil")
		}

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

			file, err := FileUnChanged(fileObj, dirPath)
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
		var newDirPath string
		if dirPath == "" {
			newDirPath = subDirName
		} else {
			newDirPath = dirPath + "/" + subDirName
		}
		subdir, err := NewDirectory(nil, newDirPath, nil, false)
		if err != nil {
			return fmt.Errorf("failed in InsertFileDeleted, cannot create subdir = %s under = %s, %s", subDirName, dirPath, err)
		}
		newFilePath := strings.Join(split[1:], "/")
		if err := subdir.InsertFileDeleted(newDirPath, newFilePath, deletedFile); err != nil {
			return fmt.Errorf("failed in InsertFileDeleted, cannot mark deletion file = %s, %s", deletedFile.Path(), err)
		}
		s.subDirs = append(s.subDirs, subdir)
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
