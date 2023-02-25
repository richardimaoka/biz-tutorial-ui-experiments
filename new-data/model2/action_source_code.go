package model2

import (
	"fmt"
	"sort"
	"strings"
)

type AddDirectory struct {
	FilePath string
}

func (a AddDirectory) filePathPtrSlice() []*string {
	split := strings.Split(a.FilePath, "/")

	var filePathSlice []*string
	for _, v := range split {
		filePathSlice = append(filePathSlice, &v)
	}

	return filePathSlice
}

func (a AddDirectory) toFileNode() *FileNode {
	dType := FileNodeTypeDirectory
	split := a.filePathPtrSlice()
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

func lessFilePath(a, b []*string) bool {
	if len(a) == 0 && len(b) == 0 {
		return false //even if len(b) == 0
	} else if /* len (a) != 0 && */ len(b) == 0 {
		return false //here, len(a) != 0
	} else if len(a) == 0 /* && len (b) != 0 */ {
		return false //even if len(b) == 0
	}

	// now len(a) != 0 AND len(b) != 0

	if a[0] == b[0] {
		return lessFilePath(a[1:], b[1:])
	} else {
		return *a[0] < *b[0]
	}
}

func (s *SourceCode) sortFileTree() {
	sort.Slice(s.FileTree, func(i, j int) bool {
		return lessFilePath(s.FileTree[i].FilePath, s.FileTree[j].FilePath)
	})
}

func (s *SourceCode) findFileNode(filePath string) *FileNode {
	for _, fn := range s.FileTree {
		if filePath == fn.FilePathString() {
			return fn
		}
	}

	return nil
}

func (s *SourceCode) existsParentDir(filePath string) bool {
	split := strings.Split(filePath, "/")
	parentPathSlice := split[:len(split)-1]

	if len(parentPathSlice) == 0 {
		return true
	} else {
		parentPath := strings.Join(parentPathSlice, "/")
		return s.findFileNode(parentPath) != nil
	}
}

func (s *SourceCode) addDirectory(add AddDirectory) error {
	if s.findFileNode(add.FilePath) != nil {
		return fmt.Errorf("filePath = %s already exists", add.FilePath)
	}
	if !s.existsParentDir(add.FilePath) {
		return fmt.Errorf("filePath = %s already exists", add.FilePath)
	}

	s.FileTree = append(s.FileTree, add.toFileNode())
	s.sortFileTree()

	return nil
}
