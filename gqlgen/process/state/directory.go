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

func (dirs Directories) sortSelf() {
	sort.Slice(dirs, func(i, j int) bool {
		return strings.ToLower(dirs[i].dirName) < strings.ToLower(dirs[j].dirName)
	})
}

func filePathInDir(parentDir, name string) string {
	if parentDir != "" {
		return parentDir + "/" + name
	} else {
		return name
	}
}

func treeFilesDirs(tree *object.Tree) ([]object.TreeEntry, []object.TreeEntry) {
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

func sortEntries(entries []object.TreeEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name) < strings.ToLower(entries[j].Name)
	})
}

func EmptyDirectory(repo *git.Repository, dirPath string) *Directory {
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

func ConstructDirectory(repo *git.Repository, dirPath string, tree *object.Tree, filesByAdd bool) (*Directory, error) {
	dir := EmptyDirectory(repo, dirPath)
	if err := dir.recursivelyConstruct(dirPath, tree, filesByAdd); err != nil {
		return nil, fmt.Errorf("failed in ConstructDirectory for dirPath = %s, %s", dirPath, err)
	}

	return dir, nil
}

func (s *Directory) recursivelyConstruct(dirPath string, tree *object.Tree, filesByAdd bool) error {
	if tree == nil {
		return fmt.Errorf("failed in recurse, tree is nil")
	}

	// 1. Find subdirs and files directly belonging to dirPath
	fileEntries, subDirEntries := treeFilesDirs(tree)
	sortEntries(fileEntries)
	sortEntries(subDirEntries)

	// 2. Construct subdirs recursively
	for _, d := range subDirEntries {
		subDirPath := filePathInDir(dirPath, d.Name)
		subTree, err := object.GetTree(s.repo.Storer, d.Hash)
		if err != nil {
			return fmt.Errorf("failed in recurse, cannot get subtree = %s, %s", subDirPath, err)
		}

		// Recursive, and depth first construction
		subDir := EmptyDirectory(s.repo, subDirPath)
		if err := subDir.recursivelyConstruct(subDirPath, subTree, filesByAdd); err != nil {
			return fmt.Errorf("failed in recurse, cannot create directory = %s, %s", subDirPath, err)
		}
		s.subDirs = append(s.subDirs, subDir)
	}

	// 3. Construct files under dirPath
	for _, f := range fileEntries {
		fileObj, err := tree.File(f.Name)
		if err != nil {
			return fmt.Errorf("failed in recurse, cannot get git file = %s in dir = %s, %s", f.Name, dirPath, err)
		}

		var file *File
		if filesByAdd { // Upon construction, all files are considered added, assuming this is the first commit
			file, err = FileAdded(fileObj, dirPath)
		} else { //        Upon construction, all files are considered unchanged, then later marked as added/updated/deleted using git patch info
			file, err = FileUnChanged(fileObj, dirPath)
		}
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
		file := FileDeleted(deletedFile.Path())
		s.files = append(s.files, file)
		s.files.sortSelf()
		return nil
	} else {
		subDirName := split[0]
		for _, subdir := range s.subDirs {
			if subdir.dirName == subDirName {
				subDirPath := filePathInDir(dirPath, subDirName)
				strippedFilePath := strings.Join(split[1:], "/")
				return subdir.InsertFileDeleted(subDirPath, strippedFilePath, deletedFile)
			}
		}

		// if no matching subdir found
		subDirPath := filePathInDir(dirPath, subDirName)
		subDir := EmptyDirectory(s.repo, subDirPath)
		strippedFilePath := strings.Join(split[1:], "/")
		if err := subDir.InsertFileDeleted(subDirPath, strippedFilePath, deletedFile); err != nil {
			return fmt.Errorf("failed in InsertFileDeleted, cannot mark deletion file = %s, %s", deletedFile.Path(), err)
		}
		s.subDirs = append(s.subDirs, subDir)
		s.subDirs.sortSelf()
		return nil
	}
}

func (s *Directory) MarkFileUpdated(relativeFilePath string, fromFile diff.File) error {
	split := strings.Split(relativeFilePath, "/")
	if len(split) == 1 {
		for i, f := range s.files {
			if f.fileName == relativeFilePath {
				added := f.ToFileUpdated()
				s.files[i] = added
				return nil
			}
		}
	} else {
		subDirName := split[0]
		for _, subdir := range s.subDirs {
			if subdir.dirName == subDirName {
				strippedFilePath := strings.Join(split[1:], "/")
				return subdir.MarkFileUpdated(strippedFilePath, fromFile)
			}
		}
	}

	return fmt.Errorf("failed in MarkFileUpdated, cannot find file = %s", fromFile.Path())
}

func (s *Directory) MarkFileRenamed(filePath string, previFile *object.File) error {
	return nil
}

func (s *Directory) MarkFileAdded(relativeFilePath string) error {
	split := strings.Split(relativeFilePath, "/")
	if len(split) == 1 {
		for i, f := range s.files {
			if f.fileName == relativeFilePath {
				added, err := f.ToFileAdded()
				if err != nil {
					return fmt.Errorf("failed in MarkFileAdded, cannot mark file = %s as added, %s", relativeFilePath, err)
				}
				s.files[i] = added
				return nil
			}
		}
	} else {
		subDirName := split[0]
		for _, subdir := range s.subDirs {
			if subdir.dirName == subDirName {
				strippedFilePath := strings.Join(split[1:], "/")
				return subdir.MarkFileAdded(strippedFilePath)
			}
		}
	}

	return fmt.Errorf("failed in MarkFileAdded, cannot find file = %s", relativeFilePath)
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

func (s *Directory) ToGraphQLOpenFileMap() map[string]model.OpenFile {
	openFileMap := make(map[string]model.OpenFile)

	for _, subdir := range s.subDirs {
		subMap := subdir.ToGraphQLOpenFileMap()
		for k, v := range subMap {
			openFileMap[k] = v
		}
	}

	for _, file := range s.files {
		openFileMap[file.filePath] = *file.ToGraphQLOpenFile()
	}

	return openFileMap
}
