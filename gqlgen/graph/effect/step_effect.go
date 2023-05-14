package effect

import (
	"encoding/json"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
)

type StepEffect struct {
	SeqNo       int    `json:"seqNo"`
	CurrentStep string `json:"currentStep"`
	NextStep    string `json:"nextStep"`
}

type GiStepEffect struct {
	CurrentStep string
	NextStep    string
	CommitHash  string
}

func ReadStepEffects(filePath string) ([]StepEffect, error) {
	var effects []StepEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := internal.JsonRead(filePath, unmarshaller)
	if err != nil {
		return nil, fmt.Errorf("ReadStepEffects failed to read file, %s", err)
	}

	for i, effect := range effects {
		if effect.SeqNo != i {
			return nil, fmt.Errorf("ReadStepEffects failed to validate, effects[%d].SeqNo must be = %d, but %d", i, i, effect.SeqNo)
		}
	}

	return effects, err
}

func GitStepEffects(repo *git.Repository) ([]GiStepEffect, error) {
	var effects []GiStepEffect

	return effects, nil
}
