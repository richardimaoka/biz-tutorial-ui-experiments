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

func (p *SourceCodeProcessor) confirmNoFileConflict(filePath string) error {
	// 1. check if node with filePath is non-existent
	exactMatchNode, exists := p.fileMap[filePath]
	if exists {
		return fmt.Errorf("path = %s already exists as a %s", exactMatchNode, exactMatchNode.NodeType())
	}

	// 2. check if no intermediate parent is a file node
	split := strings.Split(filePath, "/")
	parentPath := []string{}
	for i := 0; i < len(split)-1; /*up to direct parent of filePath*/ i++ {
		parentPath = append(parentPath, split[i])
		parentNode, exists := p.fileMap[strings.Join(parentPath, "/")]
		if exists && parentNode.NodeType() == fileType {
			return fmt.Errorf("parent path = %s already exists as a file", parentPath)
		}
	}

	return nil
}

func (p *SourceCodeProcessor) isValidNode(filePath string, t nodeType) error {
	exactMatchNode, exists := p.fileMap[filePath]
	if !exists {
		return fmt.Errorf("path = %s doesn't exist", filePath)
	} else if exactMatchNode.NodeType() != t {
		return fmt.Errorf("path = %s has wrong node type = %s, expected = %s", filePath, exactMatchNode.NodeType(), t)
	}

	return nil
}

// this must be called after confirmNoFileConflict(), otherwise behavior is not guaranteed
func (p *SourceCodeProcessor) addMissingParentDirs(filePath string) {
	split := strings.Split(filePath, "/")
	parentPath := []string{}
	for i := 0; i < len(split)-1; /*up to direct parent of filePath*/ i++ {
		parentPath = append(parentPath, split[i])
		_, exists := p.fileMap[strings.Join(parentPath, "/")]
		if !exists {
			p.fileMap[strings.Join(parentPath, "/")] = &directoryProcessorNode{filePath: strings.Join(parentPath, "/"), isUpdated: true}
		}
	}
}

func (p *SourceCodeProcessor) setAllIsUpdateFalse() {
	for _, v := range p.fileMap {
		v.SetIsUpdated(false)
	}
}

//-----------------------------------------------------//
// public methods below
//-----------------------------------------------------//

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

	// 2. check if there no file conflicts
	if err := p.confirmNoFileConflict(op.FilePath); err != nil {
		return fmt.Errorf("cannot add directory %s, %s", op.FilePath, err)
	}

	// 3 mutation
	p.setAllIsUpdateFalse()

	p.addMissingParentDirs(op.FilePath)
	p.fileMap[op.FilePath] = &directoryProcessorNode{filePath: op.FilePath, isUpdated: true}

	return nil
}

func (p *SourceCodeProcessor) AddFile(op model.FileAdd) error {
	// 1. validate file path
	if err := isValidFilePath(op.FilePath); err != nil {
		return fmt.Errorf("cannot add directory %s, %s", op.FilePath, err)
	}

	// 2. check if there no file conflicts
	if err := p.confirmNoFileConflict(op.FilePath); err != nil {
		return fmt.Errorf("cannot add directory %s, %s", op.FilePath, err)
	}

	// 3 mutation
	p.setAllIsUpdateFalse()

	p.addMissingParentDirs(op.FilePath)
	p.fileMap[op.FilePath] = &fileProcessorNode{filePath: op.FilePath, isUpdated: true, content: op.Content}

	return nil
}

func (p *SourceCodeProcessor) UpdateFile(op model.FileUpdate) error {
	// 1. validate file path
	if err := isValidFilePath(op.FilePath); err != nil {
		return fmt.Errorf("cannot update file %s, %s", op.FilePath, err)
	}

	// 2. check if there is such a file
	if err := p.isValidNode(op.FilePath, fileType); err != nil {
		return fmt.Errorf("cannot update file %s, %s", op.FilePath, err)
	}

	// 3 mutation
	p.setAllIsUpdateFalse()

	p.fileMap[op.FilePath] = &fileProcessorNode{filePath: op.FilePath, isUpdated: true, content: op.Content}

	return nil
}

func (p *SourceCodeProcessor) DeleteFile(op model.FileDelete) error {
	// 1. validate file path
	if err := isValidFilePath(op.FilePath); err != nil {
		return fmt.Errorf("cannot delete file %s, %s", op.FilePath, err)
	}

	// 2. check if there is such a file
	if err := p.isValidNode(op.FilePath, fileType); err != nil {
		return fmt.Errorf("cannot delete file %s, %s", op.FilePath, err)
	}

	// 3 mutation
	p.setAllIsUpdateFalse()

	delete(p.fileMap, op.FilePath)

	return nil
}

func (p *SourceCodeProcessor) DeleteDirectory(op model.DirectoryDelete) error {
	// 1. validate file path
	if err := isValidFilePath(op.FilePath); err != nil {
		return fmt.Errorf("cannot delete directory %s, %s", op.FilePath, err)
	}

	// 2. check if there is such a directory
	if err := p.isValidNode(op.FilePath, directoryType); err != nil {
		return fmt.Errorf("cannot delete directory %s, %s", op.FilePath, err)
	}

	// 3 mutation
	p.setAllIsUpdateFalse()

	delete(p.fileMap, op.FilePath)
	// delete op.FilePath's children
	for k := range p.fileMap {
		if strings.HasPrefix(k, op.FilePath) {
			delete(p.fileMap, k)
		}
	}

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

func (p *SourceCodeProcessor) ToGraphQLModel() *model.SourceCode {
	var resultNodes []*model.FileNode
	fileContents := make(map[string]model.OpenFile)

	keys := p.sortedFileMapKeys()
	for i := 0; i < len(keys); i++ {
		node := p.fileMap[keys[i]]
		resultNodes = append(resultNodes, node.ToGraphQLNode())

		if v, ok := node.(*fileProcessorNode); ok {
			fileContents[node.FilePath()] = *v.ToGraphQLOpenFile()
		}
	}

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
