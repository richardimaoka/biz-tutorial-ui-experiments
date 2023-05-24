package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

// combined effect, so no json meta tag, and can only be constructed from a *New* function
type PageStateEffect struct {
	seqNo            int
	Step             string
	sourceCodeEffect *SourceCodeEffect
	terminalEffect   *TerminalEffect
	markdownEffect   *MarkdownEffect
}

func NewPageStateEffect(
	seqNo int,
	step string,
	sourceCodeEffect *SourceCodeEffect,
	terminalEffect *TerminalEffect,
	markdownEffect *MarkdownEffect,
) *PageStateEffect {

	return &PageStateEffect{
		seqNo,
		step,
		sourceCodeEffect,
		terminalEffect,
		markdownEffect,
	}
}

func ConstructPageStateEffects(dirName string) ([]PageStateEffect, error) {
	//------------------------------------
	// 1. read effects from files
	//------------------------------------
	stepEffectsFile := dirName + "/step-effects.json"
	stepEffects, err := ReadStepEffects(stepEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	fileEffectsFile := dirName + "/file-effects.json"
	fileEffects, err := ReadFileEffects(fileEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	terminalEffectsFile := dirName + "/terminal-effects.json"
	terminalEffects, err := ReadTerminalEffects(terminalEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	markdownEffectsFile := dirName + "/markdown-effects.json"
	markdownEffects, err := ReadMarkdownEffects(markdownEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	//------------------------------
	// 2. construct page-sate effect
	//------------------------------
	var effects []PageStateEffect
	for _, step := range stepEffects {
		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := NewSourceCodeEffect(step.SeqNo, step.CommitHash, fEffs)

		// MarkdownEffect for seqNo
		mEff := markdownEffects.FindBySeqNo(step.SeqNo)

		// PageStateEffect for seqNo
		psEff := NewPageStateEffect(step.SeqNo, step.Step, scEff, tEff, mEff)

		effects = append(effects, *psEff)
	}

	return effects, nil
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
