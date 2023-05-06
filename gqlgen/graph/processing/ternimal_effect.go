package processing

import "fmt"

type TerminalEffect struct {
	SeqNo            int     `json:"seqNo"`
	TerminalName     string  `json:"terminalName"`
	Command          string  `json:"command"`
	Output           *string `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution
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
