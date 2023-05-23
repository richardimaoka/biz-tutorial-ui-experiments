package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

// combined effect, so no json meta tag, and can only be constructed from a *New* function

type PageStateEffect struct {
	seqNo            int
	Step             string //TODO: move this out to an outer struct?
	NextStep         string //TODO: move this out to an outer struct?
	sourceCodeEffect *SourceCodeEffect
	terminalEffect   *TerminalEffect
	markdownEffect   *MarkdownEffect
}

func NewPageStateEffect(seqNo int, step, nextStep string, sourceCodeEffect *SourceCodeEffect, terminalEffect *TerminalEffect, markdownEffect *MarkdownEffect) *PageStateEffect {
	return &PageStateEffect{seqNo, step, nextStep, sourceCodeEffect, terminalEffect, markdownEffect}
}

func ConstructPageStateEffects(stepEffectsFile, fileEffectsFile, terminalEffectsFile, markdownEffectsFile string) ([]*PageStateEffect, error) {
	//------------------------------------
	// 1. read effects from files
	//------------------------------------
	stepEffects, err := ReadStepEffects(stepEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	fileEffects, err := ReadFileEffects(fileEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	terminalEffects, err := ReadTerminalEffects(terminalEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	markdownEffects, err := ReadMarkdownEffects(markdownEffectsFile)
	if err != nil {
		return nil, fmt.Errorf("pageStateEffects failed: %v", err)
	}

	//------------------------------
	// 2. construct page-sate effect
	//------------------------------
	var pageStateEffects []*PageStateEffect
	for i, step := range stepEffects {
		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := NewSourceCodeEffect(step.SeqNo, step.CommitHash, fEffs)

		// MarkdownEffect for seqNo
		mEff := markdownEffects.FindBySeqNo(step.SeqNo)

		var nextStep string
		if i == len(stepEffects)-1 {
			nextStep = ""
		} else {
			nextStep = stepEffects[i+1].Step
		}

		// PageStateEffect for seqNo
		psEff := NewPageStateEffect(step.SeqNo, step.Step, nextStep, scEff, tEff, mEff)
		pageStateEffects = append(pageStateEffects, psEff)
	}

	return pageStateEffects, nil
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
