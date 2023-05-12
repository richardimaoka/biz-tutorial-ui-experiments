package processing

import (
	"fmt"
)

type SourceCodeEffect struct {
	SeqNo               int
	Diff                Diff //TODO: remove this
	FileEffects         []FileEffect
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

func (f FileEffect) ToOperation() (FileSystemOperation, error) {
	switch f.OperationType {
	case "FileAdd":
		return FileAdd{FilePath: f.FilePath, Content: f.Content, IsFullContent: true}, nil
	case "FileUpdate":
		return FileUpdate{FilePath: f.FilePath, Content: f.Content}, nil
	case "FileDelete":
		return FileDelete{FilePath: f.FilePath}, nil
	case "DirectoryAdd":
		return DirectoryAdd{FilePath: f.FilePath}, nil
	case "DirectoryDelete":
		return DirectoryDelete{FilePath: f.FilePath}, nil
	default:
		return nil, fmt.Errorf("FileEffect.ToOperation() found invalid OperationType = %s", f.OperationType)
	}
}
