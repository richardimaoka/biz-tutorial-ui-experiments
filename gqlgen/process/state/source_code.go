package state

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/memory"
)

type SourceCode struct {
	repo   *git.Repository
	commit plumbing.Hash
}

func NewSourceCode(repoUrl, currentCommitHash string) (*SourceCode, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCode, cannot clone repo %s, %s", repoUrl, err)
	}

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

	recursive(repo.Storer, currentRoot, 0)

	return nil, nil
}

// sub directories from tree
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

// sort directories

// sort files

func recursive(s storage.Storer, tree *object.Tree, offset int) error {
	whitespaces := strings.Repeat(" ", offset)

	var directories []*object.Tree
	var files []*object.Blob

	// needs to go through tree.Entries, as tree.Files() only returns files but not directories
	for _, e := range tree.Entries {
		obj, err := object.GetObject(s, e.Hash)
		if err != nil {
			return fmt.Errorf("failed in NewSourceCode, cannot get object = %s, %s", e.Hash.String(), err)
		}

		switch obj.Type() {

		case plumbing.TreeObject:
			fmt.Printf("%sdir  = %s\n", whitespaces, e.Name)
			tree, err := object.GetTree(s, e.Hash)
			if err != nil {
				return fmt.Errorf("failed in NewSourceCode, cannot get tree = %s, %s", e.Hash.String(), err)
			}
			directories = append(directories, tree)

		case plumbing.BlobObject:
			fmt.Printf("%sfile = %s\n", whitespaces, e.Name)
			file, err := object.GetBlob(s, e.Hash)
			if err != nil {
				return fmt.Errorf("failed in NewSourceCode, cannot get blob = %s, %s", e.Hash.String(), err)
			}
			files = append(files, file)
		}
	}

	//sort directories and store
	// for _, d := range directories {
	// 	recursive(s, d, offset+1)
	// }

	//sort files and store

	return nil
}
