package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

type PageStateEffect struct {
	SeqNo               int
	SourceCodeEffect    *SourceCodeEffect    // only either SourceCodeEffect or SourceCodeGitEffect is held
	SourceCodeGitEffect *SourceCodeGitEffect // only either SourceCodeEffect or SourceCodeGitEffect is held
	TerminalEffect      *TerminalEffect
}

func NewPageStateEffect(
	seqNo int,
	sourceCodeEffect *SourceCodeEffect,
	terminalEffect *TerminalEffect,
) *PageStateEffect {
	return &PageStateEffect{SeqNo: seqNo, SourceCodeEffect: sourceCodeEffect, TerminalEffect: terminalEffect}
}

func NewPageStateGitEffect(
	seqNo int,
	sourceCodeGitEffect *SourceCodeGitEffect,
	terminalEffect *TerminalEffect,
) *PageStateEffect {
	return &PageStateEffect{SeqNo: seqNo, SourceCodeGitEffect: sourceCodeGitEffect, TerminalEffect: terminalEffect}
}

func (p *PageStateEffect) ToOperation() (processing.PageStateOperation, error) {
	var sourceCodeOp processing.SourceCodeOperation
	var err error

	if p.SourceCodeEffect != nil && p.SourceCodeGitEffect != nil {
		return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: only either of SourceCodeEffect or SourceCodeGitEffect can be set")
	} else if p.SourceCodeEffect != nil {
		if sourceCodeOp, err = p.SourceCodeEffect.ToOperation(); err != nil {
			return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	} else if p.SourceCodeGitEffect != nil {
		if sourceCodeOp, err = p.SourceCodeGitEffect.ToOperation(); err != nil {
			return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	} else {
		sourceCodeOp = nil
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
		SourceCodeOperation: sourceCodeOp,
		TerminalOperation:   terminalOp,
	}, nil
}
