package model2

import "strings"

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

func (s *SourceCode) addDirectory(add AddDirectory) error {
	s.FileTree = append(s.FileTree, add.toFileNode())
	return nil
}
