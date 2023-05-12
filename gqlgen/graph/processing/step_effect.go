package processing

import "encoding/json"

type StepEffect struct {
	SeqNo int    `json:"seqNo"`
	Step  string `json:"step"`
}

func ReadStepEffects(filePath string) ([]StepEffect, error) {
	var effects []StepEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead("ReadStepEffects", filePath, unmarshaller)
	return effects, err
}
