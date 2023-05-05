package gitmodel

import (
	"fmt"
	"sort"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCodeFromGit struct {
	commitHash plumbing.Hash
	fileNodes  []*model.FileNode
}

func NewSourceCodeFromGit(repo *git.Repository, commitHash plumbing.Hash) (*SourceCodeFromGit, error) {
	commit, err := repo.CommitObject(commitHash)
	if err != nil {
		return nil, fmt.Errorf("error getting commit object: %v", err)
	}

	fileIter, err := commit.Files()
	if err != nil {
		return nil, fmt.Errorf("error getting file iterator: %v", err)
	}

	fileNodes := []*model.FileNode{}
	for {
		file, err := fileIter.Next()
		if err != nil {
			break
		}

		fileNode := NewFileFromGit(file.Name, false)
		fileNodes = append(fileNodes, fileNode.FileNode())
	}

	sort.Slice(fileNodes, func(i, j int) bool {
		iFilePath := fileNodes[i].FilePath
		jFilePath := fileNodes[j].FilePath
		return model.LessFilePath(*iFilePath, *jFilePath)
	})

	return &SourceCodeFromGit{
		commitHash: commitHash,
		fileNodes:  fileNodes,
	}, nil
}

func (s *SourceCodeFromGit) Step() string {
	return s.commitHash.String()
}

//method to return file node array
func (s *SourceCodeFromGit) FileNodes() []*model.FileNode {
	return []*model.FileNode{}
}

func (s *SourceCodeFromGit) OpenFile(filePath string) *model.OpenFile {
	return nil
}

func (s *SourceCodeFromGit) ToGraphQLSourceCode() *model.SourceCode {
	return nil
}
