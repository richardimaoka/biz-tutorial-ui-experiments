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

	patch, err := prevCommit.Patch(currentCommit)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot get patch between prev commit = %s and current commit = %s, %s", prevCommitStr, currentCommitStr, err)
	}

	sc := SourceCode{repo: repo, commit: currentCommitHash}
	sc.recursive("", currentRoot, 0)

	for _, p := range patch.FilePatches() {
		from, to := p.Files()
		if from == nil {
			//added
			filePath := to.Path()
			//sc.addFile(filePath, to)
			for i, node := range sc.fileNodes {
				if node.FilePath() != filePath {
					continue
				}

				v, ok := node.(*File)
				if !ok {
					continue
				}
				v.isAdded = true
				v.isUpdated = true
				sc.fileNodes[i] = v
			}
		} else if to == nil {
			//sc.deleteFile(filePath, from)
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
		var dirPath string
		if currentDir == "" {
			dirPath = d.Name
		} else {
			dirPath = currentDir + "/" + d.Name
		}

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

		file, err := NewFileUnChanged(fileObj, currentDir)
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
