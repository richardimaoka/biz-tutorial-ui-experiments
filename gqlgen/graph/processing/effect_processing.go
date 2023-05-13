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
		log.Printf("step: %d", step.SeqNo)
		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		log.Printf("fEffs")
		scEff := NewSourceCodeEffect(step.SeqNo, fEffs)
		log.Printf("scEffs")

		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)
		log.Printf("tEffs")

		// PageStateEffect for seqNo
		psEff := NewPageStateEffect(step.SeqNo, step.Step, scEff, tEff)
		log.Printf("psEff %+v", psEff)

		op, err := psEff.ToOperation()
		log.Printf("op")
		if err != nil {
			return fmt.Errorf("EffectProcessing failed: %v", err)
		}
		log.Printf("op 2")
		ops = append(ops, op)
		log.Printf("ops")
	}

	state := NewPageStateProcessor()

	for i := 0; i < len(ops); i++ {
		// after registering the next op, write to the file
		step := stepEffects[i].Step
		state.RegisterNext(step, &ops[i])
		WriteJsonToFile(state, fmt.Sprintf("data/state/page-state%s.json", step))

		// iterate over to the next state
		state.TransitionToNext()
	}
	// last state writes to the file
	lastStep := stepEffects[len(stepEffects)-1].Step
	WriteJsonToFile(state, fmt.Sprintf("data/state/page-state%s.json", lastStep))

	return nil
}
