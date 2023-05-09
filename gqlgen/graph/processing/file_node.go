package processing

import (
	"strings"

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
	ToGraphQLNode() *model.FileNode
	Clone() fileTreeNode
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

func (n *fileProcessorNode) ToGraphQLNode() *model.FileNode {
	filePath := n.filePath   // copy to avoid mutation effect afterwards
	isUpdated := n.isUpdated // copy to avoid mutation effect afterwards
	nodeType := model.FileNodeTypeFile
	split := strings.Split(filePath, "/")
	offset := len(split) - 1

	return &model.FileNode{
		NodeType:  &nodeType,
		Name:      &split[len(split)-1],
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}

func (n *directoryProcessorNode) ToGraphQLNode() *model.FileNode {
	filePath := n.filePath   // copy to avoid mutation effect afterwards
	isUpdated := n.isUpdated // copy to avoid mutation effect afterwards
	nodeType := model.FileNodeTypeDirectory
	split := strings.Split(filePath, "/")
	offset := len(split) - 1

	return &model.FileNode{
		NodeType:  &nodeType,
		Name:      &split[len(split)-1],
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}

// TODO: can this be deleted?? check and delete if possible
func (n *fileProcessorNode) Clone() fileTreeNode {
	copied := *n // copy to avoid mutation effect afterwards
	return &copied
}

func (n *directoryProcessorNode) Clone() fileTreeNode {
	copied := *n // copy to avoid mutation effect afterwards
	return &copied
}

func (n *fileProcessorNode) ToGraphQLOpenFile() *model.OpenFile {
	filePath := n.filePath // copy to avoid mutation effect afterwards
	content := n.content   // copy to avoid mutation effect afterwards
	isFullContent := true
	split := strings.Split(filePath, "/")

	return &model.OpenFile{
		FilePath:      &filePath,
		FileName:      &split[len(split)-1],
		Content:       &content,
		IsFullContent: &isFullContent,
		Language:      nil,
		Highlight:     nil,
	}
}
