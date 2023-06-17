package processing

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage"
)

type SourceCodeOperation interface {
	IsSourceCodeOperation()
}

type SourceCodeFileOperation struct {
	FileOps []FileOperation
}

type SourceCodeGitOperation struct {
	CommitHash string
}

func (o SourceCodeFileOperation) IsSourceCodeOperation() {}
func (o SourceCodeGitOperation) IsSourceCodeOperation()  {}

func FileOpsFromCommit(repo *git.Repository, currentCommit, prevCommit *object.Commit) ([]FileOperation, error) {
	patch, err := prevCommit.Patch(currentCommit)
	if err != nil {
		return nil, fmt.Errorf("FileOpsFromCommit failed to get patch, %s", err)
	}

	getContents := func(storer storage.Storer, file diff.File) (string, error) {
		blob, err := object.GetBlob(repo.Storer, file.Hash())
		if err != nil {
			return "", err
		}

		fileObj := object.NewFile(file.Path(), file.Mode(), blob)
		contents, err := fileObj.Contents()
		if err != nil {
			return "", err
		}

		return contents, nil
	}

	var ops []FileOperation
	for _, v := range patch.FilePatches() {
		from, to := v.Files()
		if from == nil {
			toContents, err := getContents(repo.Storer, to)
			if err != nil {
				return nil, fmt.Errorf("FileOpsFromCommit failed to get contents, %s", err)
			}
			fileAdd := FileAdd{FilePath: to.Path(), Content: toContents, IsFullContent: true}
			ops = append(ops, fileAdd)
		} else if to == nil {
			fileDelete := FileDelete{FilePath: from.Path()}
			ops = append(ops, fileDelete)
		} else {
			toContents, err := getContents(repo.Storer, to)
			if err != nil {
				return nil, fmt.Errorf("FileOpsFromCommit failed to get contents, %s", err)
			}
			fileUpdate := FileUpdate{FilePath: to.Path(), Content: toContents}
			ops = append(ops, fileUpdate)
		}
	}

	return ops, nil
}
