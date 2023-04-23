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

	// 2.2 check intermediate nodes
	split := strings.Split(op.FilePath, "/")
	currentPath := []string{}
	for i := 0; i < len(split)-1; /*up to last-1 depth*/ i++ {
		childDir := split[i]
		currentPath = append(currentPath, childDir)

		childNode, exists := p.fileTree[strings.Join(currentPath, "/")]
		if exists {
			if childNode.NodeType() == fileType {
				return fmt.Errorf("cannot add directory %s, path = %s already exists as a file", op.FilePath, currentPath)
			}
			// else NodeType() == directoryType, which is ok
		} else {
			// 2.2.2 if child node doesn't exist, add an intermediate directory
		}
	}

	// 3. isUpdated to false
	for _, v := range p.fileTree {
		v.SetIsUpdated(false)
	}

	// 4. add directory at the last depth
	if childNode, exists := p.fileTree[op.FilePath]; exists {
		return fmt.Errorf("cannot add directory %s, path = %s already exists as %s", op.FilePath, op.FilePath, childNode.NodeType())
	}
	p.fileTree[op.FilePath] = &directoryProcessorNode{filePath: op.FilePath, isUpdated: true}

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
	resultNodes := []*model.FileNode{}

	for _, v := range p.fileTree {
		resultNodes = append(resultNodes, createDirectoryNode(v.FilePath(), v.IsUpdated()))
	}

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
