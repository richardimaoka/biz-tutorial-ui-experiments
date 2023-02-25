package model2

import "strings"

func (f *FileNode) FilePathString() string {
	var s []string
	for _, v := range f.FilePath {
		s = append(s, *v)
	}
	return strings.Join(s, "/")
}

func filePathPtrSlice(filePath string) []*string {
	split := strings.Split(filePath, "/")

	var filePathSlice []*string
	for _, v := range split {
		filePathSlice = append(filePathSlice, &v)
	}

	return filePathSlice
}

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
