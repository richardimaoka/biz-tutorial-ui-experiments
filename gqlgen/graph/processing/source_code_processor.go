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

	currentPath := []string{}
	split := strings.Split(op.FilePath, "/")

	childDir := split[0]
	currentPath = append(currentPath, childDir)

	_, exists := p.fileTree[childDir]
	if exists {
		errorFilePath := strings.Join(currentPath, "/")
		return fmt.Errorf("cannot add directory %s, path = %s already exists", op.FilePath, errorFilePath)
	}

	p.fileTree[op.FilePath] = &directoryProcessorNode{filePath: op.FilePath, children: make(map[string]fileTreeNode)}

	// // 2. depth search
	// currentTree := p.fileTree
	// currentPath := []string{}
	// split := strings.Split(op.FilePath, "/")
	// for i := 0; i < len(split)-1; /*exclude last element*/ i++ {
	// 	childDir := split[i]
	// 	currentPath = append(currentPath, childDir)

	// 	// 2.1 if child node already exists, then error
	// 	childNode, exists := currentTree[childDir]
	// 	if exists {
	// 		errorFilePath := strings.Join(currentPath, "/")
	// 		return fmt.Errorf("cannot add directory %s, path = %s already exists", op.FilePath, errorFilePath)
	// 	}

	// 	//

	// 	if childNode.NodeType() == FileNodeTypeFile {
	// 		return fmt.Errorf("cannot add directory %s, path = %s already exists as a file", op.FilePath, dir)
	// 	}

	// }

	// last := split[len(split)-1]
	// currentTree[last] = &directoryNode{filePath: op.FilePath, children: make(map[string]fileNode)}
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

func (p *SourceCodeProcessor) ToSourceCode() *model.SourceCode {
	fileTree := []*model.FileNode{}
	for _, v := range p.fileTree {
		fileTree = append(fileTree, createDirectoryNode(v.FilePath()))
	}

	fileContents := make(map[string]model.OpenFile)
	return &model.SourceCode{
		Step:         "",
		FileTree:     fileTree,
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
