package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

type SourceCode struct {
	repo      *git.Repository
	commit    plumbing.Hash
	rootDir   *Directory
	fileNodes []FileNode
}

func NewSourceCode(repo *git.Repository, currentCommitStr string, prevCommitStr string) (*SourceCode, error) {
	currentCommitHash := plumbing.NewHash(currentCommitStr)
	if currentCommitHash.String() != currentCommitStr {
		return nil, fmt.Errorf("failed in NewSourceCode, current commit hash = %s is invalid as its re-calculated hash is mismatched = %s", currentCommitStr, currentCommitHash.String())
	}
	prevCommitHash := plumbing.NewHash(prevCommitStr)
	if prevCommitHash.String() != prevCommitStr {
		return nil, fmt.Errorf("failed in NewSourceCode, prev commit hash = %s is invalid as its re-calculated hash is mismatched = %s", prevCommitStr, prevCommitHash.String())
	}

	currentCommit, err := repo.CommitObject(currentCommitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get current commit = %s, %s", currentCommitStr, err)
	}
	prevCommit, err := repo.CommitObject(prevCommitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get prev commit = %s, %s", prevCommitStr, err)
	}

	currentRoot, err := currentCommit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get the root tree for commit = %s, %s", currentCommitStr, err)
	}

	rootDir, err := state.ConstructDirectory(repo, "", currentRoot)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot create root directory, %s", err)
	}

	patch, err := prevCommit.Patch(currentCommit)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get patch between prev commit = %s and current commit = %s, %s", prevCommitStr, currentCommitStr, err)
	}

	sc := SourceCode{repo: repo, commit: currentCommitHash, rootDir: rootDir}

	for _, p := range patch.FilePatches() {
		from, to := p.Files()
		if from == nil {
			//added
			sc.rootDir.MarkFileAdded(to.Path())
		} else if to == nil {
			sc.rootDir.InsertFileDeleted("", from.Path(), from)
			// deleted
		} else if from.Path() != to.Path() {
			//sc.renameFile(filePath, from)
			// renamed
		} else {
			// updated
		}
	}

	return &sc, nil
}

func (s *SourceCode) ToGraphQLSourceCode() *model.SourceCode {
	return &model.SourceCode{
		FileTree: s.rootDir.ToGraphQLFileNodeSlice(),
	}
}
