package processing

import (
	"fmt"
)

type SourceCodeEffect struct {
	SeqNo               int
	Diff                Diff //TODO: remove this
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

func (s *SourceCodeEffect) ToOperation() (*SourceCodeOperation, error) {
	fileOps, err := s.FileEffects.ToOperation()
	if err != nil {
		return &SourceCodeOperation{}, fmt.Errorf("ToOperation() in SourceCodeEffect failed: %v", err)
	}

	return &SourceCodeOperation{FileOps: fileOps}, nil
}

// TODO: remove this function
func CalculateSourceCodeEffect(seqNo int, effects []FileEffect) (*SourceCodeEffect, error) {
	effectsBySeqNo := fileEffectsBySeqNo(seqNo, effects)

	if len(effectsBySeqNo) == 0 {
		return nil, nil
	}

	effect := SourceCodeEffect{SeqNo: seqNo}
	for _, e := range effectsBySeqNo {
		op, err := e.ToOperation()
		if err != nil {
			return nil, fmt.Errorf("failed to calculate source code effect: %v", err)
		}

		effect.Diff.Append(op)
	}

	return &effect, nil
}
