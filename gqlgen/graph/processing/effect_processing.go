package processing

import (
	"fmt"
	"log"
)

func EffectProcessing() error {
	log.Printf("EffectProcessing started")

	stepEffects, err := ReadStepEffects("data/step-effects.json")
	if err != nil {
		return fmt.Errorf("EffectProcessing failed: %v", err)
	}
	log.Printf("%d step effects read ", len(stepEffects))

	fileEffects, err := ReadFileEffects("data/file-effects.json")
	if err != nil {
		return fmt.Errorf("EffectProcessing failed: %v", err)
	}
	log.Printf("%d file effects read ", len(fileEffects))

	terminalEffects, err := ReadTerminalEffects("data/terminal-effects.json")
	if err != nil {
		return fmt.Errorf("EffectProcessing failed: %v", err)
	}
	log.Printf("%d terminal effects read ", len(terminalEffects))

	var ops []PageStateOperation
	for _, step := range stepEffects {
		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := NewSourceCodeEffect(step.SeqNo, fEffs)

		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// PageStateEffect for seqNo
		psEff := NewPageStateEffect(step.SeqNo, scEff, tEff)

		op, err := psEff.ToOperation()
		if err != nil {
			return fmt.Errorf("EffectProcessing failed: %v", err)
		}
		ops = append(ops, op)
	}
	log.Printf("%d page state operations calculated", len(ops))

	state := NewPageStateProcessor()
	for i := 0; i < len(ops); i++ {
		// after registering the next op, write to the file
		nextStep := stepEffects[i].NextStep
		state.RegisterNext(nextStep, &ops[i])
		WriteJsonToFile(state.ToGraphQLPageState(), fmt.Sprintf("data/state/page-state%s.json", stepEffects[i].CurrentStep))

		// iterate over to the next state
		state.TransitionToNext()
	}
	// last state writes to the file
	lastStep := stepEffects[len(stepEffects)-1].CurrentStep
	WriteJsonToFile(state.ToGraphQLPageState(), fmt.Sprintf("data/state/page-state%s.json", lastStep))

	log.Printf("finished writing state into files")

	return nil
}
