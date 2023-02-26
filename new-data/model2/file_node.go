package model2

import "strings"

func directoryNode(filePath string) *FileNode {
	nodeType := FileNodeTypeDirectory
	split := filePathPtrSlice(filePath)
	offset := len(split) - 1
	trueValue := true

	node := FileNode{
		NodeType:  &nodeType,
		Name:      split[len(split)-1],
		FilePath:  split,
		Offset:    &offset,
		IsUpdated: &trueValue,
	}
	return &node
}

func fileNode(filePath string) *FileNode {
	nodeType := FileNodeTypeFile
	split := filePathPtrSlice(filePath)
	offset := len(split) - 1
	trueValue := true

	node := FileNode{
		NodeType:  &nodeType,
		Name:      split[len(split)-1],
		FilePath:  split,
		Offset:    &offset,
		IsUpdated: &trueValue,
	}
	return &node
}

func (f *FileNode) FilePathString() string {
	var s []string
	for _, v := range f.FilePath {
		s = append(s, *v)
	}
	return strings.Join(s, "/")
}
