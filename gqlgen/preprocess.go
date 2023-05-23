package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/effect"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

func processingCoreLogic(dirName string, state *processing.PageStateProcessor) error {
	pageStateEffects, err := effect.ConstructPageStateEffects(
		dirName+"/step-effects.json",
		dirName+"/file-effects.json",
		dirName+"/terminal-effects.json",
		dirName+"/markdown-effects.json",
	)
	if err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}

	//--------------------------------------------------------
	// 3. apply page-state operation and write states to files
	//--------------------------------------------------------
	for i, psEff := range pageStateEffects {
		currentStep := psEff.Step

		op, err := psEff.ToOperation()
		if err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, currentStep, err)
		}

		if err := state.RegisterNext(psEff.NextStep, &op); err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, currentStep, err)
		}

		fileName := fmt.Sprintf(dirName+"/state/state-%s.json", currentStep)
		if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("WriteJsonToFile() failed at step[%d] %s: %v", i, currentStep, err)
		}

		// iterate over to the next state
		if err := state.TransitionToNext(); err != nil {
			return fmt.Errorf("TransitionToNext() failed at step[%d] %s: %v", i, currentStep, err)
		}
	}
	// last state writes to the file
	lastStep := pageStateEffects[len(pageStateEffects)-1].Step
	fileName := fmt.Sprintf(dirName+"/state/state-%s.json", lastStep)
	if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("WriteJsonToFile() in PageStateProcessor failed: %v", err)
	}

	//--------------------------------------------------------
	// 4. write first step to file
	//--------------------------------------------------------
	firstStepJsonValue, err := json.Marshal(pageStateEffects[0].Step)
	if err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}
	if err := os.WriteFile(dirName+"/first-step.json", firstStepJsonValue, 0644); err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}
	log.Printf("wrote first step to file")

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
