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

	var ops []PageStateOperation
	for _, step := range stepEffects {
		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := NewSourceCodeEffect(step.SeqNo, fEffs)

		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// PageStateEffect for seqNo
		psEff := NewPageStateEffect(step.SeqNo, step.Step, scEff, tEff)

		op, err := psEff.ToOperation()
		if err != nil {
			return fmt.Errorf("EffectProcessing failed: %v", err)
		}
		ops = append(ops, op)
	}

	state := NewPageStateProcessor()

	for i := 0; i < len(ops); i++ {
		// after registering the next op, write to the file
		state.RegisterNext(stepEffects[i].Step, &ops[i])
		WriteJsonToFile(state, "data/page-state.json")

		// iterate over to the next state
		state.TransitionToNext()
	}
	// last state
	WriteJsonToFile(state, "data/page-state.json")

	return nil
}
