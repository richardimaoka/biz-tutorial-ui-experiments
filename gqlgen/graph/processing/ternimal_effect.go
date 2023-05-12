package processing

import (
	"encoding/json"
	"fmt"
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
	err := jsonRead("ReadTerminalEffects", filePath, unmarshaller)
	return effects, err
}

func (t TerminalEffects) FindBySeqNo(seqNo int) *TerminalEffect {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e
		}
	}

	return nil
}

// TODO: remove this function
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

func (t TerminalEffect) ToOperation() TerminalOperation {
	if t.Output == nil && t.CurrentDirectory == nil {
		return TerminalCommand{Command: t.Command}
	} else if t.Output != nil && t.CurrentDirectory == nil {
		return TerminalCommandWithOutput{Command: t.Command, Output: *t.Output}
	} else if t.Output == nil && t.CurrentDirectory != nil {
		return TerminalCommandWithCd{Command: t.Command, CurrentDirectory: *t.CurrentDirectory}
	} else if t.Output != nil && t.CurrentDirectory != nil {
		return TerminalCommandWithOutputCd{Command: t.Command, Output: *t.Output, CurrentDirectory: *t.CurrentDirectory}
	} else {
		// this should never happen
		return nil
	}
}
