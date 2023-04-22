package processing

import (
	"fmt"
	"strings"
)

type fileNodeType string

const (
	FileNodeTypeFile      fileNodeType = "FILE"
	FileNodeTypeDirectory fileNodeType = "DIRECTORY"
)

type fileTreeNode interface {
	NodeType() fileNodeType
	FilePath() string
}

type fileNode struct {
	filePath string
	content  string
}

type directoryNode struct {
	filePath string
	children map[string]fileNode
}

func (f *fileNode) NodeType() fileNodeType {
	return FileNodeTypeFile
}

func (f *directoryNode) NodeType() fileNodeType {
	return FileNodeTypeDirectory
}

func (f *fileNode) FilePath() string {
	return f.filePath
}

func (f *directoryNode) FilePath() string {
	return f.filePath
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
