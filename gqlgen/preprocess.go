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
	effects, err := effect.ConstructTransitionEffects(dirName)
	if err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}

	//--------------------------------------------------------
	// 2. prepare state dir
	//--------------------------------------------------------

	if err := os.RemoveAll(dirName + "/state"); err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}

	if err := os.MkdirAll(dirName+"/state", 0744); err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}

	//--------------------------------------------------------
	// 3. apply page-state operation and write states to files
	//--------------------------------------------------------
	initialStep := "initial"
	currentStep := initialStep //current step starts **BEFORE** seqNo = 0
	for i := 0; i < len(effects)-1; i++ {
		eff := effects[i]
		nextStep := eff.Step

		nextOperation, err := eff.Effect.ToOperation()
		if err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, nextStep, err)
		}

		if err := state.RegisterNext(nextStep, &nextOperation); err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, nextStep, err)
		}

		// after registering NEXT step, write the CURRENT state to file
		fileName := fmt.Sprintf(dirName+"/state/state-%s.json", currentStep)
		if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("WriteJsonToFile() failed at step[%d] %s: %v", i, nextStep, err)
		}

		// iterate over to the next state
		currentStep = nextStep
		if err := state.TransitionToNext(); err != nil {
			return fmt.Errorf("TransitionToNext() failed at step[%d] %s: %v", i, nextStep, err)
		}
	}
	// last state writes to the file
	lastStep := effects[len(effects)-1].Step
	fileName := fmt.Sprintf(dirName+"/state/state-%s.json", lastStep)
	if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("WriteJsonToFile() in PageStateProcessor failed: %v", err)
	}

	//--------------------------------------------------------
	// 4. write first step to file
	//--------------------------------------------------------
	firstStepJsonValue, err := json.Marshal(initialStep)
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
