package processing

import (
	"fmt"
	"sort"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCodeProcessor struct {
	step                string
	defaultOpenFilePath string
	fileMap             map[string]fileTreeNode
}

func NewSourceCodeProcessor() *SourceCodeProcessor {
	return &SourceCodeProcessor{
		step:                "init",
		defaultOpenFilePath: "",
		fileMap:             make(map[string]fileTreeNode),
	}
}

func (p *SourceCodeProcessor) AddDirectory(op model.DirectoryAdd) error {
	// 1. validate file path
	if err := isValidFilePath(op.FilePath); err != nil {
		return fmt.Errorf("cannot add directory %s, %s", op.FilePath, err)
	}

	// 2.1. check if there is no intermediate file node in the path
	split := strings.Split(op.FilePath, "/")
	currentPath := []string{}
	for i := 0; i < len(split)-1; /*up to last-1 depth*/ i++ {
		currentPath = append(currentPath, split[i])
		currentNode, exists := p.fileMap[strings.Join(currentPath, "/")]
		if exists && currentNode.NodeType() == fileType {
			return fmt.Errorf("cannot add directory %s, path = %s already exists as a file", op.FilePath, currentPath)
		}
	}

	// 2.2 check if the last node is non-existent
	lastNode, exists := p.fileMap[op.FilePath]
	if exists {
		return fmt.Errorf("cannot add directory %s, path = %s already exists as a %s", op.FilePath, currentPath, lastNode.NodeType())
	}

	// 3.1 mutation: isUpdated to false
	for _, v := range p.fileMap {
		v.SetIsUpdated(false)
	}

	// 3.2 mutation: add intermediate and last directories
	currentPath = []string{}
	for i := 0; i < len(split); i++ {
		currentPath = append(currentPath, split[i])
		_, exists := p.fileMap[strings.Join(currentPath, "/")]
		if !exists {
			p.fileMap[strings.Join(currentPath, "/")] = &directoryProcessorNode{filePath: strings.Join(currentPath, "/"), isUpdated: true}
		}
	}

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

func (p *SourceCodeProcessor) sortedFileMapKeys() []string {
	keys := make([]string, 0)
	for k := range p.fileMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return LessFilePath(keys[i], keys[j])
	})

	return keys
}

func (p *SourceCodeProcessor) ToSourceCode() *model.SourceCode {
	resultNodes := []*model.FileNode{}
	keys := p.sortedFileMapKeys()
	for i := 0; i < len(keys); i++ {
		node := p.fileMap[keys[i]]
		resultNodes = append(resultNodes, createDirectoryNode(node.FilePath(), node.IsUpdated()))
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
