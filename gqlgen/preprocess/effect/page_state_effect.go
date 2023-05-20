package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

type PageStateEffect struct {
	SeqNo            int
	SourceCodeEffect *SourceCodeEffect // only either SourceCodeEffect or SourceCodeGitEffect is held
	TerminalEffect   *TerminalEffect
	MarkdownEffect   *MarkdownEffect
}

func (p *PageStateEffect) ToOperation() (processing.PageStateOperation, error) {
	var sourceCodeOp processing.SourceCodeOperation
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
		SourceCodeOperation: sourceCodeOp,
		TerminalOperation:   terminalOp,
	}, nil
}
