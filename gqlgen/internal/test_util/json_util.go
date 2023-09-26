package test_util

import (
	"encoding/json"
	"os"
	"testing"
)

func JsonRead(t *testing.T, filePath string, v interface{}) {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read file %s, %s", filePath, err)
		return
	}

	err = json.Unmarshal(jsonBytes, v)
	if err != nil {
		t.Fatalf("failed to read file %s, %s", filePath, err)
		return
	}
}
