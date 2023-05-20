package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

type SourceCodeEffect struct {
	SeqNo               int
	CommitHash          string
	FileEffects         FileEffects
	DefaultOpenFilePath *string
}

type OpenFileEffect struct {
	SeqNo               int    `json:"seqNo"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
}

func (s *SourceCodeEffect) ToOperation() (processing.SourceCodeOperation, error) {
	if s.CommitHash == "" {
		fileOps, err := s.FileEffects.ToOperation()
		if err != nil {
			return processing.SourceCodeFileOperation{}, fmt.Errorf("ToOperation() in SourceCodeEffect failed: %v", err)
		}
		return processing.SourceCodeFileOperation{FileOps: fileOps}, nil
	} else {
		return processing.SourceCodeGitOperation{CommitHash: s.CommitHash}, nil
	}
}
