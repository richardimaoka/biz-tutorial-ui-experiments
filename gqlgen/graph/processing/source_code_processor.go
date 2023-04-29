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
	// 1. check if filePath is non-existent
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

func (p *SourceCodeProcessor) canApplyGitDiff(diff GitDiff) error {
	// pre-condition check, dupe in diff
	if diffDuplicate := diff.findDuplicate(); diffDuplicate.size() > 0 {
		return fmt.Errorf("duplicate file paths in diff = %+v", diffDuplicate)
	}

	// pre-condition check, each element's pre-condition check
	errors := []string{}
	for _, f := range diff.Added {
		if err := p.canAdd(f.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, f := range diff.Updated {
		if err := p.canDeleteOrUpdate(f.FilePath, fileType); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, f := range diff.Deleted {
		if err := p.canDeleteOrUpdate(f.FilePath, fileType); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}

func (p *SourceCodeProcessor) canApplyDirectoryDiff(diff DirectoryDiff) error {
	// pre-condition check, dupe in diff
	if diffDuplicate := diff.findDuplicate(); diffDuplicate.size() > 0 {
		return fmt.Errorf("duplicate file paths in diff = %+v", diffDuplicate)
	}

	// pre-condition check, each element's pre-condition check
	errors := []string{}
	for _, op := range diff.Added {
		if err := p.canAdd(op.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, op := range diff.Deleted {
		if err := p.canDeleteOrUpdate(op.FilePath, directoryType); err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("cannot apply diff, %s", strings.Join(errors, ", "))
	}

	return nil
}

func (p *SourceCodeProcessor) canApplyDiff(diff Diff) error {
	// 1. check if no dupe
	if diffDuplicate := diff.findDuplicate(); diffDuplicate.size() > 0 {
		return fmt.Errorf("duplicate file paths in diff = %+v", diffDuplicate)
	}

	// 2. each element's pre-condition check
	//    since there is no dupe, these checks are mutually exclusive
	errors := []string{}
	for _, op := range diff.FilesAdded {
		if err := p.canAdd(op.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, op := range diff.FilesDeleted {
		if err := p.canDeleteOrUpdate(op.FilePath, directoryType); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, op := range diff.FilesUpdated {
		if err := p.canAdd(op.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, op := range diff.DirectoriesAdded {
		if err := p.canDeleteOrUpdate(op.FilePath, directoryType); err != nil {
			errors = append(errors, err.Error())
		}
	}
	for _, op := range diff.DirectoriesDeleted {
		if err := p.canAdd(op.FilePath); err != nil {
			errors = append(errors, err.Error())
		}
	}

	// 3. report error if any
	if len(errors) != 0 {
		return fmt.Errorf("cannot apply diff, %s", strings.Join(errors, ", "))
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

func (p *SourceCodeProcessor) DeleteDirectory(op DirectoryDelete) error {
	if err := p.canDeleteOrUpdate(op.FilePath, directoryType); err != nil {
		return fmt.Errorf("cannot update file %s, %s", op.FilePath, err)
	}

	p.setAllIsUpdateFalse()
	p.deleteDirectoryMutation(op)
	return nil
}

func (p *SourceCodeProcessor) ApplyGitDiff(diff GitDiff) error {
	if err := p.canApplyGitDiff(diff); err != nil {
		return fmt.Errorf("cannot apply git diff, %s", err)
	}

	p.setAllIsUpdateFalse()
	for _, op := range diff.Added {
		p.addFileMutation(op)
	}
	for _, op := range diff.Updated {
		p.updateFileMutation(op)
	}
	for _, op := range diff.Deleted {
		p.deleteFileMutation(op)
	}

	return nil
}

func (p *SourceCodeProcessor) ApplyDirectoryDiff(diff DirectoryDiff) error {
	if err := p.canApplyDirectoryDiff(diff); err != nil {
		return fmt.Errorf("cannot apply git diff, %s", err)
	}

	p.setAllIsUpdateFalse()
	for _, op := range diff.Added {
		p.addDirectoryMutation(op)
	}
	for _, op := range diff.Deleted {
		p.deleteDirectoryMutation(op)
	}

	return nil
}

// func (p *SourceCodeProcessor) ApplyDiff(diff Diff) error {
// 	if err := p.canApplyDirectoryDiff(diff); err != nil {
// 		return fmt.Errorf("cannot apply git diff, %s", err)
// 	}

// 	p.setAllIsUpdateFalse()
// 	for _, op := range diff.Added {
// 		p.addDirectoryMutation(op)
// 	}
// 	for _, op := range diff.Deleted {
// 		p.deleteDirectoryMutation(op)
// 	}

// 	return nil
// }

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
