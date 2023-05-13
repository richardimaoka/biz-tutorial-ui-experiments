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

func (s *SourceCodeEffect) ToOperation() (*processing.SourceCodeOperation, error) {
	fileOps, err := s.FileEffects.ToOperation()
	if err != nil {
		return &processing.SourceCodeOperation{}, fmt.Errorf("ToOperation() in SourceCodeEffect failed: %v", err)
	}

	return &processing.SourceCodeOperation{FileOps: fileOps}, nil
}
