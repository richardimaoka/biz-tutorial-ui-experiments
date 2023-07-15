package state

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCode struct {
	repo      *git.Repository
	commit    plumbing.Hash
	fileNodes []FileNode
}

func NewSourceCode(repo *git.Repository, currentCommitHash string) (*SourceCode, error) {
	commitHash := plumbing.NewHash(currentCommitHash)
	if commitHash.String() != currentCommitHash {
		return nil, fmt.Errorf("failed in NewSourceCode, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", currentCommitHash, commitHash.String())
	}

	currentCommit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get commit = %s, %s", currentCommitHash, err)
	}

	currentRoot, err := currentCommit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get the root tree for commit = %s, %s", currentCommitHash, err)
	}

	sc := SourceCode{repo: repo, commit: commitHash}
	sc.recursive("", currentRoot, 0)

	return &sc, nil
}

func TreeFilesDirs(tree *object.Tree) ([]object.TreeEntry, []object.TreeEntry) {
	var files []object.TreeEntry
	var dirs []object.TreeEntry

	for _, e := range tree.Entries {
		if e.Mode.IsFile() {
			files = append(files, e)
		} else {
			dirs = append(dirs, e)
		}
	}
	return files, dirs
}

func SortEntries(entries []object.TreeEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name) < strings.ToLower(entries[j].Name)
	})
}

func (s *SourceCode) recursive(currentDir string, tree *object.Tree, offset int) error {
	files, dirs := TreeFilesDirs(tree)
	SortEntries(files)
	SortEntries(dirs)

	for _, d := range dirs {
		dirPath := currentDir + "/" + d.Name
		subTree, err := object.GetTree(s.repo.Storer, d.Hash)
		if err != nil {
			return fmt.Errorf("failed in recursive, cannot get tree = %s, %s", d.Name, err)
		}

		dir, err := NewDirectory(dirPath, nil, subTree)
		if err != nil {
			return fmt.Errorf("failed in recursive, cannot create directory = %s, %s", dirPath, err)
		}
		s.fileNodes = append(s.fileNodes, dir)

		s.recursive(dirPath, subTree, offset+1)
	}

	for _, f := range files {
		fileObj, err := tree.File(f.Name)
		if err != nil {
			return fmt.Errorf("failed in recursive, cannot get file = %s in dir = %s, %s", f.Name, currentDir, err)
		}

		file, err := NewFile(nil, fileObj, currentDir)
		if err != nil {
			return fmt.Errorf("failed in recursive, cannot create file = %s in dir = %s, %s", f.Name, currentDir, err)
		}

		s.fileNodes = append(s.fileNodes, file)
	}

	return nil
}

func (p *SourceCode) ToGraphQLSourceCode() *model.SourceCode {
	var resultNodes []*model.FileNode

	for _, node := range p.fileNodes {
		resultNodes = append(resultNodes, node.ToGraphQLFileNode())
	}

	return &model.SourceCode{
		FileTree: resultNodes,
	}
}
