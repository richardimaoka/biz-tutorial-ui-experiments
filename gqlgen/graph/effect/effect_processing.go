package effect

import (
	"fmt"
	"log"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func GitEffectProcessing(dirName, repoUrl string) error {
	log.Printf("GitEffectProcessing started")

	//------------------------------------
	// 1. read effects for git repository
	//------------------------------------
	stepEffects, err := ReadGeneralStepEffects(dirName + "/step-effects.json")
	if err != nil {
		return fmt.Errorf("processing failed: %v", err)
	}
	log.Printf("%d step effects read ", len(stepEffects))

	fileEffects, err := ReadFileEffects(dirName + "/file-effects.json")
	if err != nil {
		return fmt.Errorf("processing failed: %v", err)
	}
	log.Printf("%d file effects read ", len(fileEffects))

	terminalEffects, err := ReadTerminalEffects(dirName + "/terminal-effects.json")
	if err != nil {
		return fmt.Errorf("processing failed: %v", err)
	}
	log.Printf("%d terminal effects read ", len(terminalEffects))

	//------------------------------
	// 2. construct page-sate effect
	//------------------------------
	var pageStateEffects []PageStateEffect
	for _, step := range stepEffects {
		log.Printf("%+v", step)
		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		var psEff PageStateEffect
		if step.IsGitCommitStep() {
			// SourceCodeEffect for seqNo
			fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
			scEff := SourceCodeGitEffect{step.SeqNo, step.CommitHash, nil, fEffs}
			psEff = PageStateEffect{step.SeqNo, nil, &scEff, tEff}
		} else {
			// SourceCodeEffect for seqNo
			fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
			scEff := SourceCodeEffect{step.SeqNo, fEffs, nil}
			psEff = PageStateEffect{step.SeqNo, &scEff, nil, tEff}
		}

		// PageStateEffect for seqNo
		pageStateEffects = append(pageStateEffects, psEff)
	}
	log.Printf("%d page state effects calculated", len(pageStateEffects))

	//--------------------------------------------------------
	// 3. apply page-state operation and write states to files
	//--------------------------------------------------------
	state, err := processing.NewPageStateGitProcessorFromGit(repoUrl)
	if err != nil {
		return fmt.Errorf("NewPageStateGitProcessorFromGit() failed: %v", err)
	}

	for i, step := range stepEffects {
		state.TransitionToNext()
		op, err := pageStateEffects[i].ToOperation()
		if err != nil {
			return fmt.Errorf("ToOperation() failed at %s: %v", step.Step, err)
		}

		var nextStep string
		if i == len(stepEffects)-1 {
			nextStep = ""
		} else {
			nextStep = stepEffects[i+1].Step
		}

		if err := state.RegisterNext(nextStep, &op); err != nil {
			return fmt.Errorf("RegisterNext() failed at step %s: %v", step.Step, err)
		}

		fileName := fmt.Sprintf(dirName+"/state/state-%s.json", step.Step)
		if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("WriteJsonToFile() failed at step %s: %v", step.Step, err)
		}

		// iterate over to the next state
		if err := state.TransitionToNext(); err != nil {
			return fmt.Errorf("TransitionToNext() failed at step %s: %v", step.Step, err)
		}
	}
	// last state writes to the file
	lastStep := stepEffects[len(stepEffects)-1].Step
	fileName := fmt.Sprintf(dirName+"/state/state-%s.json", lastStep)
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
		scEff := SourceCodeEffect{step.SeqNo, fEffs, nil}

		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// PageStateEffect for seqNo
		psEff := PageStateEffect{step.SeqNo, &scEff, nil, tEff}
		pageStateEffects = append(pageStateEffects, psEff)
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

		fileName := fmt.Sprintf("data/state/state-%s.json", stepEffects[i].CurrentStep)
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
	fileName := fmt.Sprintf("data/state/state-%s.json", lastStep)
	if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("WriteJsonToFile() in PageStateProcessor failed: %v", err)
	}

	log.Printf("finished writing state into files")
	return nil
}
