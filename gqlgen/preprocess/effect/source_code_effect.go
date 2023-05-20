package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

type SourceCodeEffect struct {
	seqNo               int
	commitHash          string
	fileEffects         FileEffects
	defaultOpenFilePath *string
}

func NewSourceCodeEffect(seqNo int, commitHash string, fileEffects FileEffects) *SourceCodeEffect {
	return &SourceCodeEffect{seqNo, commitHash, fileEffects, nil}
}

type OpenFileEffect struct {
	SeqNo               int    `json:"seqNo"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
}

func (s *SourceCodeEffect) ToOperation() (processing.SourceCodeOperation, error) {
	if s.commitHash == "" {
		fileOps, err := s.fileEffects.ToOperation()
		if err != nil {
			return processing.SourceCodeFileOperation{}, fmt.Errorf("ToOperation() in SourceCodeEffect failed: %v", err)
		}
		return processing.SourceCodeFileOperation{FileOps: fileOps}, nil
	} else {
		return processing.SourceCodeGitOperation{CommitHash: s.commitHash}, nil
	}
}
