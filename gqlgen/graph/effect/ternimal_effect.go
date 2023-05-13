package effect

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

type TerminalEffect struct {
	SeqNo            int     `json:"seqNo"`
	TerminalName     string  `json:"terminalName"`
	Command          string  `json:"command"`
	Output           *string `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution
}

type TerminalEffects []TerminalEffect

func ReadTerminalEffects(filePath string) (TerminalEffects, error) {
	var effects []TerminalEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead(filePath, unmarshaller)
	if err != nil {
		return nil, fmt.Errorf("ReadFileEffects failed to read file, %s", err)
	}

	return effects, err
}

func (t TerminalEffects) FindBySeqNo(seqNo int) *TerminalEffect {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e // found!
		}
	}

	return nil
}

func (t TerminalEffect) ToOperation() (processing.TerminalOperation, error) {
	if t.Output == nil && t.CurrentDirectory == nil {
		return processing.TerminalCommand{TerminalName: t.TerminalName, Command: t.Command}, nil
	} else if t.Output != nil && t.CurrentDirectory == nil {
		return processing.TerminalCommandWithOutput{TerminalName: t.TerminalName, Command: t.Command, Output: *t.Output}, nil
	} else if t.Output == nil && t.CurrentDirectory != nil {
		return processing.TerminalCommandWithCd{TerminalName: t.TerminalName, Command: t.Command, CurrentDirectory: *t.CurrentDirectory}, nil
	} else if t.Output != nil && t.CurrentDirectory != nil {
		return processing.TerminalCommandWithOutputCd{TerminalName: t.TerminalName, Command: t.Command, Output: *t.Output, CurrentDirectory: *t.CurrentDirectory}, nil
	} else {
		// this should never happen
		return nil, fmt.Errorf("ToOperation() in TerminalEffect failed: terminal effect = %+v", t)
	}
}
