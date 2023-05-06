package processing

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProcessorTransition struct {
	Step             string            `json:"step"`
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

type SourceCodeUnitEffect struct {
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

func ReadTerminalEffects(filePath string) ([]TerminalEffect, error) {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ReadTerminalEffects failed to read file: %w", err)
	}

	var terminalEffects []TerminalEffect
	err = json.Unmarshal(jsonBytes, &terminalEffects)
	if err != nil {
		return nil, fmt.Errorf("ReadTerminalEffects failed to unmarshal: %w", err)
	}

	return terminalEffects, nil
}
