package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCode struct {
	// metadata, can be set only at initialization
	tutorial   string
	projectDir string
	repo       *git.Repository

	// inner state updated at each step
	commitHash string
	rootDir    *Directory
	step       string

	// metadata, can be set from caller anytime
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

// pure function, to make it clear it's side-effect free
func dirBetweenCommits(repo *git.Repository, prevCommit, currentCommit *object.Commit) (*Directory, error) {
	currentRoot, err := currentCommit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in dirFromTwoCommits, cannot get the root tree for commit = %s, %s", currentCommit.Hash, err)
	}

	rootDir, err := ConstructDirectory(repo, "", currentRoot, false) //false as FileUnchanged
	if err != nil {
		return nil, fmt.Errorf("failed in dirFromTwoCommits, cannot create root directory, %s", err)
	}

	patch, err := prevCommit.Patch(currentCommit)
	if err != nil {
		return nil, fmt.Errorf("failed in dirFromTwoCommits, cannot get patch between prev commit = %s and current commit = %s, %s", prevCommit.Hash, currentCommit.Hash, err)
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

	return rootDir, nil
}

// stateful, current state + arguments => next state
func (s *SourceCode) ForwardCommit(step, currentCommitStr string) error {
	// 1.1. check if it's the initial commit, before setting s.rootDir
	isInitialCommit := s.rootDir == nil

	// 1.2. variables used commonly in both of if/else blocks
	currentCommit, err := getCommit(s.repo, currentCommitStr)
	if err != nil {
		return fmt.Errorf("failed in ForwardCommit, cannot get current commit, %s", err)
	}

	var rootDir *Directory
	if isInitialCommit {
		// 2.1 All files as FileAdded for the initial commit
		currentRoot, err := currentCommit.Tree()
		if err != nil {
			return fmt.Errorf("failed in ForwardCommit, cannot get the root tree for commit = %s, %s", currentCommitStr, err)
		}

		rootDir, err = ConstructDirectory(s.repo, "", currentRoot, true) //true, as FileAdded
		if err != nil {
			return fmt.Errorf("failed in ForwardCommit, cannot create root directory, %s", err)
		}
	} else {
		// 2.2 Mark files in diff from prevCommit
		prevCommitHash := s.commitHash // s.commitHash preserves the prev commit at this point
		prevCommit, err := getCommit(s.repo, prevCommitHash)
		if err != nil {
			return fmt.Errorf("failed in ForwardCommit, cannot get prev commit = %s, %s", prevCommitHash, err)
		}

		rootDir, err = dirBetweenCommits(s.repo, prevCommit, currentCommit)
		if err != nil {
			return fmt.Errorf("failed in ForwardCommit, cannot create root directory, %s", err)
		}
	}

	s.rootDir = rootDir
	s.step = step
	s.commitHash = currentCommitStr

	return nil
}

// stateful, reset the inner state to the given commit
func (s *SourceCode) ResetCommit(step, commitStr string) error {
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

	sc := SourceCode{repo: repo, commitHash: currentCommitStr, rootDir: rootDir, tutorial: tutorial, step: step, DefaultOpenFilePath: defaultOpenFilePath, IsFoldFileTree: isFoldFileTree}

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

	sc := SourceCode{repo: repo, commitHash: currentCommitStr, rootDir: rootDir, tutorial: tutorial, step: step, DefaultOpenFilePath: defaultOpenFilePath, IsFoldFileTree: isFoldFileTree}

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
