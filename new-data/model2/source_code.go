package model2

import (
	"fmt"
	"sort"
	"strings"
)

func (s *SourceCode) sortFileTree() {
	sort.Slice(s.FileTree, func(i, j int) bool {
		iFilePath := s.FileTree[i].FilePath
		jFilePath := s.FileTree[j].FilePath
		return LessFilePath(*iFilePath, *jFilePath)
	})
}

func (s *SourceCode) findFileNode(filePath string) (int, *FileNode) {
	for i, fn := range s.FileTree {
		if filePath == *fn.FilePath {
			return i, fn
		}
	}
	return -1, nil
}

func (s *SourceCode) validateNode(filePath string, expectedNodeType FileNodeType) error {
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

func (s *SourceCode) hasParentDir(filePath string) error {
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

func (s *SourceCode) canAddDirectoryNode(directoryPath string) error {
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

func (s *SourceCode) canDeleteDirectoryNode(directoryPath string) error {
	if err := validateFilePath(directoryPath); err != nil {
		return fmt.Errorf("cannot delete directory node, %s", err)
	}

	if err := s.validateNode(directoryPath, FileNodeTypeDirectory); err != nil {
		return fmt.Errorf("cannot delete directory node, %s", err)
	}

	return nil
}

func (s *SourceCode) canAddFileNode(filePath string) error {
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

func (s *SourceCode) canDeleteFileNode(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot delete file node, %s", err)
	}

	if err := s.validateNode(filePath, FileNodeTypeFile); err != nil {
		return fmt.Errorf("cannot delete file node, %s", err)
	}

	return nil
}

func (s *SourceCode) canUpdateFileNode(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot update file node, %s", err)
	}

	if err := s.validateNode(filePath, FileNodeTypeFile); err != nil {
		return fmt.Errorf("cannot update file node, %s", err)
	}

	return nil
}

func (s *SourceCode) canAddFileContent(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot add file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; ok {
		return fmt.Errorf("cannot add file content, file path = %s already exists", filePath)
	}

	return nil
}

func (s *SourceCode) canDeleteFileContent(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot delete file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; !ok {
		return fmt.Errorf("cannot delete file content, file path = %s is non-existent", filePath)
	}

	return nil
}

func (s *SourceCode) canUpdateFileContent(filePath string) error {
	if err := validateFilePath(filePath); err != nil {
		return fmt.Errorf("cannot update file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; !ok {
		return fmt.Errorf("cannot update file content, file path = %s is non-existent", filePath)
	}

	return nil
}

func (s *SourceCode) setAllIsUpdatedFalse() {
	falseValue := false
	for _, v := range s.FileTree {
		v.IsUpdated = &falseValue
	}
}

func (s *SourceCode) addDirectoryNode(directoryPath string) {
	s.setAllIsUpdatedFalse()
	s.FileTree = append(s.FileTree, directoryNode(directoryPath))
	s.sortFileTree()
}

func (s *SourceCode) deleteDirectoryNode(filePath string) {
	s.setAllIsUpdatedFalse()
	var newFileTree []*FileNode
	for _, v := range s.FileTree {
		if !strings.HasPrefix(*v.FilePath, filePath) {
			newFileTree = append(newFileTree, v)
		}
	}
	s.FileTree = newFileTree
	s.sortFileTree()
}

func (s *SourceCode) addFileNode(filePath string) {
	s.setAllIsUpdatedFalse()
	s.FileTree = append(s.FileTree, fileNode(filePath))
	s.sortFileTree()
}

func (s *SourceCode) deleteFileNode(filePath string) {
	s.setAllIsUpdatedFalse()
	var newFileTree []*FileNode
	for _, v := range s.FileTree {
		if *v.FilePath != filePath {
			newFileTree = append(newFileTree, v)
		}
	}
	s.FileTree = newFileTree
	s.sortFileTree()
}

func (s *SourceCode) addFileContent(filePath, content string, isFullContent bool) {
	s.FileContents[filePath] = *openFile(filePath, content)
}

func (s *SourceCode) deleteFileContent(filePath string) {
	delete(s.FileContents, filePath)
}

func (s *SourceCode) updateFileContent(filePath, content string) {
	s.FileContents[filePath] = *openFile(filePath, content)
}

// public methods

func NewSourceCode() *SourceCode {
	return &SourceCode{FileContents: make(map[string]OpenFile)}
}

func (s *SourceCode) AddDirectoryNode(directoryPath string) error {
	if err := s.canAddDirectoryNode(directoryPath); err != nil {
		return fmt.Errorf("AddDirectoryNode failed, %s", err)
	}
	s.addDirectoryNode(directoryPath)
	return nil
}

func (s *SourceCode) DeleteDirectoryNode(directoryPath string) error {
	if err := s.canDeleteDirectoryNode(directoryPath); err != nil {
		return fmt.Errorf("DeleteDirectoryNode failed, %s", err)
	}
	s.deleteDirectoryNode(directoryPath)
	return nil
}

func (s *SourceCode) AddFileNode(filePath string) error {
	if err := s.canAddFileNode(filePath); err != nil {
		return fmt.Errorf("AddFileNode failed, %s", err)
	}
	s.addFileNode(filePath)
	return nil
}

func (s *SourceCode) DeleteFileNode(filePath string) error {
	if err := s.canDeleteFileNode(filePath); err != nil {
		return fmt.Errorf("DeleteFileNode failed, %s", err)
	}
	s.deleteFileNode(filePath)
	return nil
}

func (s *SourceCode) AddFileContent(filePath, content string, isFullContent bool) error {
	if err := s.canAddFileContent(filePath); err != nil {
		return fmt.Errorf("AddFileContent failed, %s", err)
	}
	s.addFileContent(filePath, content, isFullContent)
	return nil
}

func (s *SourceCode) DeleteFileContent(filePath string) error {
	if err := s.canDeleteFileContent(filePath); err != nil {
		return fmt.Errorf("DeleteFileContent failed, %s", err)
	}
	s.deleteFileContent(filePath)
	return nil
}

func (s *SourceCode) UpdateFileContent(filePath, content string) error {
	if err := s.canUpdateFileContent(filePath); err != nil {
		return fmt.Errorf("UpdateFileContent failed, %s", err)
	}
	s.updateFileContent(filePath, content)
	return nil
}

func (s *SourceCode) ApplyEffect(effect SourceCodeEffect) error {
	// Add directories
	errors := []string{}
	for _, d := range effect.DirectoriesToAdd {
		if err := s.AddDirectoryNode(d.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("failed to apply effect: %s", strings.Join(errors, ", "))
	}

	// Delete directories
	// if you come here, len(errors) = 0
	for _, d := range effect.DirectoriesToDelete {
		if err := s.DeleteDirectoryNode(d.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("failed to apply effect: %s", strings.Join(errors, ", "))
	}

	// Add files
	// if you come here, len(errors) = 0
	for _, f := range effect.FilesToAdd {
		if err := s.AddFileNode(f.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
		if err := s.AddFileContent(f.FilePath, f.Content, f.IsFullContent); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("failed to apply effect: %s", strings.Join(errors, ", "))
	}

	// Update files
	// if you come here, len(errors) = 0
	for _, f := range effect.FilesToUpdate {
		if err := s.UpdateFileContent(f.FilePath, f.Content); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("failed to apply effect: %s", strings.Join(errors, ", "))
	}

	// Delete files
	// if you come here, len(errors) = 0
	for _, f := range effect.FilesToDelete {
		if err := s.DeleteFileContent(f.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
		if err := s.DeleteFileNode(f.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("failed to apply effect: %s", strings.Join(errors, ", "))
	}
	return nil
}
