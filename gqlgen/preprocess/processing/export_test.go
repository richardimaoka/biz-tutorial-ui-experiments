package processing

type FileProcessorNode = fileProcessorNode

type DirectoryProcessorNode = directoryProcessorNode

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
