package input

import "testing"

func TestFilesInDir(t *testing.T) {
	targetDir := "../../data/test"
	targetPrefix := "input"

	resultFiles, err := FilesInDir("testdata", targetPrefix)
	if err != nil {
		t.Fatalf("error reading files in %s with prefix = %s", targetDir, targetPrefix)
	}

	if len(resultFiles) != 46 {
		t.Errorf("expected 46 but got %d", len(resultFiles))
	}
}

func TestActionProcessing(t *testing.T) {
	targetDir := "../../data/test"
	targetPrefix := "input"
	actionListFile := "testdata/action_list.json"

	SplitActionListFile(actionListFile, targetDir, targetPrefix)

	// expectedFiles, err := FilesInDir(targetDir, targetPrefix)
	// if err != nil {
	// 	t.Fatal("error reading files in %s with prefix = %s", targetDir, targetPrefix)
	// }
	// resultFiles, err := FilesInDir(targetDir, targetPrefix)
	// if err != nil {
	// 	t.Fatal("error reading files in %s with prefix = %s", targetDir, targetPrefix)
	// }
}
