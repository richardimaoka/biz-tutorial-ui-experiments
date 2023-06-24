package processing

type FileHighlight = fileHighlight

var CalcHighlight = calcHighlight

type FileProcessorNode = fileProcessorNode

type DirectoryProcessorNode = directoryProcessorNode

type FileTreeNode = fileTreeNode

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
