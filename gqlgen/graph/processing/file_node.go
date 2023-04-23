package processing

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type nodeType string

const (
	fileType      nodeType = "FILE"
	directoryType nodeType = "DIRECTORY"
)

type fileTreeNode interface {
	NodeType() nodeType
	FilePath() string
	IsUpdated() bool
	SetIsUpdated(isUpdated bool)
	//TODO: Clone() fileTreeNode
}

type fileProcessorNode struct {
	filePath  string
	content   string
	isUpdated bool
}

type directoryProcessorNode struct {
	filePath  string
	isUpdated bool
}

func (n *fileProcessorNode) NodeType() nodeType {
	return fileType
}

func (n *directoryProcessorNode) NodeType() nodeType {
	return directoryType
}

func (n *fileProcessorNode) FilePath() string {
	return n.filePath
}

func (n *directoryProcessorNode) FilePath() string {
	return n.filePath
}

func (n *fileProcessorNode) IsUpdated() bool {
	return n.isUpdated
}

func (n *directoryProcessorNode) IsUpdated() bool {
	return n.isUpdated
}

func (n *fileProcessorNode) SetIsUpdated(isUpdated bool) {
	n.isUpdated = isUpdated
}

func (n *directoryProcessorNode) SetIsUpdated(isUpdated bool) {
	n.isUpdated = isUpdated
}

func createDirectoryNode(filePath string, isUpdated bool) *model.FileNode {
	nodeType := model.FileNodeTypeDirectory
	split := filePathPtrSlice(filePath)
	offset := len(split) - 1

	node := model.FileNode{
		NodeType:  &nodeType,
		Name:      split[len(split)-1],
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
	return &node
}

func createFileNode(filePath string) *model.FileNode {
	nodeType := model.FileNodeTypeFile
	split := filePathPtrSlice(filePath)
	offset := len(split) - 1
	trueValue := true

	node := model.FileNode{
		NodeType:  &nodeType,
		Name:      split[len(split)-1],
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &trueValue,
	}
	return &node
}
