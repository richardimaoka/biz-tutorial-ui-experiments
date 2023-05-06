package processing

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProcessorTransition struct {
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
	SeqNo         int     `json:"seqNo"`
	OperationType string  `json:"operationType"`
	FilePath      *string `json:"filePath"`
	Content       *string `json:"content"`
}

type SourceCodeOpenFileEffect struct {
	SeqNo               int    `json:"seqNo"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
}

// TODO: later optimization
// type SourceCodeGitEffect struct {
// 	commitHash string
// }

func MergeEffects(
	terminalEffects []TerminalEffect,
	sourceCodeUnitEffects []FileEffect,
) ([]ProcessorTransition, error) {
	// 1. calculate the max seqNo
	maxSeqNo := 0
	for _, t := range terminalEffects {
		if t.SeqNo > maxSeqNo {
			maxSeqNo = t.SeqNo
		}
	}
	for _, t := range sourceCodeUnitEffects {
		if t.SeqNo > maxSeqNo {
			maxSeqNo = t.SeqNo
		}
	}

	var transitions []ProcessorTransition
	for i := 0; i < maxSeqNo; i++ {
		p := ProcessorTransition{SeqNo: i}
		transitions = append(transitions, p)
	}

	return transitions, nil
}

func findEffectsBySeqNo(seqNo int, sourceCodeUnitEffects []FileEffect) []FileEffect {
	var effects []FileEffect
	for _, e := range sourceCodeUnitEffects {
		if e.SeqNo == seqNo {
			effects = append(effects, e)
		}
	}
	return effects
}

func ReadTerminalEffects(filePath string) ([]TerminalEffect, error) {
	var effects []TerminalEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead("ReadTerminalEffects", filePath, unmarshaller)
	return effects, err
}

func ReadSourceCodeUnitEffect(filePath string) ([]FileEffect, error) {
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
