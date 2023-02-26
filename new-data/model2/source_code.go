package model2

import (
	"fmt"
	"sort"
	"strings"
)

type SourceCodeExtended struct {
	SourceCode
	FileContents map[string]OpenFile `json:"fileContents"`
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

func (s *SourceCodeExtended) validateNode(filePath string, expectedNodeType FileNodeType) error {
	_, node := s.findFileNode(filePath)
	if node == nil {
		return fmt.Errorf("filePath = %s has no node", filePath)
	} else if node.NodeType == nil {
		return fmt.Errorf("filePath = %s has nil node type", filePath)
	} else if *node.NodeType != expectedNodeType {
		return fmt.Errorf("filePath = %s has node type = %s, but expected %s", filePath, node.NodeType, expectedNodeType)
	} else {
		return nil
	}
}

func (s *SourceCodeExtended) hasParentDir(filePath string) error {
	parentPath := parentDirectoryPath(filePath)
	if parentPath == "" {
		return nil //parent dir = root dir
	}

	err := s.validateNode(parentPath, FileNodeTypeDirectory)
	if err != nil {
		return fmt.Errorf("no parent dir, %s", err)
	}

	return nil
}

func (s *SourceCodeExtended) canAddDirectoryNode(directoryPath string) error {
	if err := validateFilePath(directoryPath); err != nil {
		return fmt.Errorf("cannot add directory node, %s", err)
	}

	if _, node := s.findFileNode(directoryPath); node != nil {
		return fmt.Errorf("cannot add directory node, file path = %s already exists", directoryPath)
	}

	if err := s.hasParentDir(directoryPath); err != nil {
		return err
	}
	return nil
}

func (s *SourceCodeExtended) canDeleteDirectoryNode(directoryPath string) error {
	if err := validateFilePath(directoryPath); err != nil {
		return fmt.Errorf("cannot delete directory node, %s", err)
	}

	if err := s.validateNode(directoryPath, FileNodeTypeDirectory); err != nil {
		return fmt.Errorf("cannot delete directory node, %s", err)
	}

	return nil
}

func (s *SourceCodeExtended) canAddFileNode(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot add file node, %s", err)
	}

	if _, node := s.findFileNode(filePath); node != nil {
		return fmt.Errorf("cannot add file node, file path = %s already exists", filePath)
	}

	if err := s.hasParentDir(filePath); err != nil {
		return fmt.Errorf("cannot add file node, %s", err)
	}
	return nil
}

func (s *SourceCodeExtended) canDeleteFileNode(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot delete file node, %s", err)
	}

	if err := s.validateNode(filePath, FileNodeTypeFile); err != nil {
		return fmt.Errorf("cannot delete file node, %s", err)
	}

	return nil
}

func (s *SourceCodeExtended) canUpdateFileNode(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot update file node, %s", err)
	}

	if err := s.validateNode(filePath, FileNodeTypeFile); err != nil {
		return fmt.Errorf("cannot update file node, %s", err)
	}

	return nil
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

func (s *SourceCodeExtended) canDeleteFileContent(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot delete file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; !ok {
		return fmt.Errorf("cannot delete file content, file path = %s is non-existent", filePath)
	}

	return nil
}

func (s *SourceCodeExtended) canUpdateFileContent(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot update file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; !ok {
		return fmt.Errorf("cannot update file content, file path = %s is non-existent", filePath)
	}

	return nil
}

func (s *SourceCodeExtended) setAllIsUpdatedFalse() {
	falseValue := false
	for _, v := range s.FileTree {
		v.IsUpdated = &falseValue
	}
}

// public methods
func NewSourceCode() *SourceCodeExtended {
	return &SourceCodeExtended{}
}

func (s *SourceCodeExtended) AddDirectoryNode(directoryPath string) error {
	if err := s.canAddDirectoryNode(directoryPath); err != nil {
		return fmt.Errorf("AddDirectoryNode failed, %s", err)
	}

	s.setAllIsUpdatedFalse()
	s.FileTree = append(s.FileTree, directoryNode(directoryPath))
	s.sortFileTree()

	return nil
}

func (s *SourceCodeExtended) DeleteDirectoryNode(filePath string) error {
	if err := s.canDeleteDirectoryNode(filePath); err != nil {
		return fmt.Errorf("DeleteDirectoryNode failed, %s", err)
	}

	s.setAllIsUpdatedFalse()
	var newFileTree []*FileNode
	for _, v := range s.FileTree {
		if !strings.HasPrefix(v.FilePathString(), filePath) {
			newFileTree = append(newFileTree, v)
		}
	}
	s.FileTree = newFileTree
	s.sortFileTree()

	return nil
}

func (s *SourceCodeExtended) AddFileNode(filePath string) error {
	if err := s.canAddFileNode(filePath); err != nil {
		return fmt.Errorf("AddFileNode failed, %s", err)
	}

	s.setAllIsUpdatedFalse()
	s.FileTree = append(s.FileTree, fileNode(filePath))
	s.sortFileTree()

	return nil
}

func (s *SourceCodeExtended) DeleteFileNode(filePath string) error {
	if err := s.canDeleteFileNode(filePath); err != nil {
		return fmt.Errorf("DeleteFileNode failed, %s", err)
	}

	s.setAllIsUpdatedFalse()
	var newFileTree []*FileNode
	for _, v := range s.FileTree {
		if v.FilePathString() != filePath {
			newFileTree = append(newFileTree, v)
		}
	}
	s.FileTree = newFileTree
	s.sortFileTree()

	return nil
}

func (s *SourceCodeExtended) AddFileContent(filePath, content string) error {
	if err := s.canAddFileContent(filePath); err != nil {
		return fmt.Errorf("AddFileContent failed, %s", err)
	}

	s.FileContents[filePath] = *openFile(filePath, content)
	return nil
}

func (s *SourceCodeExtended) DeleteFileContent(filePath, content string) error {
	if err := s.canDeleteFileContent(filePath); err != nil {
		return fmt.Errorf("DeleteFileContent failed, %s", err)
	}

	delete(s.FileContents, filePath)
	return nil
}

func (s *SourceCodeExtended) UpdateFileContent(filePath, content string) error {
	if err := s.canUpdateFileContent(filePath); err != nil {
		return fmt.Errorf("UpdateFileContent failed, %s", err)
	}

	s.FileContents[filePath] = *openFile(filePath, content)
	return nil
}
