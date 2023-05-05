package gitmodel

import (
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type FileFromGit struct {
	filePath  string
	isUpdated bool
}

func NewFileFromGit(filePath string, isUpdated bool) *FileFromGit {

	return &FileFromGit{
		filePath:  filePath,
		isUpdated: isUpdated}
}

func (f *FileFromGit) offset() int {
	split := strings.Split(f.filePath, "/")
	return len(split) - 1
}

func (f *FileFromGit) name() string {
	split := strings.Split(f.filePath, "/")
	return split[len(split)-1]
}

func (f *FileFromGit) FileNode() *model.FileNode {
	name := f.name()
	nodeType := model.FileNodeTypeFile
	offset := f.offset()

	return &model.FileNode{
		FilePath:  &f.filePath,  //pointer is safe here, as f.filePath is effectively immutable
		IsUpdated: &f.isUpdated, //pointer is safe here, as f.filePath is effectively immutable
		Name:      &name,
		NodeType:  &nodeType,
		Offset:    &offset,
	}
}

func (f *FileFromGit) OpenFile() *model.OpenFile {
	return nil
}
