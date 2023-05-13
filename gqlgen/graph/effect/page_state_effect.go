package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
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

func (p *PageStateEffect) ToOperation() (processing.PageStateOperation, error) {
	var sourceCodeOp *processing.SourceCodeFileOperation
	if p.SourceCodeEffect == nil {
		sourceCodeOp = nil
	} else {
		var err error
		sourceCodeOp, err = p.SourceCodeEffect.ToOperation()
		if err != nil {
			return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	}

	var terminalOp processing.TerminalOperation
	if p.TerminalEffect == nil {
		terminalOp = nil
	} else {
		var err error
		terminalOp, err = p.TerminalEffect.ToOperation()
		if err != nil {
			return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	}

	return processing.PageStateOperation{
		SourceCodeFileOperation: sourceCodeOp,
		TerminalOperation:       terminalOp,
	}, nil
}
