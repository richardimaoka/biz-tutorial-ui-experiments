package model

func directoryNode(filePath string) *FileNode {
	nodeType := FileNodeTypeDirectory
	split := filePathPtrSlice(filePath)
	offset := len(split) - 1
	trueValue := true

	node := FileNode{
		NodeType:  &nodeType,
		Name:      split[len(split)-1],
		FilePath:  &filePath,
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
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &trueValue,
	}
	return &node
}
