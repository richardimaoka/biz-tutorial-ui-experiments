package read

import (
	"encoding/json"
	"fmt"
	"os"
)

type Step struct {
	SeqNo    int    `json:"seqNo"`
	Step     string `json:"step"`
	NColumns int    `json:"nColumns"`
}

type Steps []Step

func ReadSteps(filePath string) (Steps, error) {
	funcName := "ReadSteps"
	var entries Steps

	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%s failed to read file = %s, %v", funcName, filePath, err)
	}

	json.Unmarshal(jsonBytes, &entries)
	if err != nil {
		return nil, fmt.Errorf("%s failed to unmarshal file = %s, %v", funcName, filePath, err)
	}

	return entries, err
}
