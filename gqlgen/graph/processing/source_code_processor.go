package processing

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCodeProcessor struct {
	repo                *git.Repository
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
			return fmt.Errorf("parent path = %s already exists as a file", strings.Join(parentPath, "/"))
		}
	}

	return nil
}

func (p *SourceCodeProcessor) confirmNoFileConflict(filePath string) error {
	// 1. check if filePath is non-existent
	exactMatchNode, exists := p.fileMap[filePath]
	if exists {
		return fmt.Errorf("path = %s already exists as a %s", exactMatchNode.FilePath(), exactMatchNode.NodeType())
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

func (p *SourceCodeProcessor) addDirectory(op DirectoryAdd) error {
	if err := p.canAdd(op.FilePath); err != nil {
		return fmt.Errorf("cannot add directory %s, %s", op.FilePath, err)
	}

	p.addDirectoryMutation(op)
	return nil
}

func (p *SourceCodeProcessor) addFile(op FileAdd) error {
	if err := p.canAdd(op.FilePath); err != nil {
		return fmt.Errorf("cannot add file %s, %s", op.FilePath, err)
	}

	p.addFileMutation(op)
	return nil
}

func (p *SourceCodeProcessor) deleteDirectory(op DirectoryDelete) error {
	if err := p.canDeleteOrUpdate(op.FilePath, directoryType); err != nil {
		return fmt.Errorf("cannot update file %s, %s", op.FilePath, err)
	}

	p.deleteDirectoryMutation(op)
	return nil
}

func (p *SourceCodeProcessor) updateFile(op FileUpdate) error {
	if err := p.canDeleteOrUpdate(op.FilePath, fileType); err != nil {
		return fmt.Errorf("cannot update file %s, %s", op.FilePath, err)
	}

	p.updateFileMutation(op)
	return nil
}

func (p *SourceCodeProcessor) deleteFile(op FileDelete) error {
	if err := p.canDeleteOrUpdate(op.FilePath, fileType); err != nil {
		return fmt.Errorf("cannot delete file %s, %s", op.FilePath, err)
	}

	p.deleteFileMutation(op)
	return nil
}

func (p *SourceCodeProcessor) upsertFile(op FileUpsert) error {
	canAddError := p.canAdd(op.FilePath)
	canUpdateError := p.canDeleteOrUpdate(op.FilePath, fileType)

	switch {
	case canAddError == nil:
		fileAddOp := FileAdd{FilePath: op.FilePath, Content: op.Content, IsFullContent: op.IsFullContent}
		p.addFileMutation(fileAddOp)
		return nil
	case canUpdateError == nil:
		fileUpdateOp := FileUpdate{FilePath: op.FilePath, Content: op.Content}
		file, found := p.fileMap[fileUpdateOp.FilePath]
		if found && !file.Matched(&fileProcessorNode{filePath: op.FilePath, isUpdated: true, content: op.Content}) {
			p.updateFileMutation(fileUpdateOp)
		}
		return nil
	default:
		return fmt.Errorf("cannot upsert file %s, %s", op.FilePath, canAddError)
	}
}

func (p *SourceCodeProcessor) applyFileOperation(operation FileOperation) error {
	switch v := operation.(type) {
	case DirectoryAdd:
		return p.addDirectory(v)
	case DirectoryDelete:
		return p.deleteDirectory(v)
	case FileAdd:
		return p.addFile(v)
	case FileUpdate:
		return p.updateFile(v)
	case FileDelete:
		return p.deleteFile(v)
	case FileUpsert:
		return p.upsertFile(v)
	default:
		return fmt.Errorf("wrong operation type = %v", reflect.TypeOf(v))
	}
}

func (p *SourceCodeProcessor) fileUpertOpsFromGit(commitHash string) ([]FileUpsert, error) {
	if p.repo == nil {
		return nil, fmt.Errorf("git repo is not initialized")
	}

	// repo := nil // git repo is needed at initialization
	commit, err := p.repo.CommitObject(plumbing.NewHash(commitHash))
	if err != nil {
		return nil, fmt.Errorf("error in git commit %s", commitHash)
	}

	var ops []FileUpsert
	fileIter, err := commit.Files()
	for {
		file, err := fileIter.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error in commit file traversal in commit %s", commitHash)
		}

		//TODO: check the size and use "***" as file.Contents() can panic if reading buffer grows too large
		contents, err := file.Contents()
		if err != nil {
			return nil, fmt.Errorf("error in reading file from commit %s", commitHash)
		}

		operation := FileUpsert{
			FilePath:      file.Name,
			Content:       contents,
			IsFullContent: true,
		}
		ops = append(ops, operation)
	}

	return ops, nil
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

func SourceCodeProcessorFromGit(repoUrl string) (*SourceCodeProcessor, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize source code, cannot clone repo %s, %s", repoUrl, err)
	}

	return &SourceCodeProcessor{
		repo:                repo,
		step:                "",
		defaultOpenFilePath: "",
		fileMap:             make(map[string]fileTreeNode),
	}, nil
}

func (p *SourceCodeProcessor) Transition(nextStep string, operation SourceCodeOperation) error {
	cloned := p.Clone()
	cloned.setAllIsUpdateFalse()

	switch v := operation.(type) {
	case SourceCodeFileOperation:
		for _, operation := range v.FileOps {
			if err := cloned.applyFileOperation(operation); err != nil {
				return fmt.Errorf("ApplyOperation failed, %s", err)
			}
		}
	case SourceCodeGitOperation:
		ops, err := p.fileUpertOpsFromGit(v.CommitHash)
		if err != nil {
			return fmt.Errorf("cannot transition to step %s, %s", nextStep, err)
		}

		for _, operation := range ops {
			if err := cloned.applyFileOperation(operation); err != nil {
				return fmt.Errorf("cannot transition to step %s, %s", nextStep, err)
			}
		}
	default:
		return fmt.Errorf("wrong operation type = %v", reflect.TypeOf(v))
	}

	p.defaultOpenFilePath = cloned.defaultOpenFilePath
	p.fileMap = cloned.fileMap
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
