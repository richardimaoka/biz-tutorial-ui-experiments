package testio

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

func writeGoldenFile(t *testing.T, filePath string, v any) {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("error writing golden file = '%s', %v", filePath, err)
	}

	if err := os.WriteFile(filePath, jsonBytes, 0644); err != nil {
		t.Fatalf("error writing golden file = '%s', %v", filePath, err)
	}
}

func CompareAfterMarshal(t *testing.T, expectedJsonFile string, result interface{}) {
	expectedBytes, err := os.ReadFile(expectedJsonFile)
	if err != nil {
		t.Errorf("failed to read file = '%s'", expectedJsonFile)
		return
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		t.Errorf("failed to get marshaled result to compare against file = '%s'", expectedJsonFile)
		return
	}

	if err := compareJsonBytes(expectedBytes, resultBytes); err != nil {
		t.Fatalf("failed to match file = '%s' with result, %s", expectedJsonFile, err)
	}
}

func CompareWithGoldenFile(t *testing.T, updateGoldenFile bool, goldenFileName string, result interface{}) {
	if updateGoldenFile {
		writeGoldenFile(t, goldenFileName, result)
	}
	CompareAfterMarshal(t, goldenFileName, result)
}

func FilesMustUnmatch(t *testing.T, file1, file2 string) {
	bytes1, err := os.ReadFile(file1)
	if err != nil {
		t.Fatalf("FilesMustUnmatch - failed to read %s", file1)
	}

	bytes2, err := os.ReadFile(file2)
	if err != nil {
		t.Fatalf("FilesMustUnmatch - failed to read %s", file2)
	}

	if string(bytes1) == string(bytes2) {
		t.Fatalf("FilesMustUnmatch - %s and %s must unmatch but same", file1, file2)
	}
}
