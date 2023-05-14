package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

type SourceCodeEffect struct {
	SeqNo               int
	FileEffects         FileEffects
	DefaultOpenFilePath *string
}

type SourceCodeGitEffect struct {
	SeqNo               int     `json:"seqNo"`
	CommitHash          string  `json:"commitHash"`
	DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
}

type OpenFileEffect struct {
	SeqNo               int    `json:"seqNo"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
}

func NewSourceCodeEffect(seqNo int, effects []FileEffect) *SourceCodeEffect {
	return &SourceCodeEffect{SeqNo: seqNo, FileEffects: effects}
}

func (s *SourceCodeEffect) ToOperation() (*processing.SourceCodeFileOperation, error) {
	fileOps, err := s.FileEffects.ToOperation()
	if err != nil {
		return &processing.SourceCodeFileOperation{}, fmt.Errorf("ToOperation() in SourceCodeEffect failed: %v", err)
	}

	return &processing.SourceCodeFileOperation{FileOps: fileOps}, nil
}

func (s *SourceCodeGitEffect) ToOperation() (*processing.SourceCodeGitOperation, error) {
	return nil, nil
}
