package effect

import (
	"fmt"
	"log"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func GitEffectProcessing() error {
	log.Printf("GitEffectProcessing started")
	repoUrl := "https://github.com/richardimaoka/gqlgensandbox"

	//------------------------------------
	// 1. read effects from git repository
	//------------------------------------
	stepEffects, err := GitStepEffects(repoUrl)
	if err != nil {
		return fmt.Errorf("GitStepEffects failed: %v", err)
	}
	log.Printf("%d step effects read ", len(stepEffects))

	//------------------------------
	// 2. construct page-sate effect
	//------------------------------
	var pageStateEffects []PageStateEffect
	for _, step := range stepEffects {
		// SourceCodeEffect for seqNo
		scEff := NewSourceCodeGitEffect(step.SeqNo, step.CommitHash)

		// PageStateEffect for seqNo
		psEff := NewPageStateGitEffect(step.SeqNo, scEff, nil)
		pageStateEffects = append(pageStateEffects, *psEff)
	}
	log.Printf("%d page state effects calculated", len(pageStateEffects))

	//--------------------------------------------------------
	// 2. apply page-state operation and write states to files
	//--------------------------------------------------------
	state, err := processing.NewPageStateGitProcessorFromGit(repoUrl)
	if err != nil {
		return fmt.Errorf("NewPageStateGitProcessorFromGit() failed: %v", err)
	}

	for i, step := range stepEffects {
		state.TransitionToNext()
		op, err := pageStateEffects[i].ToOperation()
		if err != nil {
			return fmt.Errorf("ToOperation() failed at %s: %v", step.CurrentStep, err)
		}

		nextStep := step.NextStep
		if err := state.RegisterNext(nextStep, &op); err != nil {
			return fmt.Errorf("RegisterNext() failed at step %s: %v", step.CurrentStep, err)
		}

		fileName := fmt.Sprintf("data/state/page-state%s.json", stepEffects[i].CurrentStep)
		if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("WriteJsonToFile() failed at step %s: %v", step.CurrentStep, err)
		}

		// iterate over to the next state
		if err := state.TransitionToNext(); err != nil {
			return fmt.Errorf("TransitionToNext() failed at step %s: %v", step.CurrentStep, err)
		}
	}
	// last state writes to the file
	lastStep := stepEffects[len(stepEffects)-1].CurrentStep
	fileName := fmt.Sprintf("data/state/page-state%s.json", lastStep)
	if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("WriteJsonToFile() in PageStateProcessor failed: %v", err)
	}

	log.Printf("finished writing state into files")
	return nil
}

func EffectProcessing() error {
	log.Printf("EffectProcessing started")

	//---------------------------
	// 1. read effects from files
	//---------------------------
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

	//------------------------------
	// 2. construct page-sate effect
	//------------------------------
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

	//--------------------------------------------------------
	// 3. apply page-state operation and write states to files
	//--------------------------------------------------------
	state := processing.NewPageStateProcessor()
	for i, step := range stepEffects {
		op, err := pageStateEffects[i].ToOperation()
		if err != nil {
			return fmt.Errorf("ToOperation() in PageStateEffect failed: %v", err)
		}

		nextStep := step.NextStep
		if err := state.RegisterNext(nextStep, &op); err != nil {
			return fmt.Errorf("RegisterNext() in PageStateProcessor failed: %v", err)
		}

		fileName := fmt.Sprintf("data/state/page-state%s.json", stepEffects[i].CurrentStep)
		if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("WriteJsonToFile() in PageStateProcessor failed: %v", err)
		}

		// iterate over to the next state
		if err := state.TransitionToNext(); err != nil {
			return fmt.Errorf("TransitionToNext() in PageStateProcessor failed: %v", err)
		}
	}

	// last state writes to the file
	lastStep := stepEffects[len(stepEffects)-1].CurrentStep
	fileName := fmt.Sprintf("data/state/page-state%s.json", lastStep)
	if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("WriteJsonToFile() in PageStateProcessor failed: %v", err)
	}

	log.Printf("finished writing state into files")
	return nil
}
