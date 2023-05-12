package processing

import (
	"fmt"
)

type PageStateEffect struct {
	Step             string            `json:"step"`
	SeqNo            int               `json:"seqNo"`
	SourceCodeEffect *SourceCodeEffect `json:"sourceCodeEffect"`
	TerminalEffect   *TerminalEffect   `json:"terminalEffect"`
}

func NewPageStateEffect(
	seqNo int,
	step string,
	sourceCodeEffect *SourceCodeEffect,
	terminalEffect *TerminalEffect) *PageStateEffect {
	return &PageStateEffect{SeqNo: seqNo, Step: step, SourceCodeEffect: sourceCodeEffect, TerminalEffect: terminalEffect}
}

func MergeEffects(terminalEffects []TerminalEffect, fileEffects []FileEffect) ([]PageStateEffect, error) {
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
	var effects []PageStateEffect
	for seqNo := 0; seqNo < maxSeqNo; seqNo++ {

		tEff, err := terminalEffectBySeqNo(seqNo, terminalEffects)
		if err != nil {
			return nil, fmt.Errorf("MergeEffects failed: %v", err)
		}

		sEff, err := CalculateSourceCodeEffect(seqNo, fileEffects)
		if err != nil {
			return nil, fmt.Errorf("MergeEffects failed: %v", err)
		}

		p := PageStateEffect{SeqNo: seqNo, TerminalEffect: tEff, SourceCodeEffect: sEff}
		effects = append(effects, p)
	}

	return effects, nil
}
