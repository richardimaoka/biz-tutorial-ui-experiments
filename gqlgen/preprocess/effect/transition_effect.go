package effect

import (
	"fmt"
)

type TransitionEffect struct {
	Step   string
	Effect *PageStateEffect
}

func ConstructTransitionEffects(dirName string) ([]TransitionEffect, error) {
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
	var effects []TransitionEffect
	for _, step := range stepEffects {
		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := NewSourceCodeEffect(step.SeqNo, step.CommitHash, fEffs)

		// MarkdownEffect for seqNo
		mEff := markdownEffects.FindBySeqNo(step.SeqNo)

		// PageStateEffect for seqNo
		psEff := NewPageStateEffect(step.SeqNo, "", "", scEff, tEff, mEff)

		effects = append(effects, TransitionEffect{step.Step, psEff})
	}

	return effects, nil
}
