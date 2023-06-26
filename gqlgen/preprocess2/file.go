package preprocess2

import (
	"strings"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/uuid"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type File struct {
	filePath      string
	name          string
	offset        int
	updatedStepID uuid.UUID
	blobHash      plumbing.Hash //blobHash is equivalent to content, assuming every transition is git
	// ref        FileContentRef //when manual edit is allowed, blobHash is not enought to indicate the contnet
}

func NewFile(file *object.File, createdStepID uuid.UUID) *File {
	split := strings.Split(file.Name, "/")
	name := split[len(split)-1]
	offset := len(split) - 1

	// file.IsBinary()
	// isFullContent := file.Size < 1000000 // 1MB

	return &File{
		name:          name,
		filePath:      file.Name,
		offset:        offset,
		blobHash:      file.Hash,
		updatedStepID: createdStepID,
	}
}

// the same File can have different isUpdated value at different steps
func (n *File) IsUpdated(currentStepID uuid.UUID) bool {
	return n.updatedStepID == currentStepID
}

func (n *File) Diff(from SourceCodeStep) string {
	return ""
}

// the File file can be transformed to different GraphQL models at different steps
func (n *File) ToGraphQLFile(currentStepID uuid.UUID) *model.FileNode {
	nodeType := model.FileNodeTypeFile
	name := n.name         // copy to avoid mutation effect afterwards
	filePath := n.filePath // copy to avoid mutation effect afterwards
	offset := n.offset     // copy to avoid mutation effect afterwards
	isUpdated := n.IsUpdated(currentStepID)

	return &model.FileNode{
		NodeType:  &nodeType,
		Name:      &name,
		FilePath:  &filePath,
		Offset:    &offset,
		IsUpdated: &isUpdated,
	}
}
