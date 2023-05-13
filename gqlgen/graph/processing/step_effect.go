package processing

import (
	"encoding/json"
	"fmt"
)

type StepEffect struct {
	SeqNo       int    `json:"seqNo"`
	CurrentStep string `json:"currentStep"`
	NextStep    string `json:"nextStep"`
}

func ReadStepEffects(filePath string) ([]StepEffect, error) {
	var effects []StepEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead(filePath, unmarshaller)
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
