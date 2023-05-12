package processing

import "fmt"

func EffectProcessing() error {
	stepEffects, err := ReadStepEffects("data/step-effects.json")
	if err != nil {
		return fmt.Errorf("EffectProcessing failed: %v", err)
	}

	fileEffects, err := ReadFileEffects("data/file-effects.json")
	if err != nil {
		return fmt.Errorf("EffectProcessing failed: %v", err)
	}

	terminalEffects, err := ReadTerminalEffects("data/terminal-effects.json")
	if err != nil {
		return fmt.Errorf("EffectProcessing failed: %v", err)
	}

	for _, step := range stepEffects {
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := NewSourceCodeEffect(step.SeqNo, fEffs)

		tEff := terminalEffects.FindBySeqNo(step.SeqNo)
		psEff := NewPageStateEffect(step.SeqNo, step.Step, scEff, tEff)
		fmt.Println(psEff)
	}

	return nil
}
