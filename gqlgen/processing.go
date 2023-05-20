package main

import (
	"fmt"
	"log"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/effect"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

func processingCoreLogic(dirName string, state *processing.PageStateProcessor) error {
	//------------------------------------
	// 1. read effects for git repository
	//------------------------------------
	stepEffects, err := effect.ReadStepEffects(dirName + "/step-effects.json")
	if err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}
	log.Printf("%d step effects read ", len(stepEffects))

	fileEffects, err := effect.ReadFileEffects(dirName + "/file-effects.json")
	if err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}
	log.Printf("%d file effects read ", len(fileEffects))

	terminalEffects, err := effect.ReadTerminalEffects(dirName + "/terminal-effects.json")
	if err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}
	log.Printf("%d terminal effects read ", len(terminalEffects))

	markdownEffects, err := effect.ReadMarkdownEffects(dirName + "/markdown-effects.json")
	if err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}
	log.Printf("%d markdown effects read ", len(markdownEffects))

	//------------------------------
	// 2. construct page-sate effect
	//------------------------------
	var pageStateEffects []effect.PageStateEffect
	for _, step := range stepEffects {
		// TerminalEffect for seqNo
		tEff := terminalEffects.FindBySeqNo(step.SeqNo)

		// SourceCodeEffect for seqNo
		fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
		scEff := effect.SourceCodeEffect{SeqNo: step.SeqNo, CommitHash: step.CommitHash, FileEffects: fEffs, DefaultOpenFilePath: nil}

		// PageStateEffect for seqNo
		psEff := effect.PageStateEffect{SeqNo: step.SeqNo, SourceCodeEffect: &scEff, TerminalEffect: tEff}
		pageStateEffects = append(pageStateEffects, psEff)
	}
	log.Printf("%d page state effects calculated", len(pageStateEffects))

	//--------------------------------------------------------
	// 3. apply page-state operation and write states to files
	//--------------------------------------------------------

	for i, step := range stepEffects {
		state.TransitionToNext()
		op, err := pageStateEffects[i].ToOperation()
		if err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, step.Step, err)
		}

		var nextStep string
		if i == len(stepEffects)-1 {
			nextStep = ""
		} else {
			nextStep = stepEffects[i+1].Step
		}

		if err := state.RegisterNext(nextStep, &op); err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, step.Step, err)
		}

		fileName := fmt.Sprintf(dirName+"/state/state-%s.json", step.Step)
		if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("WriteJsonToFile() failed at step[%d] %s: %v", i, step.Step, err)
		}

		// iterate over to the next state
		if err := state.TransitionToNext(); err != nil {
			return fmt.Errorf("TransitionToNext() failed at step[%d] %s: %v", i, step.Step, err)
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

func GitEffectProcessing(dirName, repoUrl string) error {
	log.Printf("GitEffectProcessing started for dirName = %s, repoUrl = %s", dirName, repoUrl)
	state, err := processing.NewPageStateGitProcessorFromGit(repoUrl)
	if err != nil {
		return fmt.Errorf("NewPageStateGitProcessorFromGit() failed: %v", err)
	}
	return processingCoreLogic(dirName, state)
}

func EffectProcessing(dirName string) error {
	log.Printf("EffectProcessing started for dirName = %s", dirName)
	state := processing.NewPageStateProcessor()
	return processingCoreLogic(dirName, state)
}
