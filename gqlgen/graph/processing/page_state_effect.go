package processing

import (
	"fmt"
)

type PageStateEffect struct {
	SeqNo            int
	SourceCodeEffect *SourceCodeEffect
	TerminalEffect   *TerminalEffect
}

func NewPageStateEffect(
	seqNo int,
	sourceCodeEffect *SourceCodeEffect,
	terminalEffect *TerminalEffect) *PageStateEffect {
	return &PageStateEffect{SeqNo: seqNo, SourceCodeEffect: sourceCodeEffect, TerminalEffect: terminalEffect}
}

func (p *PageStateEffect) ToOperation() (PageStateOperation, error) {
	var sourceCodeOp *SourceCodeOperation
	if p.SourceCodeEffect == nil {
		sourceCodeOp = nil
	} else {
		var err error
		sourceCodeOp, err = p.SourceCodeEffect.ToOperation()
		if err != nil {
			return PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	}

	var terminalOp TerminalOperation
	if p.TerminalEffect == nil {
		terminalOp = nil
	} else {
		var err error
		terminalOp, err = p.TerminalEffect.ToOperation()
		if err != nil {
			return PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	}

	return PageStateOperation{
		SourceCodeOperation: sourceCodeOp,
		TerminalOperation:   terminalOp,
	}, nil
}

// TODO: remove this function
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
