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

func (p *SourceCodeProcessor) confirmNoParentIsFile(filePath string) error {
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

func (p *SourceCodeProcessor) confirmNoFileConflict(filePath string) error {
	// 1. check if filePath is non-existent
	exactMatchNode, exists := p.fileMap[filePath]
	if exists {
		return fmt.Errorf("path = %s already exists as a %s", exactMatchNode, exactMatchNode.NodeType())
	}

	// 2. check if no intermediate parent is a file node
	if err := p.confirmNoParentIsFile(filePath); err != nil {
		return err
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

func (p *SourceCodeProcessor) canAdd(filePath string) error {
	// 1. validate file path
	if err := isValidFilePath(filePath); err != nil {
		return err
	}

	// 2. check if there no file conflicts
	if err := p.confirmNoFileConflict(filePath); err != nil {
		return err
	}

	return nil
}

func (p *SourceCodeProcessor) canDeleteOrUpdate(filePath string, t nodeType) error {
	// 1. validate file path
	if err := isValidFilePath(filePath); err != nil {
		return err
	}

	// 2. check if there is such a file
	if err := p.isValidNode(filePath, t); err != nil {
		return err
	}

	return nil
}

func (p *SourceCodeProcessor) addDirectoryMutation(op DirectoryAdd) {
	p.addMissingParentDirs(op.FilePath)
	p.fileMap[op.FilePath] = &directoryProcessorNode{filePath: op.FilePath, isUpdated: true}
}

func (p *SourceCodeProcessor) addFileMutation(op FileAdd) {
	p.addMissingParentDirs(op.FilePath)
	p.fileMap[op.FilePath] = &fileProcessorNode{filePath: op.FilePath, isUpdated: true, content: op.Content}
}

func (p *SourceCodeProcessor) updateFileMutation(op FileUpdate) {
	p.fileMap[op.FilePath] = &fileProcessorNode{filePath: op.FilePath, isUpdated: true, content: op.Content}
}

func (p *SourceCodeProcessor) deleteFileMutation(op FileDelete) {
	delete(p.fileMap, op.FilePath)
}

func (p *SourceCodeProcessor) deleteDirectoryMutation(op DirectoryDelete) {
	delete(p.fileMap, op.FilePath)
	// delete op.FilePath's children
	for k := range p.fileMap {
		if strings.HasPrefix(k, op.FilePath) {
			delete(p.fileMap, k)
		}
	}
}

func (p *SourceCodeProcessor) applyDiifMutation(diff Diff) error {
	// does the order of operations have any implication??
	for _, op := range diff.DirectoriesDeleted {
		if err := p.DeleteDirectory(op); err != nil {
			return err
		}
	}
	for _, op := range diff.DirectoriesAdded {
		if err := p.AddDirectory(op); err != nil {
			return err
		}
	}
	for _, op := range diff.FilesDeleted {
		if err := p.DeleteFile(op); err != nil {
			return err
		}
	}
	for _, op := range diff.FilesAdded {
		if err := p.AddFile(op); err != nil {
			return err
		}
	}
	for _, op := range diff.FilesUpdated {
		if err := p.UpdateFile(op); err != nil {
			return err
		}
	}

	return nil
}

//-----------------------------------------------------//
// public methods below
//-----------------------------------------------------//

func NewSourceCodeProcessor() *SourceCodeProcessor {
	return &SourceCodeProcessor{
		step:                "",
		defaultOpenFilePath: "",
		fileMap:             make(map[string]fileTreeNode),
	}
}

func (p *SourceCodeProcessor) AddDirectory(op DirectoryAdd) error {
	if err := p.canAdd(op.FilePath); err != nil {
		return fmt.Errorf("cannot add directory %s, %s", op.FilePath, err)
	}

	p.setAllIsUpdateFalse()
	p.addDirectoryMutation(op)
	return nil
}

func (p *SourceCodeProcessor) AddFile(op FileAdd) error {
	if err := p.canAdd(op.FilePath); err != nil {
		return fmt.Errorf("cannot add file %s, %s", op.FilePath, err)
	}

	p.setAllIsUpdateFalse()
	p.addFileMutation(op)
	return nil
}

func (p *SourceCodeProcessor) UpdateFile(op FileUpdate) error {
	if err := p.canDeleteOrUpdate(op.FilePath, fileType); err != nil {
		return fmt.Errorf("cannot update file %s, %s", op.FilePath, err)
	}

	p.setAllIsUpdateFalse()
	p.updateFileMutation(op)
	return nil
}

func (p *SourceCodeProcessor) DeleteFile(op FileDelete) error {
	if err := p.canDeleteOrUpdate(op.FilePath, fileType); err != nil {
		return fmt.Errorf("cannot delete file %s, %s", op.FilePath, err)
	}

	p.setAllIsUpdateFalse()
	p.deleteFileMutation(op)
	return nil
}

func (p *SourceCodeProcessor) UpsertFile(op FileUpsert) error {
	fileAddOp := FileAdd{FilePath: op.FilePath, Content: op.Content, IsFullContent: op.IsFullContent}
	if errAdd := p.AddFile(fileAddOp); errAdd != nil {
		fileUpdateOp := FileUpdate{FilePath: op.FilePath, Content: op.Content}
		if errUpd := p.UpdateFile(fileUpdateOp); errUpd != nil {
			return fmt.Errorf("cannot upsert file %s, %s, %s", op.FilePath, errAdd, errUpd)
		}
	}

	return nil
}

func (p *SourceCodeProcessor) DeleteDirectory(op DirectoryDelete) error {
	if err := p.canDeleteOrUpdate(op.FilePath, directoryType); err != nil {
		return fmt.Errorf("cannot update file %s, %s", op.FilePath, err)
	}

	p.setAllIsUpdateFalse()
	p.deleteDirectoryMutation(op)
	return nil
}

func (p *SourceCodeProcessor) ApplyDiff(diff Diff) error {
	cloned := p.Clone()
	cloned.setAllIsUpdateFalse()
	if err := cloned.applyDiifMutation(diff); err != nil {
		return fmt.Errorf("cannot apply diff, %s", err)
	}

	p.step = cloned.step
	p.defaultOpenFilePath = cloned.defaultOpenFilePath
	p.fileMap = cloned.fileMap

	return nil
}

func (p *SourceCodeProcessor) SetStep(step string) {
	p.step = step
}

func (p *SourceCodeProcessor) Transition(nextStep string, effect SourceCodeEffect) error {
	if err := p.ApplyDiff(effect.Diff); err != nil {
		return fmt.Errorf("cannot transition to step %s, %s", nextStep, err)
	}
	p.step = nextStep
	return nil
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
		Step:         p.step,
		FileTree:     resultNodes,
		FileContents: fileContents,
	}
}

func (p *SourceCodeProcessor) Clone() *SourceCodeProcessor {
	// clone to avoid receiver's mutation effect afterwards
	fileMap := make(map[string]fileTreeNode)
	for k := range p.fileMap {
		fileMap[k] = p.fileMap[k].Clone()
	}

	return &SourceCodeProcessor{
		step:                p.step,
		defaultOpenFilePath: p.defaultOpenFilePath,
		fileMap:             fileMap,
	}
}
