package gitmodel

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCodeFromGit struct {
	commitHash plumbing.Hash
	commit     object.Commit
}

func NewSourceCodeFromGit(repoUrl string, commitHash plumbing.Hash) *SourceCodeFromGit {
	return &SourceCodeFromGit{}
}

func (s *SourceCodeFromGit) setFiles() int64 {
	//sort and set files
	//files, err := s.commit.Files()
	return 0
}

func (s *SourceCodeFromGit) Step() string {
	return ""
}

//method to return file node array
func (s *SourceCodeFromGit) FileNodes() []*model.FileNode {
	return []*model.FileNode{}
}

func (s *SourceCodeFromGit) OpenFile(filePath string) *model.OpenFile {
	return nil
}
