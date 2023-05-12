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

	pageStateEffects, err := MergeEffects(terminalEffects, fileEffects)

	fmt.Println(stepEffects)
	fmt.Println(pageStateEffects)

	return nil
}
