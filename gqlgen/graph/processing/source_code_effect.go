package processing

import (
	"encoding/json"
	"fmt"

	"github.com/go-git/go-git/v5/plumbing"
)

type SourceCodeEffect struct {
	SeqNo               int     `json:"seqNo"`
	Diff                Diff    `json:"diff"`
	DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
}

type SourceCodeGitEffect struct {
	SeqNo               int           `json:"seqNo"`
	CommitHash          plumbing.Hash `json:"commitHash"`
	DefaultOpenFilePath *string       `json:"defaultOpenFilePath"`
}

type FileEffect struct {
	SeqNo         int    `json:"seqNo"`
	OperationType string `json:"operationType"`
	FilePath      string `json:"filePath"`
	Content       string `json:"content"`
}

type OpenFileEffect struct {
	SeqNo               int    `json:"seqNo"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
}

// TODO: later optimization
// type SourceCodeGitEffect struct {
// 	commitHash string
// }

func ReadFileEffects(filePath string) ([]FileEffect, error) {
	var effects []FileEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead("ReadFileEffects", filePath, unmarshaller)
	return effects, err
}

func calculateSourceCodeEffect(seqNo int, effects []FileEffect) (*SourceCodeEffect, error) {
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

func fileEffectsBySeqNo(seqNo int, effects []FileEffect) []FileEffect {
	var effectsBySeqNo []FileEffect
	for _, e := range effects {
		if e.SeqNo == seqNo {
			effectsBySeqNo = append(effectsBySeqNo, e)
		}
	}
	return effectsBySeqNo
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
