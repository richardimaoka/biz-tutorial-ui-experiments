package model2

import "testing"

func TestFilePathLessFile(t *testing.T) {
	path1 := "aaa"
	path2 := "aaab"
	if lessFilePath2(path1, path2) != true {
		t.Errorf("path1 = %s < path2 = %s is expected, but they did not make it", path1, path2)
	}
}
