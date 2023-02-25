package model2

import (
	"fmt"
	"sort"
	"strings"
)

type AddDirectory struct {
	FilePath string
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

func parentDirectoryPath(filePath string) string {
	split := strings.Split(filePath, "/")
	return strings.Join(split[:len(split)-1], "")
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

func (s *SourceCode) hasParentDir(filePath string) error {
	parentPath := parentDirectoryPath(filePath)
	if parentPath == "" {
		return nil //parent dir = root dir
	}

	parentNode := s.findFileNode(parentPath)
	if parentNode == nil {
		return fmt.Errorf("parent path = %s has no directory", parentPath)
	} else if parentNode.NodeType == nil {
		return fmt.Errorf("parent path = %s has nil node type", parentPath)
	} else if *parentNode.NodeType == FileNodeTypeFile {
		return fmt.Errorf("parent path = %s is a file node, not directory", parentPath)
	} else {
		return nil
	}
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

	if err := s.hasParentDir(directoryPath); err != nil {
		return err
	}
	return nil
}

func (s *SourceCode) canAddFile(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("cannot add file with empty path")
	}
	if strings.HasSuffix(filePath, "/") {
		return fmt.Errorf("file path = %s ends in slash", filePath)
	}

	if s.findFileNode(filePath) != nil {
		return fmt.Errorf("file path = %s already exists", filePath)
	}

	if err := s.hasParentDir(filePath); err != nil {
		return err
	}
	return nil
}

func (s *SourceCode) addDirectory(directoryPath string) error {
	if err := s.canAddDirectory(directoryPath); err != nil {
		return fmt.Errorf("addDirectory failed, %s", err)
	}

	s.FileTree = append(s.FileTree, directoryNode(directoryPath))
	s.sortFileTree()

	return nil
}

func (s *SourceCode) addFile(filePath string) error {
	if err := s.canAddFile(filePath); err != nil {
		return fmt.Errorf("addFile failed, %s", err)
	}

	s.FileTree = append(s.FileTree, fileNode(filePath))
	s.sortFileTree()

	return nil
}
