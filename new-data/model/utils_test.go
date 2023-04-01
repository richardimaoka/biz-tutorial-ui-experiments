package model

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func address(s string) *string {
	return &s
}

func compareJsonBytes(expectedBytes, resultBytes []byte) error {
	var resultMap map[string]interface{}
	err := json.Unmarshal(resultBytes, &resultMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal result json %s", err)
	}

	var expectedMap map[string]interface{}
	err = json.Unmarshal(expectedBytes, &expectedMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal expected json %s", err)
	}

	if diff := cmp.Diff(expectedMap, resultMap); diff != "" {
		return fmt.Errorf("mismatch (-expected +result):\n%s", diff)
	}

	return nil
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

	if err := compareJsonBytes(expectedBytes, resultBytes); err != nil {
		t.Fatalf("failed to compare after marshal where expected file = %s, %s", expectedJsonFile, err)
	}
}
