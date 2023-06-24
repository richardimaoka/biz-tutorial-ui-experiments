package processing

func NewFileProcessorNode(filePath string) *FileProcessorNode {
	return &FileProcessorNode{
		filePath: filePath,
	}
}

func NewDirectoryProcessorNode(filePath string) *DirectoryProcessorNode {
	return &DirectoryProcessorNode{
		filePath: filePath,
	}
}
