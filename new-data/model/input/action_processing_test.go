package input

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
func TestFilesInDir(t *testing.T) {
	dir := "testdata"
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
	dataDir := "../../data"
	targetDir := fmt.Sprintf("%s/test", dataDir)
	targetPrefix := "input"
	actionListFile := "testdata/action_list.json"

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
	SplitActionListFile(actionListFile, targetDir, targetPrefix)

	// from here checking result
	expectedFiles, err := FilesInDir("testdata", targetPrefix)
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
			return
		}

		resultBytes, err := os.ReadFile(resultFiles[i])
		if err != nil {
			t.Errorf("failed to read %s", resultFiles[i])
			return
		}

		if err := compareJsonBytes(expectedBytes, resultBytes); err != nil {
			t.Fatalf("failed to compare files = %s vs. %s", expectedFiles[i], resultFiles[i])
		}
	}
}
