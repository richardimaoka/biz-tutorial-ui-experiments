package main

import (
	"fmt"
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/effect"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

func processingCoreLogic(dirName string, state *processing.PageStateProcessor) error {
	//--------------------------------------------------------
	// 1. construct page state effects
	//--------------------------------------------------------
	effects, err := effect.ConstructPageStateEffects(dirName)
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
	// 3. write initial step to file
	//--------------------------------------------------------
	initialStep := state.CurrentStep()
	if err := internal.WriteJsonValueToFile(initialStep, dirName+"/initial-step.json"); err != nil {
		return fmt.Errorf("processingCoreLogic failed: %v", err)
	}
	log.Printf("wrote initial step to file")

	//--------------------------------------------------------
	// 4. apply page-state operation and write states to files
	//--------------------------------------------------------
	for i := 0; i < len(effects); i++ {
		eff := effects[i]
		nextStep := eff.Step

		nextOperation, err := eff.ToOperation()
		if err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, nextStep, err)
		}

		if err := state.RegisterNext(nextStep, &nextOperation); err != nil {
			return fmt.Errorf("processingCoreLogic failed at step[%d] %s: %v", i, nextStep, err)
		}

		// after registering NEXT step, write the CURRENT state to file
		fileName := fmt.Sprintf(dirName+"/state/%s.json", state.CurrentStep())
		if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("WriteJsonToFile() failed at step[%d] %s: %v", i, nextStep, err)
		}

		// iterate over to the next state
		if err := state.TransitionToNext(); err != nil {
			return fmt.Errorf("TransitionToNext() failed at step[%d] %s: %v", i, nextStep, err)
		}
	}
	// last state writes to the file
	lastStep := effects[len(effects)-1].Step
	fileName := fmt.Sprintf(dirName+"/state/%s.json", lastStep)
	if err := internal.WriteJsonToFile(state.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("WriteJsonToFile() in PageStateProcessor failed: %v", err)
	}

	log.Printf("finished writing state into files")
	return nil
}

func GitEffectProcessing(dirName, repoUrl string) error {
	log.Printf("GitEffectProcessing started for dirName = %s, repoUrl = %s", dirName, repoUrl)
	state, err := processing.NewPageStateProcessorGit(repoUrl)
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
