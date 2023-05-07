package processing

import (
	"fmt"
)

type ProcessorEffect struct {
	Step             string            `json:"step"`
	SeqNo            int               `json:"seqNo"`
	SourceCodeEffect *SourceCodeEffect `json:"sourceCodeEffect"`
	TerminalEffect   *TerminalEffect   `json:"terminalEffect"`
}

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
