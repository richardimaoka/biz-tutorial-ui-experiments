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

func validateFilePath(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("empty path")
	}
	if strings.HasSuffix(filePath, "/") {
		return fmt.Errorf("directory path = %s ends in slash", filePath)
	}
	return nil
}

func (s *SourceCodeExtended) sortFileTree() {
	sort.Slice(s.FileTree, func(i, j int) bool {
		return lessFilePath(s.FileTree[i].FilePath, s.FileTree[j].FilePath)
	})
}

func (s *SourceCodeExtended) findFileNode(filePath string) (int, *FileNode) {
	for i, fn := range s.FileTree {
		if filePath == fn.FilePathString() {
			return i, fn
		}
	}

	return -1, nil
}

func (s *SourceCodeExtended) hasParentDir(filePath string) error {
	parentPath := parentDirectoryPath(filePath)
	if parentPath == "" {
		return nil //parent dir = root dir
	}

	_, parentNode := s.findFileNode(parentPath)
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

func (s *SourceCodeExtended) canAddDirectory(directoryPath string) error {
	if err := validateFilePath(directoryPath); err != nil {
		return fmt.Errorf("cannot add directory, %s", err)
	}

	if _, node := s.findFileNode(directoryPath); node != nil {
		return fmt.Errorf("cannot add directory, file path = %s already exists", directoryPath)
	}

	if err := s.hasParentDir(directoryPath); err != nil {
		return err
	}
	return nil
}

func (s *SourceCodeExtended) canAddFile(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot add file, %s", err)
	}

	if _, node := s.findFileNode(filePath); node != nil {
		return fmt.Errorf("cannot add file, file path = %s already exists", filePath)
	}

	if err := s.hasParentDir(filePath); err != nil {
		return fmt.Errorf("cannot add file, %s", err)
	}
	return nil
}

func (s *SourceCodeExtended) canDeleteFile(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot delete file, %s", err)
	}

	_, node := s.findFileNode(filePath)
	if node == nil {
		return fmt.Errorf("cannot add file, file path = %s does not exists", filePath)
	}

	if node.NodeType == nil {
		return fmt.Errorf("file path = %s has nil node type", filePath)
	} else if *node.NodeType == FileNodeTypeDirectory {
		return fmt.Errorf("file path = %s is a directory, not file", filePath)
	} else if *node.NodeType == FileNodeTypeFile {
		return nil
	} else {
		return fmt.Errorf("file path = %s has unkown node type = %s", filePath, *node.NodeType)
	}
}

func (s *SourceCodeExtended) canAddFileContent(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot add file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; ok {
		return fmt.Errorf("cannot add file content, file path = %s already exists", filePath)
	}

	return nil
}

func (s *SourceCodeExtended) addDirectory(directoryPath string) error {
	if err := s.canAddDirectory(directoryPath); err != nil {
		return fmt.Errorf("addDirectory failed, %s", err)
	}

	s.FileTree = append(s.FileTree, directoryNode(directoryPath))
	s.sortFileTree()

	return nil
}

func (s *SourceCodeExtended) addFile(filePath string) error {
	if err := s.canAddFile(filePath); err != nil {
		return fmt.Errorf("addFile failed, %s", err)
	}

	s.FileTree = append(s.FileTree, fileNode(filePath))
	s.sortFileTree()

	return nil
}

func (s *SourceCodeExtended) addFileContent(filePath, content string) error {
	if err := s.canAddFileContent(filePath); err != nil {
		return fmt.Errorf("addFile failed, %s", err)
	}

	s.FileTree = append(s.FileTree, fileNode(filePath))
	s.sortFileTree()

	return nil
}

func (s *SourceCodeExtended) deleteFile(filePath string) error {
	if err := s.canDeleteFile(filePath); err != nil {
		return fmt.Errorf("addFile failed, %s", err)
	}

	i, _ := s.findFileNode(filePath)
	if len(s.FileTree) == 1 && i != -1 {
		s.FileTree = nil
	} else {
		s.FileTree = append(s.FileTree[:i], s.FileTree[i+1:]...)
	}
	s.sortFileTree()

	return nil
}
