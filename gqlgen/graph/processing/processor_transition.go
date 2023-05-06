package processing

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProcessorEffect struct {
	Step             string            `json:"step"`
	SeqNo            int               `json:"seqNo"`
	SourceCodeEffect *SourceCodeEffect `json:"sourceCodeEffect"`
	TerminalEffect   *TerminalEffect   `json:"terminalEffect"`
}

type TerminalEffect struct {
	SeqNo            int     `json:"seqNo"`
	TerminalName     string  `json:"terminalName"`
	Command          string  `json:"command"`
	Output           *string `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution
}

type SourceCodeEffect struct {
	SeqNo               int     `json:"seqNo"`
	Diff                Diff    `json:"diff"`
	DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
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

func MergeEffects(terminalEffects []TerminalEffect, fileEffects []FileEffect) ([]ProcessorEffect, error) {
	// 1. calculate the max seqNo
	maxSeqNo := 0
	for _, t := range terminalEffects {
		if t.SeqNo > maxSeqNo {
			maxSeqNo = t.SeqNo
		}
	}
	for _, t := range fileEffects {
		if t.SeqNo > maxSeqNo {
			maxSeqNo = t.SeqNo
		}
	}

	// 2. construct ProcessorEffect for each seqNo
	var transitions []ProcessorEffect
	for seqNo := 0; seqNo < maxSeqNo; seqNo++ {
		tEff, err := terminalEffectBySeqNo(seqNo, terminalEffects)
		if err != nil {
			return nil, fmt.Errorf("MergeEffects failed: %v", err)
		}

		sEff, err := calculateSourceCodeEffect(seqNo, fileEffects)
		if err != nil {
			return nil, fmt.Errorf("MergeEffects failed: %v", err)
		}

		p := ProcessorEffect{SeqNo: seqNo, TerminalEffect: tEff, SourceCodeEffect: sEff}
		transitions = append(transitions, p)
	}

	return transitions, nil
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

func terminalEffectBySeqNo(seqNo int, effects []TerminalEffect) (*TerminalEffect, error) {
	var effectsBySeqNo []TerminalEffect
	for _, e := range effects {
		if e.SeqNo == seqNo {
			effectsBySeqNo = append(effectsBySeqNo, e)
		}
	}

	if len(effectsBySeqNo) > 1 {
		return nil, fmt.Errorf("multiple (%d of) TerminalEffect with the same seqNo = %d found", len(effectsBySeqNo), seqNo)
	}

	if len(effectsBySeqNo) == 1 { // must be len(effectsBySeqNo) == 1
		return &effectsBySeqNo[0], nil
	} else { // must be len(effectsBySeqNo) == 0
		return nil, nil
	}
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

func ReadTerminalEffects(filePath string) ([]TerminalEffect, error) {
	var effects []TerminalEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead("ReadTerminalEffects", filePath, unmarshaller)
	return effects, err
}

func ReadFileEffects(filePath string) ([]FileEffect, error) {
	var effects []FileEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead("ReadSourceCodeUnitEffect", filePath, unmarshaller)
	return effects, err
}

func jsonRead(callerName, filePath string, unmarshaller func(jsonBytes []byte) error) error {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("%s failed to read file, %s", callerName, err)
	}

	err = unmarshaller(jsonBytes)
	if err != nil {
		return fmt.Errorf("%s failed to unmarshal, %s", callerName, err)
	}

	return nil
}
