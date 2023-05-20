package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

// combined effect, so no json meta tag, and can only be constructed from a *New* function

type PageStateEffect struct {
	seqNo            int
	sourceCodeEffect *SourceCodeEffect
	terminalEffect   *TerminalEffect
	markdownEffect   *MarkdownEffect
}

func NewPageStateEffect(seqNo int, sourceCodeEffect *SourceCodeEffect, terminalEffect *TerminalEffect, markdownEffect *MarkdownEffect) *PageStateEffect {
	return &PageStateEffect{seqNo, sourceCodeEffect, terminalEffect, markdownEffect}
}

func (p *PageStateEffect) ToOperation() (processing.PageStateOperation, error) {
	var sourceCodeOp processing.SourceCodeOperation
	if p.sourceCodeEffect == nil {
		sourceCodeOp = nil
	} else {
		var err error
		sourceCodeOp, err = p.sourceCodeEffect.ToOperation()
		if err != nil {
			return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	}

	var terminalOp processing.TerminalOperation
	if p.terminalEffect == nil {
		terminalOp = nil
	} else {
		var err error
		terminalOp, err = p.terminalEffect.ToOperation()
		if err != nil {
			return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	}

	var markdownOp *processing.MarkdownOperation
	if p.markdownEffect == nil {
		markdownOp = nil
	} else {
		var err error
		markdownOp, err = p.markdownEffect.ToOperation()
		if err != nil {
			return processing.PageStateOperation{}, fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}
	}

	return processing.PageStateOperation{
		SourceCodeOperation: sourceCodeOp,
		TerminalOperation:   terminalOp,
		MarkdownOperation:   markdownOp,
	}, nil
}
