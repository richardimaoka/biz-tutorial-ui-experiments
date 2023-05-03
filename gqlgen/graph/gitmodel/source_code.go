package gitmodel

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type SourceCodeFromGit struct {
}

func NewSourceCodeFromGit() *SourceCodeFromGit {
	return &SourceCodeFromGit{}
}

func (s *SourceCodeFromGit) Step() string {
	return ""
}

//method to return file node array
func (s *SourceCodeFromGit) FileNodes() []*model.FileNode {
	return []*model.FileNode{}
}

func (s *SourceCodeFromGit) OpenFile() *model.OpenFile {
	return nil
}
