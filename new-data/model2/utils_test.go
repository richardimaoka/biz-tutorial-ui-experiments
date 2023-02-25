package model2

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func prettyString(m map[string]interface{}) string {
	jsonString, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonString)
}

func compareJsonBytes(t *testing.T, expectedBytes, resultBytes []byte) {
	var resultMap map[string]interface{}
	err := json.Unmarshal(resultBytes, &resultMap)
	if err != nil {
		t.Errorf("failed to unmarshal result json")
		return
	}

	var expectedMap map[string]interface{}
	err = json.Unmarshal(expectedBytes, &expectedMap)
	if err != nil {
		t.Errorf("failed to unmarshal expected json")
		return
	}

	if diff := cmp.Diff(expectedMap, resultMap); diff != "" {
		t.Errorf("mismatch (-expected +result):\n%s", diff)
	}
}

func compareAfterMarshal(t *testing.T, expectedJsonFile string, result interface{}) {
	expectedBytes, err := os.ReadFile(expectedJsonFile)
	if err != nil {
		t.Errorf("failed to read %s", expectedJsonFile)
		return
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		t.Errorf("failed to marshal result")
		return
	}

	compareJsonBytes(t, expectedBytes, resultBytes)
}
