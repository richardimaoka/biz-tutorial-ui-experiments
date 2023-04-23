package processing

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCodeProcessor struct {
	step                string
	defaultOpenFilePath string
	fileTree            map[string]fileTreeNode
}

func NewSourceCodeProcessor() *SourceCodeProcessor {
	return &SourceCodeProcessor{
		step:                "init",
		defaultOpenFilePath: "",
		fileTree:            make(map[string]fileTreeNode),
	}
}

func (p *SourceCodeProcessor) AddDirectory(op model.DirectoryAdd) error {
	// 1. validate file path
	if err := isValidFilePath(op.FilePath); err != nil {
		return fmt.Errorf("cannot add directory %s, %s", op.FilePath, err)
	}

	// 2.1 initialize currentTree to root
	currentTree := p.fileTree
	currentPath := []string{}

	// 2.2 traverse p.fileTree up to last-1 depth
	split := strings.Split(op.FilePath, "/")
	for i := 0; i < len(split)-1; /*up to last-1 depth*/ i++ {
		childDir := split[i]
		currentPath = append(currentPath, childDir)

		childNode, exists := currentTree[childDir]
		if exists {
			// 2.2.1 if child node already exists, it should be directory
			switch v := childNode.(type) {
			case *directoryProcessorNode:
				currentTree = v.children
			default:
				return fmt.Errorf("cannot add directory %s, path = %s already exists as a file", op.FilePath, currentPath)
			}
		} else {
			// 2.2.2 if child node doesn't exist, add an intermediate directory
		}
	}

	// 3. add directory at the last depth
	childDir := split[len(split)-1]
	currentPath = append(currentPath, childDir)
	childNode, exists := currentTree[childDir]
	if exists {
		switch childNode.(type) {
		case *directoryProcessorNode:
			return fmt.Errorf("cannot add directory %s, path = %s already exists as a directory", op.FilePath, currentPath)
		default:
			return fmt.Errorf("cannot add directory %s, path = %s already exists as a file", op.FilePath, currentPath)
		}
	}

	currentTree[childDir] = &directoryProcessorNode{filePath: op.FilePath, children: make(map[string]fileTreeNode)}

	return nil
}

func (p *SourceCodeProcessor) AddFile(op model.FileAdd) error {
	return nil
}

func (p *SourceCodeProcessor) UpdateFile(op model.FileUpdate) error {
	return nil
}

func (p *SourceCodeProcessor) DeleteFile(op model.FileDelete) error {
	return nil
}

func (p *SourceCodeProcessor) DeleteDirectory(op model.DirectoryDelete) error {
	return nil
}

func recurse(resultNodes []*model.FileNode, currentTree map[string]fileTreeNode) []*model.FileNode {
	for _, v := range currentTree {
		switch n := v.(type) {
		case *directoryProcessorNode:
			resultNodes = append(resultNodes, createDirectoryNode(n.FilePath()))
			resultNodes = recurse(resultNodes, n.children)
		case *fileProcessorNode:
			resultNodes = append(resultNodes, createFileNode(v.FilePath()))
		}
	}

	return resultNodes
}

func (p *SourceCodeProcessor) ToSourceCode() *model.SourceCode {
	resultNodes := []*model.FileNode{}
	resultNodes = recurse(resultNodes, p.fileTree)

	fileContents := make(map[string]model.OpenFile)
	return &model.SourceCode{
		Step:         "",
		FileTree:     resultNodes,
		FileContents: fileContents,
	}
}

/*
	file node with content
	dir node with content

	find file
	find directory
	validate file/directory node (find node)

	can add file
	  - canAddFileNode
	  - canAddFileContent
	can update file
	can delete file

	can add directory
	can delete directory

	add/udpdate/delete file
	add/delete directory

	apply diff

*/
