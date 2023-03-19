package model

import (
	"fmt"
	"os"
	"testing"
)

func TestFilesInDir(t *testing.T) {
	dir := "testdata/action/input"
	prefix := "input"

	resultFiles, err := FilesInDir(dir, prefix)
	if err != nil {
		t.Fatalf("error reading files in %s with prefix = %s", dir, prefix)
	}

	if len(resultFiles) != 46 {
		t.Errorf("expected 46 but got %d", len(resultFiles))
	}
}

func TestActionProcessing(t *testing.T) {
	dataDir := "../data"
	targetDir := fmt.Sprintf("%s/test", dataDir)
	targetPrefix := "input"
	actionListFile := "testdata/action/action_list.json"

	// initial setup
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, 0755); err != nil {
			t.Fatalf("failed to create %s", dataDir)
		}
	}
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		if err := os.Mkdir(targetDir, 0755); err != nil {
			t.Fatalf("failed to create %s", targetDir)
		}
	}

	// the function to test
	if err := SplitActionList(actionListFile, targetDir, targetPrefix); err != nil {
		t.Fatal(err)
	}

	// from here checking result
	expectedFiles, err := FilesInDir("testdata/action/input", targetPrefix)
	if err != nil {
		t.Fatalf("error reading files in testdata with prefix = %s", targetPrefix)
	}

	resultFiles, err := FilesInDir(targetDir, targetPrefix)
	if err != nil {
		t.Fatalf("error reading files in %s with prefix = %s", targetDir, targetPrefix)
	}

	if len(expectedFiles) != len(resultFiles) {
		t.Fatalf("expected %d files but result is %d files", len(expectedFiles), len(resultFiles))
	}

	for i := range expectedFiles {
		expectedBytes, err := os.ReadFile(expectedFiles[i])
		if err != nil {
			t.Errorf("failed to read %s", expectedFiles[i])
			continue
		}

		resultBytes, err := os.ReadFile(resultFiles[i])
		if err != nil {
			t.Errorf("failed to read %s", resultFiles[i])
			continue
		}

		if err := compareJsonBytes(expectedBytes, resultBytes); err != nil {
			t.Errorf("failed to compare files = %s vs. %s, %s", expectedFiles[i], resultFiles[i], err)
		}
	}
}
