package model2

import (
	"fmt"
	"sort"
	"strings"
)

type AddDirectory struct {
	FilePath string
}

func splitFilePath(filePath string) []*string {
	split := strings.Split(filePath, "/")

	var filePathSlice []*string
	for _, v := range split {
		filePathSlice = append(filePathSlice, &v)
	}

	return filePathSlice
}

func (a AddDirectory) toFileNode() *FileNode {
	dType := FileNodeTypeDirectory
	split := splitFilePath(a.FilePath)
	offset := len(split) - 1
	trueValue := true

	fileNode := FileNode{
		NodeType:  &dType,
		Name:      split[len(split)-1],
		FilePath:  split,
		Offset:    &offset,
		IsUpdated: &trueValue,
	}
	return &fileNode
}

func filePathLess(a, b []*string) bool {
	if len(a) == 0 && len(b) == 0 {
		return false //even if len(b) == 0
	} else if /* len (a) != 0 && */ len(b) == 0 {
		return false //here, len(a) != 0
	} else if len(a) == 0 /* && len (b) != 0 */ {
		return false //even if len(b) == 0
	}

	// now len(a) != 0 AND len(b) != 0

	if a[0] == b[0] {
		return filePathLess(a[1:], b[1:])
	} else {
		return *a[0] < *b[0]
	}
}

func (s *SourceCode) sortFileTree() {
	sort.Slice(s.FileTree, func(i, j int) bool {
		return filePathLess(s.FileTree[i].FilePath, s.FileTree[j].FilePath)
	})
}

func (s *SourceCode) filePathExists(add AddDirectory) bool {
	for _, f := range s.FileTree {
		if add.FilePath == f.FilePathString() {
			return true
		}
	}
	return false
}

func (s *SourceCode) addDirectory(add AddDirectory) error {
	if s.filePathExists(add) {
		return fmt.Errorf("filePath = %s already exists", add.FilePath)
	}

	s.FileTree = append(s.FileTree, add.toFileNode())
	s.sortFileTree()

	return nil
}
