package effect

import (
	"fmt"
	"log"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

// func GitEffectProcessing() error {
// 	// 1. read effects from repository
// 	log.Printf("EffectProcessing started")

// 	stepEffects, err := GitStepEffects(nil)
// 	if err != nil {
// 		return fmt.Errorf("GitStepEffects failed: %v", err)
// 	}

// 	// 3. apply page-state operation and write states to files
// 	state := processing.NewPageStateProcessor()
// 	for i, step := range stepEffects {
// 		state.TransitionToNext()
// 		op, err := stepEffects[i].ToOperation()
// 		if err != nil {
// 			return fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
// 		}

// 		// after registering the next op, write to the file
// 		nextStep := step.NextStep
// 		state.RegisterNext(nextStep, &op)
// 		WriteJsonToFile(state.ToGraphQLPageState(), fmt.Sprintf("data/state/page-state%s.json", stepEffects[i].CurrentStep))

// 		// iterate over to the next state
// 		state.TransitionToNext()
// 	}
// 	// last state writes to the file
// 	lastStep := stepEffects[len(stepEffects)-1].CurrentStep
// 	WriteJsonToFile(state.ToGraphQLPageState(), fmt.Sprintf("data/state/page-state%s.json", lastStep))

// 	return nil
// }

func EffectProcessing() error {
	// 1. read effects from files
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

	// 2. construct page-sate effect
	var pageStateEffects []PageStateEffect
	for _, step := range stepEffects {
		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := NewSourceCodeEffect(step.SeqNo, fEffs)

		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// PageStateEffect for seqNo
		psEff := NewPageStateEffect(step.SeqNo, scEff, tEff)
		pageStateEffects = append(pageStateEffects, *psEff)
	}
	log.Printf("%d page state effects calculated", len(pageStateEffects))

	// 3. apply page-state operation and write states to files
	state := processing.NewPageStateProcessor()
	for i, step := range stepEffects {
		op, err := pageStateEffects[i].ToOperation()
		if err != nil {
			return fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}

		// after registering the next op, write to the file
		nextStep := step.NextStep
		state.RegisterNext(nextStep, &op)
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
