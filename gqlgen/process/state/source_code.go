package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCode struct {
	repo       *git.Repository
	commit     plumbing.Hash
	rootDir    *Directory
	fileNodes  []FileNode
	tutorial   string
	projectDir string
	step       string

	// can be set from caller
	DefaultOpenFilePath string
	IsFoldFileTree      bool
}

func getCommit(repo *git.Repository, hashStr string) (*object.Commit, error) {
	commitHash := plumbing.NewHash(hashStr)
	if commitHash.String() != hashStr {
		return nil, fmt.Errorf("commit hash = %s mismatched with re-calculated hash = %s", hashStr, commitHash.String())
	}

	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get commit = %s, %s", hashStr, err)
	}

	return commit, nil
}

func NewSourceCode2(repo *git.Repository, projectDir, tutorial string) *SourceCode {
	return &SourceCode{
		repo:           repo,
		projectDir:     projectDir,
		tutorial:       tutorial,
		IsFoldFileTree: true,
	}
}

// stateless, predictable from arguments
func (s *SourceCode) ConstructStep(step, prevCommitHash, currentCommitHash string) error {
	// 1. Construct source code root dir as if all files are unchanged
	currentCommit, err := getCommit(s.repo, currentCommitHash)
	if err != nil {
		return fmt.Errorf("failed in ConstructStep, cannot get current commit = %s, %v", currentCommitHash, err)
	}

	currentRoot, err := currentCommit.Tree()
	if err != nil {
		return fmt.Errorf("failed in ConstructStep, cannot get the root tree for commit = %s, %s", currentCommitHash, err)
	}

	rootDir, err := ConstructDirectory(s.repo, "", currentRoot, false)
	if err != nil {
		return fmt.Errorf("failed in ConstructStep, cannot create root directory, %s", err)
	}

	// 2. From git patches, mark added/deleted/updated/renamed files
	prevCommit, err := getCommit(s.repo, prevCommitHash)
	if err != nil {
		return fmt.Errorf("failed in ConstructStep, cannot get prev commit, %s", err)
	}

	patch, err := prevCommit.Patch(currentCommit)
	if err != nil {
		return fmt.Errorf("failed in ConstructStep, cannot get patch between prev commit = %s and current commit = %s, %s", prevCommitHash, currentCommitHash, err)
	}

	for _, p := range patch.FilePatches() {
		from, to := p.Files() // See Files() method's comment about when 'from' and 'to' become nil
		if from == nil {
			//added
			rootDir.MarkFileAdded(to.Path())
		} else if to == nil {
			// deleted
			rootDir.InsertFileDeleted("", from.Path(), from)
		} else if from.Path() != to.Path() {
			// renamed
		} else {
			// updated
			rootDir.MarkFileUpdated(from.Path(), from, p)
		}
	}

	s.step = step
	s.rootDir = rootDir

	return nil
}

// stateful, current state + arguments => next state
func (s *SourceCode) ForwardCommit(nextStep, commitStr string) error {
	return nil
}

// stateful, current state + arguments => next state
// func (s *SourceCode) ForwardStepOps(nextStep string, fileOps []string, defaultOpenFilePath string, isFoldFileTree bool) error {
// 	return nil
// }

// SourceCodeFromChange(prevCommit, currentCommit)
func NewSourceCode(repo *git.Repository, currentCommitStr, prevCommitStr, tutorial, step, defaultOpenFilePath string, isFoldFileTree bool) (*SourceCode, error) {
	// 1. Construct source code root dir as if all files are unchanged
	currentCommit, err := getCommit(repo, currentCommitStr)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get current commit, %s", err)
	}

	currentRoot, err := currentCommit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get the root tree for commit = %s, %s", currentCommitStr, err)
	}

	rootDir, err := ConstructDirectory(repo, "", currentRoot, false)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot create root directory, %s", err)
	}

	sc := SourceCode{repo: repo, commit: plumbing.NewHash(currentCommitStr), rootDir: rootDir, tutorial: tutorial, step: step, DefaultOpenFilePath: defaultOpenFilePath, IsFoldFileTree: isFoldFileTree}

	// 2. From git patches, mark added/deleted/updated/renamed files
	prevCommit, err := getCommit(repo, prevCommitStr)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get prev commit, %s", err)
	}

	patch, err := prevCommit.Patch(currentCommit)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get patch between prev commit = %s and current commit = %s, %s", prevCommitStr, currentCommitStr, err)
	}

	for _, p := range patch.FilePatches() {
		from, to := p.Files() // See Files() method's comment about when 'from' and 'to' become nil
		if from == nil {
			//added
			sc.rootDir.MarkFileAdded(to.Path())
		} else if to == nil {
			// deleted
			sc.rootDir.InsertFileDeleted("", from.Path(), from)
		} else if from.Path() != to.Path() {
			// renamed
		} else {
			// updated
			sc.rootDir.MarkFileUpdated(from.Path(), from, p)
		}
	}

	return &sc, nil
}

// SourceCodeFromCommit(commit) // all files are FileUnchanged
func InitialSourceCode(repo *git.Repository, currentCommitStr, step, defaultOpenFilePath, tutorial string, isFoldFileTree bool) (*SourceCode, error) {
	// 1. Construct source code root dir as if all files are unchanged
	currentCommit, err := getCommit(repo, currentCommitStr)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get current commit, %s", err)
	}

	currentRoot, err := currentCommit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get the root tree for commit = %s, %s", currentCommitStr, err)
	}

	rootDir, err := ConstructDirectory(repo, "", currentRoot, true)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot create root directory, %s", err)
	}

	sc := SourceCode{repo: repo, commit: plumbing.NewHash(currentCommitStr), rootDir: rootDir, tutorial: tutorial, step: step, DefaultOpenFilePath: defaultOpenFilePath, IsFoldFileTree: isFoldFileTree}

	return &sc, nil
}

func (s *SourceCode) ToGraphQLSourceCode() *model.SourceCode {

	return &model.SourceCode{
		FileTree:            s.rootDir.ToGraphQLFileNodeSlice(),
		FileContents:        s.rootDir.ToGraphQLOpenFileMap(),
		Tutorial:            s.tutorial,
		Step:                s.step,
		DefaultOpenFilePath: s.DefaultOpenFilePath,
		IsFoldFileTree:      s.IsFoldFileTree,
		ProjectDir:          s.projectDir,
	}
}
