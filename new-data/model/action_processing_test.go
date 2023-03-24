package model

import (
	"os"
	"testing"
)

const (
	dataDir string = "../data"
	testDir string = "../data/test"
)

func address(s string) *string {
	return &s
}

func TestFilesInDir(t *testing.T) {
	dir := "testdata/action/input"
	prefix := "action"

	resultFiles, err := FilesInDir(dir, prefix)
	if err != nil {
		t.Fatalf("error reading files in %s with prefix = %s", dir, prefix)
	}

	if len(resultFiles) != 47 {
		t.Errorf("expected 47 but got %d", len(resultFiles))
	}
}

func TestSplitActoinList(t *testing.T) {
	targetPrefix := "action"
	actionListFile := "testdata/action/action_list.json"

	// the function to test

	if err := SplitActionList(actionListFile, testDir, targetPrefix); err != nil {
		t.Fatal(err)
	}

	// expectation and results

	expectedFiles, err := FilesInDir("testdata/action/input", targetPrefix)
	if err != nil {
		t.Fatalf("error reading files in testdata with prefix = %s", targetPrefix)
	}

	resultFiles, err := FilesInDir(testDir, targetPrefix)
	if err != nil {
		t.Fatalf("error reading files in %s with prefix = %s", testDir, targetPrefix)
	}

	// run check result

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

func TestMain(m *testing.M) {
	// initial setup
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, 0755); err != nil {
			panic("failed to create " + dataDir)
		}
	}
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		if err := os.Mkdir(testDir, 0755); err != nil {
			panic("failed to create " + testDir)
		}
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}
