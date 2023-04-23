package processing

import (
	"fmt"
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
	SetIsUpdated(isUpdated bool)
}

type fileProcessorNode struct {
	filePath  string
	content   string
	isUpdated bool
}

type directoryProcessorNode struct {
	filePath  string
	children  map[string]fileTreeNode
	isUpdated bool
}

func (f *fileProcessorNode) NodeType() nodeType {
	return fileType
}

func (f *directoryProcessorNode) NodeType() nodeType {
	return directoryType
}

func (f *fileProcessorNode) FilePath() string {
	return f.filePath
}

func (f *directoryProcessorNode) FilePath() string {
	return f.filePath
}

func (f *fileProcessorNode) SetIsUpdated(isUpdated bool) {
	f.isUpdated = isUpdated
}

func (f *directoryProcessorNode) SetIsUpdated(isUpdated bool) {
	f.isUpdated = isUpdated
}

func isValidFilePath(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("empty path")
	}
	if strings.HasSuffix(filePath, "/") {
		return fmt.Errorf("directory path = %s ends in slash", filePath)
	}
	return nil
}

func filePathPtrSlice(filePath string) []*string {
	split := strings.Split(filePath, "/")

	var filePathSlice []*string
	for i := range split {
		filePathSlice = append(filePathSlice, &split[i]) // cannot use v of `for i, v := range ...` because v has the same address throughout the loop
	}

	return filePathSlice
}

func createDirectoryNode(filePath string) *model.FileNode {
	nodeType := model.FileNodeTypeDirectory
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
