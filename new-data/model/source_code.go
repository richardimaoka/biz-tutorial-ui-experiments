package model

import (
	"fmt"
	"sort"
	"strings"
)

func (s *SourceCode) findFileNode(filePath string) (int, *FileNode) {
	for i, fn := range s.FileTree {
		if filePath == *fn.FilePath {
			return i, fn
		}
	}
	return -1, nil
}

func (s *SourceCode) isValidNode(filePath string, expectedNodeType FileNodeType) error {
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

	err := s.isValidNode(parentPath, FileNodeTypeDirectory)
	if err != nil {
		return fmt.Errorf("no parent dir, %s", err)
	}

	return nil
}

// canXxxYyy pre-condition checks

func (s *SourceCode) canAddDirectory(directoryPath string) error {
	if err := isValidFilePath(directoryPath); err != nil {
		return fmt.Errorf("cannot add directory, %s", err)
	}

	if _, node := s.findFileNode(directoryPath); node != nil {
		return fmt.Errorf("cannot add directory, file path = %s already exists", directoryPath)
	}

	// if err := s.hasParentDir(directoryPath); err != nil {
	// 	return err
	// }
	return nil
}

func (s *SourceCode) canDeleteDirectory(directoryPath string) error {
	if err := isValidFilePath(directoryPath); err != nil {
		return fmt.Errorf("cannot delete directory, %s", err)
	}

	if err := s.isValidNode(directoryPath, FileNodeTypeDirectory); err != nil {
		return fmt.Errorf("cannot delete directory, %s", err)
	}

	return nil
}

func (s *SourceCode) canAddFileNode(filePath string) error {
	if err := isValidFilePath(filePath); err != nil {
		return fmt.Errorf("cannot add file node, %s", err)
	}

	if _, node := s.findFileNode(filePath); node != nil {
		return fmt.Errorf("cannot add file node, file path = %s already exists", filePath)
	}

	// if err := s.hasParentDir(filePath); err != nil {
	// 	return fmt.Errorf("cannot add file node, %s", err)
	// }
	return nil
}

func (s *SourceCode) canDeleteFileNode(filePath string) error {
	if err := isValidFilePath(filePath); err != nil {
		return fmt.Errorf("cannot delete file node, %s", err)
	}

	if err := s.isValidNode(filePath, FileNodeTypeFile); err != nil {
		return fmt.Errorf("cannot delete file node, %s", err)
	}

	return nil
}

func (s *SourceCode) canUpdateFileNode(filePath string) error {
	if err := isValidFilePath(filePath); err != nil {
		return fmt.Errorf("cannot update file node, %s", err)
	}

	if err := s.isValidNode(filePath, FileNodeTypeFile); err != nil {
		return fmt.Errorf("cannot update file node, %s", err)
	}

	return nil
}

func (s *SourceCode) canAddFileContent(filePath string) error {
	if err := isValidFilePath(filePath); err != nil {
		return fmt.Errorf("cannot add file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; ok {
		return fmt.Errorf("cannot add file content, file path = %s already exists", filePath)
	}

	return nil
}

func (s *SourceCode) canDeleteFileContent(filePath string) error {
	if err := isValidFilePath(filePath); err != nil {
		return fmt.Errorf("cannot delete file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; !ok {
		return fmt.Errorf("cannot delete file content, file path = %s is non-existent", filePath)
	}

	return nil
}

func (s *SourceCode) canUpdateFileContent(filePath string) error {
	if err := isValidFilePath(filePath); err != nil {
		return fmt.Errorf("cannot update file content, %s", err)
	}

	if _, ok := s.FileContents[filePath]; !ok {
		return fmt.Errorf("cannot update file content, file path = %s is non-existent", filePath)
	}

	return nil
}

// canXxxFile pre-condition checks

func (s *SourceCode) canAddFile(op FileAdd) error {
	if err := s.canAddFileNode(op.FilePath); err != nil {
		return fmt.Errorf("cannot add file, %s", err)
	}
	if err := s.canAddFileContent(op.FilePath); err != nil {
		return fmt.Errorf("cannot add file, %s", err)
	}

	return nil
}

func (s *SourceCode) canUpdateFile(op FileUpdate) error {
	if err := s.canUpdateFileNode(op.FilePath); err != nil {
		return fmt.Errorf("cannot update file, %s", err)
	}
	if err := s.canUpdateFileContent(op.FilePath); err != nil {
		return fmt.Errorf("cannot update file, %s", err)
	}

	return nil
}

func (s *SourceCode) canDeleteFile(op FileDelete) error {
	if err := s.canDeleteFileNode(op.FilePath); err != nil {
		return fmt.Errorf("cannot delete file, %s", err)
	}
	if err := s.canDeleteFileContent(op.FilePath); err != nil {
		return fmt.Errorf("cannot delete file, %s", err)
	}

	return nil
}

func (s *SourceCode) canApplyDiff(diff GitDiff) error {
	// pre-condition check, dupe in diff
	if diffDuplicate := diff.findDuplicate(); diffDuplicate.size() > 0 {
		return fmt.Errorf("cannot apply diff, duplicate file paths in diff = %+v", diffDuplicate)
	}

	// pre-condition check, each element's pre-condition check
	errors := []string{}
	for _, f := range diff.Added {
		if err := s.canAddFile(f); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, f := range diff.Updated {
		if err := s.canUpdateFile(f); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, f := range diff.Deleted {
		if err := s.canDeleteFile(f); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("cannot apply diff, %s", strings.Join(errors, ", "))
	}

	return nil
}

// internal mutation methods

func (s *SourceCode) preMutation() {
	//set all IsUpdated to false
	falseValue := false
	for _, v := range s.FileTree {
		v.IsUpdated = &falseValue
	}
}

func (s *SourceCode) postMutation() {
	// soft fileTree
	sort.Slice(s.FileTree, func(i, j int) bool {
		iFilePath := s.FileTree[i].FilePath
		jFilePath := s.FileTree[j].FilePath
		return LessFilePath(*iFilePath, *jFilePath)
	})
}

func (s *SourceCode) addMissingParentDirs(directoryPath string) {
	splitPath := strings.Split(directoryPath, "/")
	incremental := []string{}
	for i, c := range splitPath {
		// skip the last in splitPath
		if i == len(splitPath)-1 {
			continue
		}
		incremental = append(incremental, c)
		incrementalPath := strings.Join(incremental, "/")
		if found, _ := s.findFileNode(incrementalPath); found == -1 {
			s.appendDirectoryNode(incrementalPath)
		}
	}
}

func (s *SourceCode) popNode(filePath string) {
	var newFileTree []*FileNode
	for _, v := range s.FileTree {
		if !strings.HasPrefix(*v.FilePath, filePath) {
			newFileTree = append(newFileTree, v)
		}
	}
	s.FileTree = newFileTree
}

func (s *SourceCode) appendDirectoryNode(directoryPath string) {
	s.FileTree = append(s.FileTree, directoryNode(directoryPath))
}

func (s *SourceCode) addDirectoryNode(directoryPath string) {
	s.addMissingParentDirs(directoryPath)
	s.appendDirectoryNode(directoryPath)
}

func (s *SourceCode) deleteDirectoryNode(filePath string) {
	s.popNode(filePath)
}

func (s *SourceCode) addFileNode(filePath string) {
	s.addMissingParentDirs(filePath)
	s.FileTree = append(s.FileTree, fileNode(filePath))
}

func (s *SourceCode) deleteFileNode(filePath string) {
	s.popNode(filePath)
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

func (s *SourceCode) applyDiff(diff GitDiff) {
	for _, op := range diff.Added {
		s.addFileNode(op.FilePath)
		s.addFileContent(op.FilePath, op.Content, op.IsFullContent)
	}
	for _, op := range diff.Updated {
		s.updateFileContent(op.FilePath, op.Content)
	}
	for _, op := range diff.Deleted {
		s.deleteFileContent(op.FilePath)
		s.deleteFileNode(op.FilePath)
	}
}

// public methods

func NewSourceCode() *SourceCode {
	return &SourceCode{FileContents: make(map[string]OpenFile)}
}

func (s *SourceCode) AddDirectory(op DirectoryAdd) error {
	if err := s.canAddDirectory(op.FilePath); err != nil {
		return fmt.Errorf("AddDirectory failed, %s", err)
	}

	s.preMutation()
	s.addDirectoryNode(op.FilePath)
	s.postMutation()

	return nil
}

func (s *SourceCode) DeleteDirectory(op DirectoryDelete) error {
	if err := s.canDeleteDirectory(op.FilePath); err != nil {
		return fmt.Errorf("DeleteDirectoryNode failed, %s", err)
	}

	s.preMutation()
	s.deleteDirectoryNode(op.FilePath)
	s.postMutation()

	return nil
}

func (s *SourceCode) AddFile(op FileAdd) error {
	if err := s.canAddFile(op); err != nil {
		return fmt.Errorf("AddFile failed, %s", err)
	}

	s.preMutation()
	s.addFileNode(op.FilePath)
	s.addFileContent(op.FilePath, op.Content, op.IsFullContent)
	s.postMutation()

	return nil
}

func (s *SourceCode) UpdateFile(op FileUpdate) error {
	if err := s.canUpdateFile(op); err != nil {
		return fmt.Errorf("UpdateFile failed, %s", err)
	}

	//no need to execute preMutation() / postMutation()
	s.updateFileContent(op.FilePath, op.Content)

	return nil
}

func (s *SourceCode) DeleteFile(op FileDelete) error {
	if err := s.canDeleteFile(op); err != nil {
		return fmt.Errorf("DeleteFile failed, %s", err)
	}

	s.preMutation()
	s.deleteFileContent(op.FilePath)
	s.deleteFileNode(op.FilePath)
	s.postMutation()

	return nil
}

func (s *SourceCode) ApplyDiff(diff GitDiff) error {
	if err := s.canApplyDiff(diff); err != nil {
		return fmt.Errorf("failed to apply diff, %s", err)
	}

	s.preMutation()
	s.applyDiff(diff)
	s.postMutation()

	return nil
}
