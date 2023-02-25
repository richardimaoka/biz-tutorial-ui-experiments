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

func parentDirectoryPath(filePath string) string {
	split := strings.Split(filePath, "/")
	return strings.Join(split[:len(split)-1], "")
}

func (s *SourceCode) canAddDirectory(directoryPath string) error {
	if directoryPath == "" {
		return fmt.Errorf("cannot add directory with empty path")
	}
	if strings.HasSuffix(directoryPath, "/") {
		return fmt.Errorf("directory path = %s ends in slash", directoryPath)
	}

	if s.findFileNode(directoryPath) != nil {
		return fmt.Errorf("file path = %s already exists", directoryPath)
	}

	parentPath := parentDirectoryPath(directoryPath)
	if parentPath == "" {
		return nil //parent dir = root dir
	}

	node := s.findFileNode(parentPath)
	if node.NodeType == nil {
		return fmt.Errorf("parent path = %s has nil node type", parentPath)
	} else if *node.NodeType == FileNodeTypeFile {
		return fmt.Errorf("parent path = %s is a file node, not directory", parentPath)
	} else {
		return nil
	}
}

func (s *SourceCode) addDirectory(directoryPath string) error {
	if err := s.canAddDirectory(directoryPath); err != nil {
		return fmt.Errorf("addDirectory failed, %s", err)
	}

	s.FileTree = append(s.FileTree, directoryNode(directoryPath))
	s.sortFileTree()

	return nil
}

// func (s *SourceCode) addDirectory(directoryPath string) error {
// 	if s.findFileNode(directoryPath) != nil {
// 		return fmt.Errorf("filePath = %s already exists", directoryPath)
// 	}
// 	if !s.existsParentDir(directoryPath) {
// 		return fmt.Errorf("parent directory = %s does not exist", directoryPath)
// 	}

// 	s.FileTree = append(s.FileTree, add.toFileNode())
// 	s.sortFileTree()

// 	return nil
// }
