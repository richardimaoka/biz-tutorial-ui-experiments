package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Address(s string) *string {
	return &s
}

func compareJsonBytes(expectedBytes, resultBytes []byte) error {
	var resultMap interface{}
	err := json.Unmarshal(resultBytes, &resultMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal result json, %s", err)
	}

	var expectedMap interface{}
	err = json.Unmarshal(expectedBytes, &expectedMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal expected json, %s", err)
	}

	if diff := cmp.Diff(expectedMap, resultMap); diff != "" {
		return fmt.Errorf("mismatch (-expected +result):\n%s", diff)
	}

	return nil
}

func CompareAfterMarshal(t *testing.T, expectedJsonFile string, result interface{}) {
	expectedBytes, err := os.ReadFile(expectedJsonFile)
	if err != nil {
		t.Errorf("%s failed to read", expectedJsonFile)
		return
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		t.Errorf("%s failed to get marshaled result to compare against", expectedJsonFile)
		return
	}

	if err := compareJsonBytes(expectedBytes, resultBytes); err != nil {
		t.Fatalf("%s failed to match with result, %s", expectedJsonFile, err)
	}
}

func writeGoldenFile(t *testing.T, filePath string, v any) {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("error writing golden file %s: %v", filePath, err)
	}

	if err := os.WriteFile(filePath, jsonBytes, 0644); err != nil {
		t.Fatalf("error writing golden file %s: %v", filePath, err)
	}
}

func CompareWitGoldenFile(t *testing.T, updateGoldenFile bool, goldenFileName string, result interface{}) {
	if updateGoldenFile {
		writeGoldenFile(t, goldenFileName, result)
	}
	CompareAfterMarshal(t, goldenFileName, result)
}
