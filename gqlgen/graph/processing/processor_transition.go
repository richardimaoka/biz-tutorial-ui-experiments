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
	var effects []TerminalEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead("ReadTerminalEffects", filePath, unmarshaller)
	return effects, err
}

func ReadSourceCodeUnitEffect(filePath string) ([]SourceCodeUnitEffect, error) {
	var effects []SourceCodeUnitEffect
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
