package state

import (
	"fmt"
	"sort"
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

func sortEntries(entries []object.TreeEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name) < strings.ToLower(entries[j].Name)
	})
}

func emptyDirectory(fullPath string) *Directory {
	split := strings.Split(fullPath, "/")
	dirName := split[len(split)-1]
	offset := len(split) - 1

	dir := Directory{
		dirPath: fullPath,
		dirName: dirName,
		offset:  offset,
	}

	return &dir
}

func constructDirectory(files Files) (*Directory, error) {
	rootDir := emptyDirectory("")

	for _, f := range files {
		if err := rootDir.addFile(f.filePath, f); err != nil {
			return nil, fmt.Errorf("constructDirectory failed, %s", err)
		}
	}

	return rootDir, nil
}

func (s *Directory) addFile(relativeFilePath string, file *File) error {
	split := strings.Split(relativeFilePath, "/")
	isBareFile := len(split) == 1 // bare file (i.e.) this file is not in a directory like `main.tsx`, not `src/main.tsx`

	if isBareFile {
		// `relativeFilePath` is a bare file, not in a directory.
		// So try to find it from files in Directory (= receiver of this method)
		for _, f := range s.files {
			if f.fileName == relativeFilePath {
				return fmt.Errorf("addFile failed, in dir = '%s' file = '%s' already exists", s.dirPath, relativeFilePath)
			}
		}

		// if not found, add a file and return it
		s.files = append(s.files, file)
		s.files.sortSelf()
		return nil

	} else {
		// `relativeFilePath` is a file within a directory.
		// So recursively find a file in sub directories
		targetSubDir := split[0]
		// strip the current directory, and get a relative path to call this function recursively.
		nextRelativeFilePath := strings.Join(split[1:], "/")

		for _, subdir := range s.subDirs {
			if subdir.dirName == targetSubDir {
				return subdir.addFile(nextRelativeFilePath, file)
			}
		}

		// if not found, add a directory
		var targetSubDirFullPath string
		if s.dirPath == "" {
			targetSubDirFullPath = targetSubDir
		} else {
			targetSubDirFullPath = s.dirPath + "/" + targetSubDir
		}
		subDir := emptyDirectory(targetSubDirFullPath)
		s.subDirs = append(s.subDirs, subDir)
		s.subDirs.sortSelf()
		return subDir.addFile(nextRelativeFilePath, file)
	}
}

func (s *Directory) findFile(relativeFilePath string) (*File, error) {
	split := strings.Split(relativeFilePath, "/")
	isBareFile := len(split) == 1 // bare file (i.e.) this file is not in a directory like `main.tsx`, not `src/main.tsx`

	if isBareFile {
		// Here, `relativeFilePath` is a bare file, NOT in a directory.
		// So try to find it from files in Directory (= receiver of this method)
		for _, f := range s.files {
			if f.fileName == relativeFilePath {
				return f, nil
			}
		}

		// if not found, error
		return nil, fmt.Errorf("findFile failed, in dir = '%s' file = '%s' does not exist", s.dirPath, relativeFilePath)
	} else {
		// Here,`relativeFilePath` is a file within a directory.
		// So recursively find a file in sub directories
		targetSubDir := split[0]
		// strip the current directory, and get the next relative path to call this function recursively.
		nextRelativeFilePath := strings.Join(split[1:], "/")

		for _, subdir := range s.subDirs {
			if subdir.dirName == targetSubDir {
				return subdir.findFile(nextRelativeFilePath)
			}
		}

		// if not found, error
		return nil, fmt.Errorf("findFile failed, in dir = '%s' no directory = '%s' exists", s.dirPath, targetSubDir)
	}
}

func (s *Directory) ToGraphQLFileNode() *model.FileNode {
	//copy to avoid mutation effects afterwards
	dirType := model.FileNodeTypeDirectory
	dirPath := s.dirPath
	name := s.dirName
	offset := s.offset
	falseValue := false

	return &model.FileNode{
		NodeType:  dirType,
		FilePath:  dirPath,
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
